package cmd

import (
	"context"
	"time"

	"github.com/italolelis/reachable/pkg/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type (
	// RootOptions represents the ahoy global options
	RootOptions struct {
		timeout time.Duration
		verbose bool
	}
)

// NewRootCmd creates the root command
func NewRootCmd() *cobra.Command {
	opts := RootOptions{}
	ctx := log.NewContext(context.Background())

	cmd := cobra.Command{
		Use:   "reachable",
		Short: "Reachable is a CLI tool to check if a domain is up",
		PersistentPreRun: func(ccmd *cobra.Command, args []string) {
			if opts.verbose {
				log.WithContext(context.Background()).SetLevel(logrus.DebugLevel)
			}
		},
	}

	cmd.PersistentFlags().DurationVarP(&opts.timeout, "timeout", "t", 30*time.Second, "Defines a timeout")
	cmd.PersistentFlags().BoolVarP(&opts.verbose, "verbose", "v", false, "Make the operation more talkative")

	// Aggregates Root commands
	cmd.AddCommand(NewCheckCmd(ctx, opts.timeout))
	cmd.AddCommand(NewVersionCmd(ctx))

	return &cmd
}
