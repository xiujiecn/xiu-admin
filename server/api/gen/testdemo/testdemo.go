// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package testdemo

import (
	"context"

	"xiujieadmin/api/gen/testdemo/v1"
)

type ITestdemoV1 interface {
	TestDemoList(ctx context.Context, req *v1.TestDemoListReq) (res *v1.TestDemoListRes, err error)
	TestDemoExport(ctx context.Context, req *v1.TestDemoExportReq) (res *v1.TestDemoExportRes, err error)
	TestDemoView(ctx context.Context, req *v1.TestDemoViewReq) (res *v1.TestDemoViewRes, err error)
	TestDemoEdit(ctx context.Context, req *v1.TestDemoEditReq) (res *v1.TestDemoEditRes, err error)
	TestDemoDelete(ctx context.Context, req *v1.TestDemoDeleteReq) (res *v1.TestDemoDeleteRes, err error)
}
