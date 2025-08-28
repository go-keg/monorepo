package seeds

import (
	"context"

	"github.com/go-keg/monorepo/internal/app/account/conf"
	"github.com/go-keg/monorepo/internal/app/account/data"
	"github.com/go-keg/monorepo/internal/data/account/ent"
	"github.com/go-keg/monorepo/internal/data/account/ent/permission"
	"github.com/go-keg/monorepo/internal/pkg/utils"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "seeds",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("conf")
		cfg := conf.MustLoad(path)
		client, err := data.NewEntClient(cfg)
		if err != nil {
			panic(err)
		}
		ctx := cmd.Context()

		_ = client.User.Create().SetID(1).
			SetEmail("admin@example.com").
			SetNickname("admin").
			SetAvatar("https://avatars.githubusercontent.com/u/22978107").
			SetPassword(utils.GeneratePassword("123456")).
			SetIsAdmin(true).
			OnConflict().Ignore().Exec(ctx)

		_ = client.Tenant.Create().SetID(1).
			SetName("Default Tenant").
			OnConflict().Ignore().Exec(ctx)

		_ = SeedPermissions(ctx, client)

		_ = client.TenantRole.Create().
			SetID(1).
			SetTenantID(1).
			SetName("管理员").
			AddPermissionIDs(lo.Map(permissions, func(item *ent.Permission, index int) int {
				return item.ID
			})...).
			OnConflict().Ignore().Exec(ctx)

		_ = client.TenantUser.Create().
			SetID(1).
			SetTenantID(1).
			SetUserID(1).
			SetIsOwner(true).
			AddRoleIDs(1).
			OnConflict().Ignore().Exec(ctx)

	},
}

var permissions = []*ent.Permission{
	{
		ID:   1,
		Name: "管理组织架构",
		Type: permission.TypeMenu,
		Key:  lo.ToPtr("manage.organizations"),
		Path: "/organizations",
	},
	{
		ID:   2,
		Name: "管理角色",
		Type: permission.TypeMenu,
		Key:  lo.ToPtr("manage.roles"),
		Path: "/roles",
	},
}

func SeedPermissions(ctx context.Context, client *ent.Client) error {

	bulk := lo.Map(permissions, func(item *ent.Permission, index int) *ent.PermissionCreate {
		return client.Permission.Create().SetID(item.ID).
			SetName(item.Name).
			SetType(item.Type).
			SetNillableKey(item.Key).
			SetPath(item.Path)
	})
	return client.Permission.CreateBulk(bulk...).OnConflict().Ignore().Exec(ctx)
}
