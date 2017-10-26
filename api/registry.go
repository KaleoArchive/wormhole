package api

import "github.com/kaleocheng/docker-registry-client/registry"

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

// GetRegistry ...
func GetRegistry(id string) *registry.Registry {
	return registryMap[id]
}

// AddRegistry ...
func AddRegistry(id string, r *registry.Registry) {
	registryMap[id] = r
}

// RegistryExist ...
func RegistryExist(r *Registry) bool {
	return true
}
