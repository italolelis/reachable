package cmd

import (
	"context"
	"time"

	"github.com/italolelis/reachable/pkg/log"
	"github.com/italolelis/reachable/pkg/reachable"
	"github.com/spf13/cobra"
)

// NewCheckCmd creates a new check command
func NewCheckCmd(ctx context.Context, timeout time.Duration) *cobra.Command {
	return &cobra.Command{
		Use:     "check",
		Short:   "Checks if a domain is reachable",
		Aliases: []string{"v"},
		Run: func(cmd *cobra.Command, args []string) {
			logger := log.WithContext(ctx)

			result, err := reachable.IsReachable(ctx, args[0], timeout)
			if err != nil {
				logger.Debug(err.Error())
				logger.Error("Not Reachable!")
				return
			}

			logger.Debugf("Domain %s", result.Domain)
			logger.Debugf("IP %s", result.IP)
			logger.Debugf("Status Code %d", result.StatusCode)
			logger.Debugf("DNS Lookup %d ms", int(result.Response.DNSLookup/time.Millisecond))
			logger.Debugf("TCP Connection %d ms", int(result.Response.TCPConnection/time.Millisecond))
			logger.Debugf("TLS Handshake %d ms", int(result.Response.TLSHandshake/time.Millisecond))
			logger.Debugf("Server Processing %d ms", int(result.Response.ServerProcessing/time.Millisecond))
			logger.Debugf("Content Transfer %d ms", int(result.Response.ContentTransfer(time.Now())/time.Millisecond))
			logger.Debugf("Total Time %d ms", int(result.Response.Total(time.Now())/time.Millisecond))

			logger.Info("Reachable!")
		},
	}
}
