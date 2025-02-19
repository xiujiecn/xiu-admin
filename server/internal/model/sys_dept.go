package model

type DeptListQuery struct {
	DeptName string `json:"deptName"`
	Status   string `json:"status"`
}

type SysDept struct {
	DeptId   int64  `json:"deptId"`
	DeptName string `json:"deptName"`
	Status   string `json:"status"`
}
