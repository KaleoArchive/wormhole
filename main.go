package main

import (
	"fmt"

	"github.com/kaleocheng/docker-registry-client/registry"
	"github.com/kaleocheng/wormhole/job"
	"github.com/kaleocheng/wormhole/trans"
)

func main() {

	url := "http://localhost:5001"
	username := "" // anonymous
	password := "" // anonymous
	hub, err := registry.New(url, username, password)
	if err != nil {
		return
	}

	url2 := "http://localhost:5002"
	username2 := "" // anonymous
	password2 := "" // anonymous
	hub2, err := registry.New(url2, username2, password2)
	if err != nil {
		return
	}

	t := trans.NewTrans(hub, hub2)

	image, err := trans.GetImage(hub, "library/alpine", "latest")
	if err != nil {
		fmt.Println(err)
		return
	}

	job.Start(t)
	defer job.Close()
	job.SetRateLimit(10000 * 1024)

	if _, err := job.Add(image); err != nil {
		fmt.Println(err)
		return
	}
}
