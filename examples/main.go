package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/spf13/cobra"
	snowgo "github.com/zeiss/snow-go"
	"github.com/zeiss/snow-go/apis"
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

	client, err := snowgo.New(cfg.Flags.URL, apis.WithRequestEditorFn(basicAuuth.Intercept))
	if err != nil {
		return err
	}

	res, err := client.GetApiNowTableTableName(ctx, "ecc_event", &apis.GetApiNowTableTableNameParams{})
	if err != nil {
		return err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
