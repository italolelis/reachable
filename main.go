package main

import (
	"context"
	"fmt"
	"os"

	"github.com/italolelis/reachable/cmd"
)

func main() {
	ctx := context.Background()
	if err := cmd.NewRootCmd(ctx).Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
