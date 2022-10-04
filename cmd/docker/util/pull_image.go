package util

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/moby/term"
	"github.com/spf13/viper"
	"os"
)

func getOutputWriter() *os.File {
	if viper.GetBool("verbose") {
		return os.Stdout
	}
	return nil
}

func PullImage(image string) {
	cli, ctx := GetClientAndContext()

	reader, err := cli.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer reader.Close()
	termFd, isTerm := term.GetFdInfo(os.Stderr)
	jsonmessage.DisplayJSONMessagesStream(reader, os.Stdout, termFd, isTerm, nil)
}
