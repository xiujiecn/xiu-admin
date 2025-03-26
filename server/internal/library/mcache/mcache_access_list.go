package mcache

import (
	"context"
	"fmt"
	"time"
	"xiujieadmin/internal/consts"
	"xiujieadmin/internal/library/event"
	"xiujieadmin/internal/service"
)

func init() {
	event.EventsInstance().Register(consts.EventKeyUserLogout, func(ctx context.Context, eventKey string, args ...interface{}) {
		if len(args) == 0 {
			return
		}
		userId := args[0].(int64)
		RemoveUserAccessCodeList(ctx, userId)
	})

}

func GetUserAccessCodeList(ctx context.Context, userId int64) ([]string, error) {
	accessCodeList, err := Instance().Get(ctx, fmt.Sprintf(consts.MemCacheUserAccessCodeList, userId))
	if err != nil || accessCodeList == nil {
		accessCodeCurrList, err := service.SysAuth().GetUserAccessCodeList(ctx, userId)
		if err != nil {
			return nil, err
		}
		if len(accessCodeCurrList) == 0 {
			accessCodeCurrList = []string{"null"}
		}
		Instance().Set(ctx, fmt.Sprintf(consts.MemCacheUserAccessCodeList, userId), accessCodeCurrList, time.Hour*24)
		return accessCodeCurrList, nil
	}
	return accessCodeList.Strings(), nil
}

func SetUserAccessCodeList(ctx context.Context, userId int64, accessCodeList []string) error {
	return Instance().Set(ctx, fmt.Sprintf(consts.MemCacheUserAccessCodeList, userId), accessCodeList, time.Hour*24)
}

func RemoveUserAccessCodeList(ctx context.Context, userId int64) error {
	_, err := Instance().Remove(ctx, fmt.Sprintf(consts.MemCacheUserAccessCodeList, userId))
	return err
}
