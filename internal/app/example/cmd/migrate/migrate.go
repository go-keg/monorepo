package migrate

import (
	"fmt"
	"log"

	"github.com/go-keg/monorepo/internal/app/example/conf"
	"github.com/go-keg/monorepo/internal/app/example/data"
	"github.com/go-keg/monorepo/internal/data/example/ent"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "migrate",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("conf")
		cfg := conf.MustLoad(path)
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
		if err := client.Debug().Schema.Create(cmd.Context()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	},
}
