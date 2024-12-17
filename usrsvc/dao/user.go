package dao

import (
	"context"
	"dogapm"
	"encoding/json"
	"fmt"
	"time"
)

type userDao struct {
}

var UserDao = &userDao{}

func (u *userDao) Get(ctx context.Context, uid int64) map[string]interface{} {
	userCache, err := dogapm.Infra.Rdb.Get(ctx, fmt.Sprintf("%s:%s:%d", "usersvc", "uinfo", uid)).Result()
	if len(userCache) != 0 {
		userInfo := make(map[string]interface{})
		err = json.Unmarshal([]byte(userCache), &userInfo)
		if err == nil {
			return userInfo
		}
	}
	userDbInfo := dogapm.DBUtil.QueryFirst(dogapm.Infra.Db.QueryContext(ctx, "select * from t_user where id = ?", uid))
	if len(userDbInfo) == 0 {
		dogapm.Logger.Info(ctx, "GetUserInfo", map[string]interface{}{
			"msg": "user not found",
			"uid": uid,
		})
		return nil
	}
	cacheUserStr, err := json.Marshal(userDbInfo)
	if err == nil {
		dogapm.Infra.Rdb.Set(ctx, fmt.Sprintf("%s:%s:%d", "usersvc", "uinfo", uid), cacheUserStr, 10*time.Minute)
	}
	return userDbInfo
}
