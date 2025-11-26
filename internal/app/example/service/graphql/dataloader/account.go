package dataloader

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/go-keg/keg/contrib/gql"
	"github.com/go-keg/monorepo/internal/data/example/ent"
	"github.com/go-keg/monorepo/internal/data/example/ent/membership"
	"github.com/go-keg/monorepo/internal/data/example/ent/organization"
	"github.com/go-keg/monorepo/internal/data/example/ent/permission"
	"github.com/go-keg/monorepo/internal/data/example/ent/tenantuser"
	"github.com/graph-gophers/dataloader"
	"github.com/samber/lo"
)

type AccountLoader struct {
	client *ent.Client
}

type countItem struct {
	ID    int64 `json:"id"`
	Count int   `json:"count"`
}

type parentCountItem struct {
	ParentID int64 `json:"parent_id"`
	Count    int   `json:"count"`
}

func (r AccountLoader) permissionChildrenCount() gql.LoaderFunc {
	return func(ctx context.Context, keys dataloader.Keys) (map[dataloader.Key]any, error) {
		var items []parentCountItem
		err := r.client.Permission.Query().Where(permission.ParentIDIn(gql.ToInts(keys)...)).
			GroupBy(permission.FieldParentID).Aggregate(ent.Count()).Scan(ctx, &items)
		if err != nil {
			return nil, err
		}
		return lo.SliceToMap(items, func(item parentCountItem) (dataloader.Key, any) {
			return gql.IntKey(item.ParentID), item.Count
		}), nil
	}
}

func (r AccountLoader) tenantUserRoleCount() gql.LoaderFunc {
	return func(ctx context.Context, keys dataloader.Keys) (map[dataloader.Key]any, error) {
		var result = make(map[dataloader.Key]any)
		var items []countItem
		err := r.client.TenantUser.Query().Where(
			tenantuser.IDIn(gql.ToInts(keys)...),
		).Modify(func(s *sql.Selector) {
			ur := sql.Table(tenantuser.RolesTable).As("ur")
			s.Select(s.C(tenantuser.FieldID), sql.As(sql.Count(ur.C("tenant_role_id")), "count")).
				LeftJoin(ur).On(s.C(tenantuser.FieldID), ur.C("tenant_user_id")).
				GroupBy(s.C(tenantuser.FieldID))
		}).Scan(ctx, &items)
		if err != nil {
			return nil, err
		}
		for _, item := range items {
			result[gql.IntKey(item.ID)] = item.Count
		}
		return result, nil
	}
}

func (r AccountLoader) orgMemberCount() gql.LoaderFunc {
	return func(ctx context.Context, keys dataloader.Keys) (map[dataloader.Key]any, error) {
		var items []countItem
		err := r.client.Membership.Query().Where(membership.HasOrganizationWith(organization.IDIn(gql.ToInts(keys)...))).Modify(func(s *sql.Selector) {
			om := sql.Table(organization.MembershipsTable).As("or")
			s.Select(s.C(organization.FieldID), sql.As(sql.Count(om.C("role_id")), "count")).
				LeftJoin(om).On(s.C(organization.FieldID), om.C(organization.MembershipsColumn)).
				GroupBy(s.C(organization.FieldID))
		}).Scan(ctx, &items)
		if err != nil {
			return nil, err
		}
		return lo.SliceToMap(items, func(item countItem) (dataloader.Key, any) {
			return gql.IntKey(item.ID), item.Count
		}), nil
	}
}

func (r AccountLoader) orgChildrenCount() gql.LoaderFunc {
	return func(ctx context.Context, keys dataloader.Keys) (map[dataloader.Key]any, error) {
		var items []parentCountItem
		err := r.client.Organization.Query().Where(organization.ParentIDIn(gql.ToInts(keys)...)).Modify(func(s *sql.Selector) {
			s.Select(s.C(organization.FieldParentID), sql.As(sql.Count(s.C(organization.FieldID)), "count")).
				GroupBy(s.C(organization.FieldParentID))
		}).Scan(ctx, &items)
		if err != nil {
			return nil, err
		}
		return lo.SliceToMap(items, func(item parentCountItem) (dataloader.Key, any) {
			return gql.IntKey(item.ParentID), item.Count
		}), nil
	}
}
