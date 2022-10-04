package util

import (
	"context"
	"github.com/docker/docker/client"
)

var ctx = context.Background()

func GetClientAndContext() (*client.Client, context.Context) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	return cli, ctx
}
