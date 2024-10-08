package dataloader

import (
	"context"
	"github.com/go-keg/example/internal/data/example/ent"
	"github.com/go-keg/keg/contrib/gql"
	"github.com/graph-gophers/dataloader"
	"github.com/spf13/cast"
	"net/http"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

// DataLoader offers data loaders scoped to a context
type DataLoader struct {
	userRoleCountLoader           *dataloader.Loader
	permissionChildrenCountLoader *dataloader.Loader
}

// NewDataLoader returns the instantiated Loaders struct for use in a request
func NewDataLoader(client *ent.Client) *DataLoader {
	opts := []dataloader.Option{dataloader.WithCache(&dataloader.NoCache{})}
	admin := AdminLoader{client: client}
	return &DataLoader{
		userRoleCountLoader:           dataloader.NewBatchedLoader(gql.BatchFunc(admin.userRoleCount()), opts...),
		permissionChildrenCountLoader: dataloader.NewBatchedLoader(gql.BatchFunc(admin.permissionChildrenCount()), opts...),
	}
}

func (l DataLoader) GetUserRoleCount(ctx context.Context, id int) (int, error) {
	thunk := l.userRoleCountLoader.Load(ctx, dataloader.StringKey(cast.ToString(id)))
	result, err := thunk()
	if err != nil || result == nil {
		return 0, err
	}
	return result.(int), nil
}

func (l DataLoader) GetPermissionChildrenCount(ctx context.Context, id int) (int, error) {
	thunk := l.permissionChildrenCountLoader.Load(ctx, dataloader.StringKey(cast.ToString(id)))
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
