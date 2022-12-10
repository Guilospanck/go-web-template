package main

import (
	"context"
	"go-web-template/cmd"
	_ "go-web-template/pkg/infrastructure/environments"
	"log"
)

func main() {
	ctx := context.Background()

	log.Fatal(cmd.Execute(ctx))
}
