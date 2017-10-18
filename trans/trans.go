package trans

import (
	"log"

	"github.com/docker/distribution/digest"
	"github.com/kaleocheng/docker-registry-client/registry"
)

// Trans struct
type Trans struct {
	Src *registry.Registry
	Dst *registry.Registry
}

// NewTrans return a new Trans
func NewTrans(src *registry.Registry, dst *registry.Registry) *Trans {
	return &Trans{
		Src: src,
		Dst: dst,
	}
}

// NewJob return a Migretion
func (t *Trans) NewJob(repository, reference string) (*Job, error) {
	j := &Job{
		Repository: repository,
		Reference:  reference,
	}
	manifest, err := t.Src.ManifestV2(j.Repository, j.Reference)
	if err != nil {
		return nil, err
	}

	j.Manifest = manifest

	digest, err := t.Src.ManifestDigest(j.Repository, j.Reference)
	if err != nil {
		return nil, err
	}
	j.Digest = digest

	return j, nil
}

// Migrate migrates repo:tag from SrcClient to DstClient
// If the image already exists in DstClient, It does nothing.
func (t *Trans) Migrate(repository, reference string) error {
	j, err := t.NewJob(repository, reference)
	if err != nil {
		return err
	}
	ok, err := t.Check(j)
	if err != nil {
		return err
	}
	if ok {
		return t.Start(j)
	}
	return nil
}

// Start starts a Job
func (t *Trans) Start(j *Job) error {
	if err := t.migrateConfig(j); err != nil {
		return err
	}
	if err := t.migrateLayers(j); err != nil {
		return err
	}
	digest, err := t.migrateManifest(j)
	log.Println(digest)
	return err
}

// Check if the image already exists in DstClient
func (t *Trans) Check(j *Job) (bool, error) {
	exist, err := t.Dst.HasManifest(j.Repository, j.Reference)
	if err != nil {
		return false, err
	}

	if !exist {
		return true, nil
	}

	digestDst, err := t.Dst.ManifestDigest(j.Repository, j.Reference)
	if err != nil {
		return false, err
	}

	if j.Digest != digestDst {
		return true, nil
	}

	return false, nil
}

func (t *Trans) migrateLayer(digest digest.Digest, repository string) error {
	exist, err := t.Dst.HasLayer(repository, digest)
	if err != nil {
		return err
	}

	if exist {
		return nil
	}

	reader, err := t.Src.DownloadLayer(repository, digest)
	if reader != nil {
		defer reader.Close()
	}

	if err != nil {
		return err
	}

	err = t.Dst.UploadLayer(repository, digest, reader)
	return err
}

func (t *Trans) migrateConfig(j *Job) error {
	return t.migrateLayer(j.Manifest.Config.Digest, j.Repository)
}

func (t *Trans) migrateLayers(j *Job) error {
	for _, l := range j.Manifest.Layers {
		if err := t.migrateLayer(l.Digest, j.Repository); err != nil {
			return err
		}
	}
	return nil
}

func (t *Trans) migrateManifest(j *Job) (string, error) {
	mediaType, payload, err := j.Manifest.Payload()
	if err != nil {
		return "", err
	}

	digest, err := t.Dst.PushManifest(j.Repository, j.Reference, mediaType, payload)
	if err != nil {
		return "", err
	}
	return digest, nil
}
