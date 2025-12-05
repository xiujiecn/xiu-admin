// package hook
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package hook

import (
	"context"
	"xiuadmin/utility"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type MemberSumma struct {
	UserId   int64  `json:"userId"   orm:"user_id"   description:"用户ID"`
	UserName string `json:"userName" orm:"user_name" description:"用户账号"`
	NickName string `json:"nickName" orm:"nick_name" description:"用户昵称"`
	Avatar   string `json:"avatar"   orm:"avatar"    description:"头像地址"`
	TenantId string `json:"tenantId"    orm:"tenant_id"    description:"租户编号"`
	DeptId   int64  `json:"deptId"      orm:"dept_id"      description:"部门ID"`
}

// MemberSummary 操作人摘要信息
var MemberSummary = gdb.HookHandler{
	Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
		result, err = in.Next(ctx)
		if err != nil {
			return
		}

		var (
			createdByIds []int64
			updatedByIds []int64
			deletedByIds []int64
			memberIds    []int64
		)

		for _, record := range result {
			if record["created_by"].Int64() > 0 {
				createdByIds = append(createdByIds, record["created_by"].Int64())
			}
			if record["updated_by"].Int64() > 0 {
				updatedByIds = append(updatedByIds, record["updated_by"].Int64())
			}
			if record["deleted_by"].Int64() > 0 {
				deletedByIds = append(deletedByIds, record["deleted_by"].Int64())
			}
			if record["user_id"].Int64() > 0 {
				memberIds = append(memberIds, record["user_id"].Int64())
			}
		}

		memberIds = append(memberIds, createdByIds...)
		memberIds = append(memberIds, updatedByIds...)
		memberIds = append(memberIds, deletedByIds...)
		memberIds = utility.UniqueSlice(memberIds)
		if len(memberIds) == 0 {
			return
		}

		var members []*MemberSumma
		if err = g.Model("sys_user").Ctx(ctx).WhereIn("user_id", memberIds).Scan(&members); err != nil {
			g.Log().Errorf(ctx, "[MemberSummary] 查询用户信息失败: %v", err)
			return nil, err
		}

		if len(members) == 0 {
			g.Log().Warningf(ctx, "[MemberSummary] 没有找到任何用户信息，用户IDs: %v", memberIds)
			return
		}

		findMember := func(id *gvar.Var) *MemberSumma {
			userId := id.Int64()
			for _, v := range members {
				if v.UserId == userId {
					return v
				}
			}
			// 只在找不到用户时记录警告（避免日志过多）
			g.Log().Warningf(ctx, "[MemberSummary] 未找到用户ID: %d", userId)
			return nil
		}

		for _, record := range result {
			if record["created_by"].Int64() > 0 {
				member := findMember(record["created_by"])
				if member != nil {
					record["createdBySumma"] = gvar.New(member)
				}
			}

			if record["updated_by"].Int64() > 0 {
				member := findMember(record["updated_by"])
				if member != nil {
					record["updatedBySumma"] = gvar.New(member)
				}
			}

			if record["deleted_by"].Int64() > 0 {
				member := findMember(record["deleted_by"])
				if member != nil {
					record["deletedBySumma"] = gvar.New(member)
				}
			}

			if record["user_id"].Int64() > 0 {
				member := findMember(record["user_id"])
				if member != nil {
					record["memberBySumma"] = gvar.New(member)
				}
			}
		}
		return
	},
}
