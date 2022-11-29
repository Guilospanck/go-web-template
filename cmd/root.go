package cmd

import (
	"context"
)

func Execute(ctx context.Context) error {
	container := NewContainer(ctx)

	container.httpServer.Setup(ctx)

	// register routes into http server interface
	container.authenticationPresenter.Register(container.httpServer)

	if err := container.httpServer.Run(); err != nil {
		return err
	}

	return nil
}
