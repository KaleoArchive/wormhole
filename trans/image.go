package trans

import (
	"github.com/docker/distribution/digest"
	"github.com/docker/distribution/manifest/schema2"
	"github.com/kaleocheng/docker-registry-client/registry"
)

// Image struct
type Image struct {
	Repository string
	Reference  string
	Manifest   *schema2.DeserializedManifest
	Digest     digest.Digest
}

// GetImage return a Image struct
func GetImage(r *registry.Registry, repository, reference string) (*Image, error) {
	i := &Image{
		Repository: repository,
		Reference:  reference,
	}
	manifest, err := r.ManifestV2(i.Repository, i.Reference)
	if err != nil {
		return nil, err
	}

	i.Manifest = manifest

	digest, err := r.ManifestDigest(i.Repository, i.Reference)
	if err != nil {
		return nil, err
	}
	i.Digest = digest

	return i, nil
}
