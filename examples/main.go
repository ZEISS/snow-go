package main

import (
	"context"
	"log"
	"os"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/spf13/cobra"
	snowgo "github.com/zeiss/snow-go"
)

// Config ...
type Config struct {
	Flags *Flags
}

// Flags ...
type Flags struct {
	URL string
}

var cfg = &Config{
	Flags: &Flags{},
}

var rootCmd = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(cmd.Context())
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfg.Flags.URL, "url", "", "url")

	rootCmd.SilenceUsage = true
}

func run(ctx context.Context) error {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	basicAuuth, err := securityprovider.NewSecurityProviderBasicAuth("demo", "")
	if err != nil {
		return err
	}

	client := snowgo.New(cfg.Flags.URL, snowgo.WithRequestEditorFn(basicAuuth.Intercept))
	if err != nil {
		return err
	}

	err = client.Do(ctx, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
