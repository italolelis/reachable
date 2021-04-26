package cmd

import (
	"context"

	log "github.com/italolelis/reachable/pkg/log"
	"github.com/spf13/cobra"
)

type (
	// RootOptions represents the reachable global options
	RootOptions struct {
		verbose bool
	}
)

// NewRootCmd creates the root command
func NewRootCmd(ctx context.Context) *cobra.Command {
	var opts RootOptions
	ctx = log.NewContext(ctx)

	cmd := cobra.Command{
		Use:   "reachable",
		Short: "Reachable is a CLI tool to check if a domain is up",
		Example: `
		reachable check google.com
		
		reachable check google.com facebook.com twitter.com

		reachable check google.com -v
		`,
		PersistentPreRun: func(ccmd *cobra.Command, args []string) {
			if opts.verbose {
				log.SetLevel("debug")
			}
		},
	}

	cmd.PersistentFlags().BoolVarP(&opts.verbose, "verbose", "v", false, "Make the operation more talkative")

	// Aggregates Root commands
	cmd.AddCommand(NewCheckCmd(ctx))
	cmd.AddCommand(NewVersionCmd(ctx))

	return &cmd
}
