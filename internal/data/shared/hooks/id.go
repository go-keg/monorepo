package hooks

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent"
	"github.com/sony/sonyflake"
)

func IDHook() ent.Hook {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{
		// 从 2025-01-01 00:00:00
		StartTime: time.Unix(1735660800, 0),
	})
	type IDSetter interface {
		SetID(int64)
		ID() (int64, bool)
	}
	return func(mutator ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
			if mutation.Op().Is(ent.OpCreate) {
				is, ok := mutation.(IDSetter)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation %T", mutation)
				}
				if _, ok = is.ID(); !ok {
					id, err := sf.NextID()
					if err != nil {
						return nil, err
					}
					is.SetID(int64(id))
				}
			}
			return mutator.Mutate(ctx, mutation)
		})
	}
}
