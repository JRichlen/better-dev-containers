package docker

import (
	"better-dev-container/cmd/docker/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RunCommandInContainer(cmd *cobra.Command, containerCmd []string) {
	containerImage := viper.GetString("image")
	containerId := util.CreateContainer(containerImage, containerCmd)

	util.ExecuteContainer(containerId)
}
