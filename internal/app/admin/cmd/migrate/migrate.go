package migrate

import (
	"context"
	"fmt"
	"github.com/go-keg/example/internal/app/admin/conf"
	"github.com/go-keg/example/internal/app/admin/data"
	"github.com/go-keg/example/internal/data/example/ent"
	"github.com/go-keg/keg/contrib/config"
	"github.com/spf13/cobra"
	"log"
)

var Cmd = &cobra.Command{
	Use: "migrate",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("conf")
		cfg, err := config.Load[conf.Config](path)
		if err != nil {
			panic(err)
		}
		client, err := data.NewEntClient(cfg)
		if err != nil {
			panic(err)
		}
		defer func(client *ent.Client) {
			err := client.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(client)
		// Run the auto migration tool.
		if err := client.Debug().Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	},
}
