package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/zeiss/snow-go/push"

	cloudevents "github.com/cloudevents/sdk-go/v2"
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

	basicAuuth, err := snowgo.NewBasicAuth("demo", "")
	if err != nil {
		return err
	}

	client := snowgo.New(cfg.Flags.URL, snowgo.WithRequestEditorFn(basicAuuth.Intercept))

	event := cloudevents.NewEvent()
	event.SetID("example-uuid-32943bac6fea")
	event.SetSource("example/uri")
	event.SetType("example.type")
	event.SetData(cloudevents.ApplicationJSON, map[string]string{"hello": "world"})

	url := push.NewPushConnectorUrl("", "typhoon")

	req := push.NewRequest(url, event)
	res := &push.Response{}

	err = client.Do(ctx, req, res)
	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
