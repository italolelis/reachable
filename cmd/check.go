package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/apcera/termtables"
	"github.com/italolelis/reachable/pkg/log"
	"github.com/italolelis/reachable/pkg/reachable"
	"github.com/spf13/cobra"
)

type (
	// CheckOptions represents the check command options
	CheckOptions struct {
		timeout time.Duration
		method  string
	}
)

// NewCheckCmd creates a new check command
func NewCheckCmd(ctx context.Context) *cobra.Command {
	var opts CheckOptions

	cmd := cobra.Command{
		Use:     "check",
		Short:   "Checks if a domain is reachable",
		Aliases: []string{"v"},
		Run:     check(ctx, &opts),
	}

	cmd.PersistentFlags().DurationVarP(&opts.timeout, "timeout", "t", 30*time.Second, "Defines a timeout")
	cmd.PersistentFlags().StringVarP(&opts.method, "method", "m", "get", "Defines a HTTP method to be used")

	return &cmd
}

func check(ctx context.Context, opts *CheckOptions) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		logger := log.WithContext(ctx)

		chErr := make(chan error)
		defer close(chErr)

		go func() {
			for err := range chErr {
				logger.Error(err)
			}
		}()

		timeoutCtx, cancel := context.WithTimeout(ctx, opts.timeout)
		defer cancel()

		ch := make(chan reachable.Reachable)

		go reachable.IsReachableAsync(timeoutCtx, ch, chErr, args...)

		table := termtables.CreateTable()
		table.AddHeaders(
			"Domain",
			"Status Code",
			"DNS Lookup",
			"TCP Connection",
			"TLS Handshake",
			"Server Processing",
			"Content Transfer",
			"Total Time",
		)

		for {
			r, ok := <-ch
			if ok == false {
				break
			}

			table.AddRow(
				r.Domain,
				r.StatusCode,
				r.Response.DNSLookup.Round(time.Millisecond),
				r.Response.TCPConnection.Round(time.Millisecond),
				r.Response.TLSHandshake.Round(time.Millisecond),
				r.Response.ServerProcessing.Round(time.Microsecond),
				fmt.Sprintf("%4d ms", int(r.Response.ContentTransfer(time.Now())/time.Millisecond)),
				fmt.Sprintf("%4d ms", int(r.Response.Total(time.Now())/time.Millisecond)),
			)
		}

		fmt.Fprint(os.Stderr, table.Render())
	}
}
