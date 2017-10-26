package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kaleocheng/docker-registry-client/registry"
	"github.com/kaleocheng/wormhole/job"
	"github.com/kaleocheng/wormhole/trans"
)

func main() {

	router := gin.Default()
	router.Run()

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

	image, err := trans.GetImage(hub, "library/test", "latest")
	if err != nil {
		fmt.Println(err)
		return
	}

	image2, err := trans.GetImage(hub, "library/test2", "latest")
	if err != nil {
		fmt.Println(err)
		return
	}

	image3, err := trans.GetImage(hub, "library/test3", "latest")
	if err != nil {
		fmt.Println(err)
		return
	}

	image4, err := trans.GetImage(hub, "library/test4", "latest")
	if err != nil {
		fmt.Println(err)
		return
	}

	job.Start(t)
	defer job.Close()
	job.SetRateLimit(5000 * 1024)

	job.Add(image)
	job.Add(image2)
	job.Add(image3)
	job.Add(image4)

	always := make(chan int)
	<-always
}
