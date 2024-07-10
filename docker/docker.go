package docker

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/akshay0074700747/my-sandbox/model"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func ExecuteOnDocker(req model.CodeExecutionRequest) (string, error) {

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return "", err
	}

	filename := strings.Split(req.FileName, ".")[0]
	tmpfile, err := os.CreateTemp("", fmt.Sprintf("%s_*.go", filename))
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(req.SourceCode); err != nil {
		return "", err
	}

	if err := tmpfile.Close(); err != nil {
		return "", err
	}

	// Create a container
	resp, err := cli.ContainerCreate(context.Background(),
		&container.Config{
			Image: "golang:alpine",
			Cmd:   []string{"go", "run", "/code/main.go"},
		},
		&container.HostConfig{
			Binds: []string{fmt.Sprintf("%s:/code/main.go", tmpfile.Name())},
		},
		nil, nil, "")
	if err != nil {
		return "", err
	}

	// Start the container
	if err := cli.ContainerStart(context.Background(), resp.ID, container.StartOptions{}); err != nil {
		return "", err
	}

	// Wait for the container to finish
	statusCh, errCh := cli.ContainerWait(context.Background(), resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return "", err
		}
	case <-statusCh:
	}

	// Get the logs from the container
	out, err := cli.ContainerLogs(context.Background(), resp.ID, container.LogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		return "", err
	}

	// io.Copy(os.Stdout, out)
	res := make([]byte, 10000)
	if _, err = out.Read(res); err != nil {
		return "", err
	}

	// Remove the container
	err = cli.ContainerRemove(context.Background(), resp.ID, container.RemoveOptions{})
	if err != nil {
		return string(res), err
	}

	return string(res), nil
}
