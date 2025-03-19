package model

import (
	"xiujieadmin/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

type SysJobListModel struct {
	JobId          int64  `json:"jobId"          orm:"job_id"          description:"任务ID"`
	JobName        string `json:"jobName"        orm:"job_name"        description:"任务名称"`
	JobParams      string `json:"jobParams"      orm:"job_params"      description:"参数"`
	JobGroup       string `json:"jobGroup"       orm:"job_group"       description:"任务组名"`
	InvokeTarget   string `json:"invokeTarget"   orm:"invoke_target"   description:"调用目标字符串"`
	CronExpression string `json:"cronExpression" orm:"cron_expression" description:"cron执行表达式"`
	MisfirePolicy  int    `json:"misfirePolicy"  orm:"misfire_policy"  description:"计划执行策略（1多次执行 2执行一次）"`
	Concurrent     int    `json:"concurrent"     orm:"concurrent"      description:"是否并发执行（0允许 1禁止）"`
	Status         string `json:"status"         orm:"status"          description:"状态（0正常 1暂停）"`
	Remark         string `json:"remark"         orm:"remark"          description:"备注信息"`
	CreatedDept    int64  `json:"createdDept"    orm:"created_dept"    description:"创建部门"`
}

type SysJobViewModel struct {
	entity.SysJob
}

type SysJobAddModel struct {
	JobId          int64       `json:"jobId"          orm:"job_id"          description:"任务ID"`
	JobName        string      `v:"required" json:"jobName"        orm:"job_name"        description:"任务名称"`
	JobParams      string      `json:"jobParams"      orm:"job_params"      description:"参数"`
	JobGroup       string      `json:"jobGroup"       orm:"job_group"       description:"任务组名"`
	InvokeTarget   string      `v:"required" json:"invokeTarget"   orm:"invoke_target"   description:"调用目标字符串"`
	CronExpression string      `v:"required" json:"cronExpression" orm:"cron_expression" description:"cron执行表达式"`
	MisfirePolicy  int         `v:"required" json:"misfirePolicy"  orm:"misfire_policy"  description:"计划执行策略（1多次执行 2执行一次）"`
	Concurrent     int         `json:"concurrent"     orm:"concurrent"      description:"是否并发执行（0允许 1禁止）"`
	Status         string      `v:"required" json:"status"         orm:"status"          description:"状态（0正常 1暂停）"`
	Remark         string      `json:"remark"         orm:"remark"          description:"备注信息"`
	CreatedDept    int64       `json:"createdDept"    orm:"created_dept"    description:"创建部门"`
	CreatedBy      int64       `json:"createdBy"      orm:"created_by"      description:"创建者"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`
	UpdatedBy      int64       `json:"updatedBy"      orm:"updated_by"      description:"更新者"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:"更新时间"`
}

type SysJobUpdateModel struct {
	JobId          int64       `v:"required" json:"jobId"          orm:"job_id"          description:"任务ID"`
	JobName        string      `v:"required" json:"jobName"        orm:"job_name"        description:"任务名称"`
	JobParams      string      `json:"jobParams"      orm:"job_params"      description:"参数"`
	JobGroup       string      `json:"jobGroup"       orm:"job_group"       description:"任务组名"`
	InvokeTarget   string      `v:"required" json:"invokeTarget"   orm:"invoke_target"   description:"调用目标字符串"`
	CronExpression string      `v:"required" json:"cronExpression" orm:"cron_expression" description:"cron执行表达式"`
	MisfirePolicy  int         `v:"required" json:"misfirePolicy"  orm:"misfire_policy"  description:"计划执行策略（1多次执行 2执行一次）"`
	Concurrent     int         `json:"concurrent"     orm:"concurrent"      description:"是否并发执行（0允许 1禁止）"`
	Status         string      `v:"required" json:"status"         orm:"status"          description:"状态（0正常 1暂停）"`
	Remark         string      `json:"remark"         orm:"remark"          description:"备注信息"`
	UpdatedBy      int64       `json:"updatedBy"      orm:"updated_by"      description:"更新者"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:"更新时间"`
}

type SysJobDeleteModel struct {
	JobIds    []int64     `v:"required" json:"jobId"          orm:"job_id"          description:"任务ID"`
	DeletedBy int64       `json:"deletedBy"      orm:"deleted_by"      description:"删除人"`
	DeletedAt *gtime.Time `json:"deletedAt"      orm:"deleted_at"      description:"删除时间"`
}
type SysJobListParam struct {
	JobName  string `json:"jobName"        orm:"job_name"        description:"任务名称"`
	JobGroup string `json:"jobGroup"       orm:"job_group"       description:"任务组名"`
	Status   string `json:"status"         orm:"status"          description:"状态（0正常 1暂停）"`
}
