package containers

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var FtpContainer testcontainers.Container

func RunFtpContainer(ctx context.Context, port string, user string, password string) error {
	req := testcontainers.ContainerRequest{
		Image:        "delfer/alpine-ftp-server:latest",
		ExposedPorts: []string{port},
		WaitingFor:   wait.ForLog("passwd"),
		Env: map[string]string{
			"USERS":    fmt.Sprintf("%s|%s", user, password),
			"ADDRESS":  "localhost",
			"MIN_PORT": "21000",
			"MAX_PORT": "21010",
		},
	}

	var err error
	FtpContainer, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return err
	}

	// defer ftpC.Terminate(ctx)

	return nil
}
