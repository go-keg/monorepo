package dataloader

import (
	"context"

	"github.com/go-keg/keg/contrib/gql"
	"github.com/graph-gophers/dataloader"
	"github.com/samber/lo"
)

type Loader struct {
}

func (r Loader) exampleCount() gql.LoaderFunc {
	return func(ctx context.Context, keys dataloader.Keys) (map[dataloader.Key]any, error) {
		i := 0
		return lo.SliceToMap(keys, func(item dataloader.Key) (dataloader.Key, any) {
			i++
			return item, i
		}), nil
	}
}
