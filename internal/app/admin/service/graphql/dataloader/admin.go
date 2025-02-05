package dataloader

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/go-keg/keg/contrib/gql"
	"github.com/go-keg/monorepo/internal/data/example/ent"
	"github.com/go-keg/monorepo/internal/data/example/ent/permission"
	"github.com/go-keg/monorepo/internal/data/example/ent/user"
	"github.com/graph-gophers/dataloader"
	"github.com/samber/lo"
)

type AdminLoader struct {
	client *ent.Client
}

func (r AdminLoader) permissionChildrenCount() gql.LoaderFunc {
	return func(ctx context.Context, keys dataloader.Keys) (map[dataloader.Key]any, error) {
		type item struct {
			ParentId int64 `json:"parent_id"`
			Count    int   `json:"count"`
		}
		var items []item
		err := r.client.Permission.Query().Where(permission.ParentIDIn(gql.ToInts(keys)...)).
			GroupBy(permission.FieldParentID).Aggregate(ent.Count()).Scan(ctx, &items)
		if err != nil {
			return nil, err
		}
		return lo.SliceToMap(items, func(item item) (dataloader.Key, any) {
			return gql.ToStringKey(item.ParentId), item.Count
		}), nil
	}
}

func (r AdminLoader) userRoleCount() gql.LoaderFunc {
	type item struct {
		ID    int64 `json:"id"`
		Count int   `json:"count"`
	}
	return func(ctx context.Context, keys dataloader.Keys) (map[dataloader.Key]any, error) {
		var items []item
		err := r.client.User.Query().Where(user.IDIn(gql.ToInts(keys)...)).Modify(func(s *sql.Selector) {
			ur := sql.Table(user.RolesTable).As("ur")
			s.Select(s.C(user.FieldID), sql.As(sql.Count(ur.C("role_id")), "count")).
				LeftJoin(ur).On(s.C(user.FieldID), ur.C("user_id")).
				GroupBy(s.C(user.FieldID))
		}).Scan(ctx, &items)
		if err != nil {
			return nil, err
		}
		return lo.SliceToMap(items, func(item item) (dataloader.Key, any) {
			return gql.ToStringKey(item.ID), item.Count
		}), nil
	}
}
