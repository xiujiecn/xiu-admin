// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package testtree

import (
	"context"

	"xiuadmin/api/gen/testtree/v1"
)

type ITesttreeV1 interface {
	TestTreeList(ctx context.Context, req *v1.TestTreeListReq) (res *v1.TestTreeListRes, err error)
	TestTreeExport(ctx context.Context, req *v1.TestTreeExportReq) (res *v1.TestTreeExportRes, err error)
	TestTreeView(ctx context.Context, req *v1.TestTreeViewReq) (res *v1.TestTreeViewRes, err error)
	TestTreeEdit(ctx context.Context, req *v1.TestTreeEditReq) (res *v1.TestTreeEditRes, err error)
	TestTreeDelete(ctx context.Context, req *v1.TestTreeDeleteReq) (res *v1.TestTreeDeleteRes, err error)
	TestTreeTreeOption(ctx context.Context, req *v1.TestTreeTreeOptionReq) (res *v1.TestTreeTreeOptionRes, err error)
}
