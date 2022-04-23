package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func main() {
	imageName := flag.String("image-name", "nginx", "Image name")
	flag.Parse()

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	ii, err := cli.ImageList(ctx, types.ImageListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{
			Key:   "reference",
			Value: *imageName,
		}),
	})
	if err != nil {
		panic(err)
	}

	for _, i := range ii {
		fmt.Println(strings.Join(i.RepoTags, ","))
	}
}
