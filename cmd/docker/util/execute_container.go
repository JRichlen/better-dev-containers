package util

import (
	"bufio"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
)

func ExecuteContainer(containerId string) {
	cli, ctx := GetClientAndContext()

	waiter, err := cli.ContainerAttach(ctx, containerId, types.ContainerAttachOptions{
		Stderr: true,
		Stdout: true,
		Stdin:  true,
		Stream: true,
	})

	go io.Copy(os.Stdout, waiter.Reader)
	go io.Copy(os.Stderr, waiter.Reader)

	err = cli.ContainerStart(ctx, containerId, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	fd := int(os.Stdin.Fd())
	var oldState *terminal.State
	if terminal.IsTerminal(fd) {
		oldState, err = terminal.MakeRaw(fd)
		if err != nil {
			panic(err)
		}

		listenAndExitForCtrlC(containerId, waiter)
	}

	statusCh, errCh := cli.ContainerWait(ctx, containerId, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	if terminal.IsTerminal(fd) {
		terminal.Restore(fd, oldState)
	}
}

func listenAndExitForCtrlC(containerId string, waiter types.HijackedResponse) {
	// Wrapper around Stdin for the container, to detect Ctrl+C
	go func() {
		consoleReader := bufio.NewReaderSize(os.Stdin, 1)
		for {
			input, _ := consoleReader.ReadByte()
			// Ctrl-C = 3
			if input == 3 {
				cli, ctx := GetClientAndContext()
				// Tell docker to forcefully remove the container
				cli.ContainerRemove(ctx, containerId, types.ContainerRemoveOptions{
					Force: true,
				})
			}
			waiter.Conn.Write([]byte{input})
		}
	}()
}
