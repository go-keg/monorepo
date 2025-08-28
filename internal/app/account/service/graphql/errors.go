package graphql

import "github.com/go-keg/keg/contrib/gql"

// region Account

// ErrAccountOrPasswordInvalid 账号或密码无效
var ErrAccountOrPasswordInvalid = gql.Error("account or password invalid", gql.WithErrCode("AccountOrPasswordInvalid"))

// ErrVerifyCodeInvalid 验证码无效
var ErrVerifyCodeInvalid = gql.Error("verify code invalid", gql.WithErrCode("VerifyCodeInvalid"))

// endregion
