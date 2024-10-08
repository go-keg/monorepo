package service

import (
	"github.com/go-keg/example/internal/data/example/ent"
)

type AccountService struct {
	//v1.UnimplementedAccountServiceServer
	ent *ent.Client
}

func NewAccountService(ent *ent.Client) *AccountService {
	return &AccountService{ent: ent}
}

//
//func (r AccountService) Profile(ctx context.Context, req *emptypb.Empty) (*v1.Account, error) {
//	account, err := r.ent.Account.Get(ctx, auth.GetAccountID(ctx))
//	if err != nil {
//		return nil, err
//	}
//	return &v1.Account{
//		Id:       int64(account.ID),
//		Nickname: account.Nickname,
//	}, nil
//}
