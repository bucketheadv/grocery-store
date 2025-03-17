package errcode

import "github.com/bucketheadv/infra-gin/api"

var (
	ErrInvalidNameOrPassword = api.NewBizError(10001, "用户名或密码不正确")
)
