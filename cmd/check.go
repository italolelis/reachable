package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/apcera/termtables"
	alog "github.com/apex/log"
	"github.com/italolelis/reachable/pkg/log"
	"github.com/italolelis/reachable/pkg/reachable"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

// NewCheckCmd creates a new check command
func NewCheckCmd(ctx context.Context, timeout time.Duration) *cobra.Command {
	return &cobra.Command{
		Use:     "check",
		Short:   "Checks if a domain is reachable",
		Aliases: []string{"v"},
		Run: func(cmd *cobra.Command, args []string) {
			var wg errgroup.Group
			var results [][]interface{}
			// wg, ctx := errgroup.WithContext(ctx)
			logger := log.WithContext(ctx)

			for _, domain := range args {
				domain := domain
				wg.Go(func() error {
					result, err := reachable.IsReachable(ctx, domain, timeout)
					if err != nil {
						if logger.Level == alog.DebugLevel {
							logger.Error(err.Error())
						}

						logger.WithField("domain", domain).Error("Unreachable!")
						return err
					}

					results = append(results, []interface{}{
						domain,
						result.StatusCode,
						int(result.Response.DNSLookup / time.Millisecond),
						int(result.Response.TCPConnection / time.Millisecond),
						int(result.Response.TLSHandshake / time.Millisecond),
						int(result.Response.ServerProcessing / time.Millisecond),
						int(result.Response.ContentTransfer(time.Now()) / time.Hour),
						int(result.Response.Total(time.Now()) / time.Millisecond),
					})

					logger.WithField("domain", domain).Info("Reachable!")

					return nil
				})
			}

			wg.Wait()

			if logger.Level == alog.DebugLevel {
				table := termtables.CreateTable()
				table.AddHeaders("Domain", "Status Code", "DNS Lookup", "TCP Connection", "TLS Handshake", "Server Processing", "Content Transfer", "Total Time")

				for _, r := range results {
					table.AddRow(r...)
				}

				fmt.Printf(table.Render())
			}
		},
	}
}
