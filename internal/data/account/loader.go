package account

import (
	"context"

	"github.com/go-keg/keg/contrib/gql"
	"github.com/go-keg/monorepo/internal/data/account/ent"
	"github.com/go-keg/monorepo/internal/data/account/ent/user"
	"github.com/graph-gophers/dataloader"
	"github.com/samber/lo"
)

type Loader struct {
	ent  *ent.Client
	user *dataloader.Loader
}

func NewLoader(ent *ent.Client, opts ...dataloader.Option) *Loader {
	loader := &Loader{ent: ent}
	loader.user = dataloader.NewBatchedLoader(gql.BatchFunc(loader.loadUser()), opts...)
	return loader
}

func (r Loader) loadUser() gql.LoaderFunc {
	return func(ctx context.Context, keys dataloader.Keys) (map[dataloader.Key]any, error) {
		items, err := r.ent.User.Query().Where(user.IDIn(gql.ToInts(keys)...)).All(ctx)
		if err != nil {
			return nil, err
		}
		return lo.SliceToMap(items, func(item *ent.User) (dataloader.Key, any) {
			return gql.IntKey(item.ID), item
		}), nil
	}
}

func (r Loader) GetUser(ctx context.Context, id int) (*ent.User, error) {
	thunk := r.user.Load(ctx, gql.IntKey(id))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	if result == nil {
		return &ent.User{}, nil
	}
	return result.(*ent.User), nil
}

func (r Loader) GetUsers(ctx context.Context, ids []int) ([]*ent.User, error) {
	thunk := r.user.LoadMany(ctx, lo.Map(ids, func(item int, _ int) dataloader.Key {
		return gql.IntKey(item)
	}))
	return gql.LoadManyResult[*ent.User](thunk())
}
