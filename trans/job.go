package trans

import (
	"github.com/docker/distribution/digest"
	"github.com/docker/distribution/manifest/schema2"
)

// Job struct
type Job struct {
	Repository string
	Reference  string
	Manifest   *schema2.DeserializedManifest
	Digest     digest.Digest
}
