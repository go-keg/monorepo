package auth

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport"
	"github.com/spf13/cast"
)

const (
	HeaderAccountID = "X-ACCOUNT-ID"
)

func GetAccountID(ctx context.Context) int {
	if tr, ok := transport.FromServerContext(ctx); ok {
		return cast.ToInt(tr.RequestHeader().Get(HeaderAccountID))
	}
	return 0
}
