package dataloader

import (
	"context"
	"net/http"

	"github.com/go-keg/keg/contrib/gql"
	"github.com/go-keg/monorepo/internal/data/account/ent"
	"github.com/graph-gophers/dataloader"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

// DataLoader offers data loaders scoped to a context
type DataLoader struct {
	tenantUserRoleCount     *dataloader.Loader
	permissionChildrenCount *dataloader.Loader
	orgChildrenCount        *dataloader.Loader
	orgMemberCount          *dataloader.Loader
}

// NewDataLoader returns the instantiated Loaders struct for use in a request
func NewDataLoader(client *ent.Client) *DataLoader {
	opts := []dataloader.Option{dataloader.WithCache(&dataloader.NoCache{})}
	account := AccountLoader{client: client}
	return &DataLoader{
		tenantUserRoleCount:     dataloader.NewBatchedLoader(gql.BatchFunc(account.tenantUserRoleCount()), opts...),
		permissionChildrenCount: dataloader.NewBatchedLoader(gql.BatchFunc(account.permissionChildrenCount()), opts...),
		orgChildrenCount:        dataloader.NewBatchedLoader(gql.BatchFunc(account.orgChildrenCount()), opts...),
		orgMemberCount:          dataloader.NewBatchedLoader(gql.BatchFunc(account.orgMemberCount()), opts...),
	}
}

func (l DataLoader) GetTenantUserRoleCount(ctx context.Context, tenantUserID int) (int, error) {
	thunk := l.tenantUserRoleCount.Load(ctx, gql.IntKey(tenantUserID))
	result, err := thunk()
	if err != nil || result == nil {
		return 0, err
	}
	return result.(int), nil
}

func (l DataLoader) GetPermissionChildrenCount(ctx context.Context, id int) (int, error) {
	thunk := l.permissionChildrenCount.Load(ctx, gql.IntKey(id))
	result, err := thunk()
	if err != nil || result == nil {
		return 0, err
	}
	return result.(int), nil
}

func (l DataLoader) GetOrgChildrenCount(ctx context.Context, id int) (int, error) {
	thunk := l.orgChildrenCount.Load(ctx, gql.IntKey(id))
	result, err := thunk()
	if err != nil || result == nil {
		return 0, err
	}
	return result.(int), nil
}

func (l DataLoader) GetOrgMemberCount(ctx context.Context, id int) (int, error) {
	thunk := l.orgMemberCount.Load(ctx, gql.IntKey(id))
	result, err := thunk()
	if err != nil || result == nil {
		return 0, err
	}
	return result.(int), nil
}

func Middleware(loader *DataLoader, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), loadersKey, loader)
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}

// For returns the dataloader for a given context
func For(ctx context.Context) *DataLoader {
	return ctx.Value(loadersKey).(*DataLoader)
}
