package main

import (
	"context"
	"fmt"

	"github.com/kutsenkoilya/go-demo-project/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
