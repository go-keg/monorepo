package service

import (
	"context"
	v1 "github.com/go-keg/apis/api/account/v1"
	"github.com/go-keg/monorepo/internal/data/example/ent"
	"github.com/go-keg/monorepo/internal/pkg/auth"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AccountService struct {
	v1.UnimplementedAccountServiceServer
	ent *ent.Client
}

func NewAccountService(ent *ent.Client) *AccountService {
	return &AccountService{ent: ent}
}

func (r AccountService) Profile(ctx context.Context, _ *emptypb.Empty) (*v1.Account, error) {
	account, err := r.ent.Account.Get(ctx, auth.GetAccountID(ctx))
	if err != nil {
		return nil, err
	}
	return &v1.Account{
		Id:       int64(account.ID),
		Nickname: account.Nickname,
	}, nil
}
