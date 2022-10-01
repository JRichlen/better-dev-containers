package util

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/moby/term"
	"github.com/spf13/viper"
	"os"
)

func getClientAndContext() (*client.Client, context.Context) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	panicIfError(err)
	return cli, ctx
}

func PullContainerImage() {
	cli, ctx := getClientAndContext()

	containerImage := viper.GetString("image")
	reader, err := cli.ImagePull(ctx, containerImage, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer reader.Close()
	termFd, isTerm := term.GetFdInfo(os.Stderr)
	jsonmessage.DisplayJSONMessagesStream(reader, os.Stderr, termFd, isTerm, nil)
}

func RunCommandInContainer(cmd []string) {
	containerImage := viper.GetString("image")
	cli, ctx := getClientAndContext()

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: containerImage,
		Cmd:   cmd,
		Tty:   false,
	}, nil, nil, nil, "")
	panicIfError(err)

	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	panicIfError(err)

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		panicIfError(err)
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	panicIfError(err)

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

	cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{
		RemoveVolumes: true,
		RemoveLinks:   true,
		Force:         true,
	})
}
