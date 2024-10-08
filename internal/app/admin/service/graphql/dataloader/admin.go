package dataloader

import (
	"context"
	"fmt"
	"github.com/go-keg/example/internal/data/example/ent"
	"github.com/go-keg/example/internal/data/example/ent/permission"
	"github.com/go-keg/keg/contrib/gql"
	"github.com/graph-gophers/dataloader"
	"github.com/samber/lo"
	"strings"
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
	return func(ctx context.Context, keys dataloader.Keys) (map[dataloader.Key]any, error) {
		placeholder := strings.TrimSuffix(strings.Repeat("?,", len(keys)), ",")
		var rows, err = r.client.QueryContext(ctx, fmt.Sprintf(`SELECT user_id, COUNT(role_id) FROM users LEFT JOIN user_roles ON users.id = user_roles.user_id WHERE user_id IN (%s) GROUP BY user_id`, placeholder), gql.ToAnySlice(keys)...)
		if err != nil {
			return nil, err
		}
		var result = make(map[dataloader.Key]any)
		for rows.Next() {
			var userId, count int
			err = rows.Scan(&userId, &count)
			if err != nil {
				return nil, err
			}
			result[gql.ToStringKey(userId)] = count
		}
		return result, nil
	}
}
