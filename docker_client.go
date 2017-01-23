package dockersnitch

import (
    "fmt"
    docker "github.com/fsouza/go-dockerclient"
)

func Images() string {
    endpoint := "unix:///var/run/docker.sock"
    client, err := docker.NewClient(endpoint)
    if err != nil {
        panic(err)
    }
    imgs, err := client.ListImages(docker.ListImagesOptions{All: false})
    if err != nil {
        panic(err)
    }
    imageIDs := ""
    for _, img := range imgs {
        fmt.Println("ID: ", img.ID)
        fmt.Println("RepoTags: ", img.RepoTags)
        fmt.Println("Created: ", img.Created)
        fmt.Println("Size: ", img.Size)
        fmt.Println("VirtualSize: ", img.VirtualSize)
        fmt.Println("ParentId: ", img.ParentID)

        imageIDs += img.ID + ", "
    }

    return imageIDs
}