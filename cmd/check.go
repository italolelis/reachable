package cmd

import (
	"context"
	"fmt"
	"time"

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
			// wg, ctx := errgroup.WithContext(ctx)
			logger := log.WithContext(ctx)

			for _, domain := range args {
				domain := domain
				wg.Go(func() error {
					lg := logger.WithField("domain", domain)
					result, err := reachable.IsReachable(ctx, domain, timeout)
					if err != nil {
						if lg.Logger.Level == alog.DebugLevel {
							lg.Error(err.Error())
						}

						lg.Error("Unreachable!")
						return err
					}

					lg.Debugf("Domain %s", result.Domain)
					lg.Debugf("Port %s", result.Port)
					lg.Debugf("Status Code %d", result.StatusCode)
					lg.Debugf("DNS Lookup %d ms", int(result.Response.DNSLookup/time.Millisecond))
					lg.Debugf("TCP Connection %d ms", int(result.Response.TCPConnection/time.Millisecond))
					lg.Debugf("TLS Handshake %d ms", int(result.Response.TLSHandshake/time.Millisecond))
					lg.Debugf("Server Processing %d ms", int(result.Response.ServerProcessing/time.Millisecond))
					lg.Debugf("Content Transfer %d ms", int(result.Response.ContentTransfer(time.Now())/time.Millisecond))
					lg.Debugf("Total Time %d ms", int(result.Response.Total(time.Now())/time.Millisecond))

					lg.Info("Reachable!")

					if lg.Logger.Level == alog.DebugLevel {
						fmt.Println("")
					}
					return nil
				})
			}

			wg.Wait()
		},
	}
}
