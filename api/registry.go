package api

import (
	log "github.com/Sirupsen/logrus"
	"github.com/kaleocheng/docker-registry-client/registry"
	"github.com/segmentio/ksuid"
)

// Registry ...
type Registry struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	URL      string `json:"url"`
}

var registryMap map[string]*registry.Registry

func init() {
	registryMap = make(map[string]*registry.Registry)
}

// NewRegistry ...
func NewRegistry(r *Registry) (string, error) {
	ok, err := RegistryExist(r)
	if err != nil {
		return "", err
	}
	if ok {
		return r.ID, nil
	}

	hub, err := registry.New(r.URL, r.Username, r.Password)
	if err != nil {
		log.Error(err)
		return "", errCreateRegistryFailed
	}
	id := ksuid.New()
	r.ID = id.String()
	AddRegistry(r.ID, hub)
	return r.ID, nil
}

// GetRegistry ...
func GetRegistry(id string) *registry.Registry {
	return registryMap[id]
}

// AddRegistry ...
func AddRegistry(id string, r *registry.Registry) {
	registryMap[id] = r
}

// RegistryExist ...
func RegistryExist(r *Registry) (bool, error) {
	return false, nil
}
