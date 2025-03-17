package adminsrv

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/bucketheadv/infra-core/modules/utils"
	"grocery-store/constants/errcode"
	"grocery-store/database"
	"grocery-store/model/params"
	"grocery-store/service"
	"time"
)

const userTokenKey = "user:token:%s"

func AuthUser(param params.LoginParam) (string, error) {
	user, err := service.GetUserByUsername(param.Username)
	if err != nil {
		return "", err
	}
	password := md5.Sum([]byte(param.Password))
	if user.Password != string(password[:]) {
		return "", errcode.ErrInvalidNameOrPassword
	}
	var token = utils.Uuid(16)
	var key = fmt.Sprintf(userTokenKey, token)

	var redisClient = database.RedisClient
	redisClient.SetEX(context.Background(), key, user, 24*time.Hour)

	return token, nil
}
