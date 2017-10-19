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

// Migrate migrates repo:tag from SrcClient to DstClient
// If the image already exists in DstClient, It does nothing.
func (t *Trans) Migrate(i *Image) error {
	ok, err := t.Check(i)
	if err != nil {
		return err
	}
	if ok {
		return t.Start(i)
	}
	return nil
}

// Start starts a Job
func (t *Trans) Start(i *Image) error {
	if err := t.migrateConfig(i); err != nil {
		return err
	}
	if err := t.migrateLayers(i); err != nil {
		return err
	}
	digest, err := t.migrateManifest(i)
	log.Println(digest)
	return err
}

// Check if the image already exists in DstClient
func (t *Trans) Check(i *Image) (bool, error) {
	exist, err := t.Dst.HasManifest(i.Repository, i.Reference)
	if err != nil {
		return false, err
	}

	if !exist {
		return true, nil
	}

	digestDst, err := t.Dst.ManifestDigest(i.Repository, i.Reference)
	if err != nil {
		return false, err
	}

	if i.Digest != digestDst {
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

func (t *Trans) migrateConfig(i *Image) error {
	return t.migrateLayer(i.Manifest.Config.Digest, i.Repository)
}

func (t *Trans) migrateLayers(i *Image) error {
	for _, l := range i.Manifest.Layers {
		if err := t.migrateLayer(l.Digest, i.Repository); err != nil {
			return err
		}
	}
	return nil
}

func (t *Trans) migrateManifest(i *Image) (string, error) {
	mediaType, payload, err := i.Manifest.Payload()
	if err != nil {
		return "", err
	}

	digest, err := t.Dst.PushManifest(i.Repository, i.Reference, mediaType, payload)
	if err != nil {
		return "", err
	}
	return digest, nil
}
