package main

import (
	"context"
	"guest-bag-authentication/cmd"
	_ "guest-bag-authentication/pkg/infrastructure/environments"
	"log"
)

func main() {
	ctx := context.Background()

	log.Fatal(cmd.Execute(ctx))
}
