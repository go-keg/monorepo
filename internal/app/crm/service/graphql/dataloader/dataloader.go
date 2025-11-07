package dataloader

import (
	"context"
	"net/http"

	"github.com/go-keg/keg/contrib/gql"
	"github.com/graph-gophers/dataloader"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

// DataLoader offers data loaders scoped to a context
type DataLoader struct {
	exampleCount *dataloader.Loader
}

// NewDataLoader returns the instantiated Loaders struct for use in a request
func NewDataLoader() *DataLoader {
	opts := []dataloader.Option{dataloader.WithCache(&dataloader.NoCache{})}
	loader := Loader{}
	return &DataLoader{
		exampleCount: dataloader.NewBatchedLoader(gql.BatchFunc(loader.exampleCount()), opts...),
	}
}

func (l DataLoader) GetExampleCount(ctx context.Context, id int) (int, error) {
	thunk := l.exampleCount.Load(ctx, gql.IntKey(id))
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
