package util

import (
	"better-dev-container/cmd/lang"
	"github.com/docker/docker/api/types/container"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

func CreateContainer(containerImage string, cmd []string) (containerId string) {
	dockerWorkingDir := "/usr/src/project"
	workingDir := getWorkingDir()

	containerConfig := getContainerConfig(containerImage, cmd, dockerWorkingDir)
	hostConfig := getHostConfig(workingDir, dockerWorkingDir)

	PullImage(containerImage)

	cli, ctx := GetClientAndContext()
	resp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, "")
	if err != nil {
		panic(err)
	}

	return resp.ID
}

func getContainerConfig(containerImage string, cmd []string, dockerWorkingDir string) *container.Config {
	mappedCmd := lang.MapLanguageCommand(containerImage, cmd)
	return &container.Config{
		WorkingDir:   dockerWorkingDir,
		Image:        containerImage,
		Cmd:          mappedCmd,
		AttachStderr: true,
		AttachStdin:  true,
		Tty:          true,
		AttachStdout: true,
		OpenStdin:    true,
	}
}

func getHostConfig(workingDir string, dockerWorkingDir string) *container.HostConfig {
	return &container.HostConfig{
		AutoRemove: true,
		Binds: []string{
			strings.Join([]string{workingDir, dockerWorkingDir}, ":"),
		},
	}
}

func getWorkingDir() string {
	workingDir, err := filepath.Abs(viper.GetString("workingDir"))
	if err != nil {
		panic(err)
	}
	return workingDir
}
