package main

import (
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
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

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
