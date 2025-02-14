package system

import (
	"context"

	v1 "server/api/system/v1"
)

func (c *ControllerV1) MenuAll(ctx context.Context, req *v1.MenuAllReq) (res *v1.MenuAllRes, err error) {
	res = &[]v1.RouteMenu{
		{
			Path:      "/dashboard",
			Name:      "Dashboard",
			Component: "/dashboard/index",
			Meta: &v1.RouteMeta{
				Order: -1,
				Title: "概览",
				Icon:  "lucide:layout-dashboard",
			},
			Children: []v1.RouteMenu{
				{
					Path:      "/analytics",
					Name:      "Analytics",
					Component: "/dashboard/analytics/index",
					Meta: &v1.RouteMeta{
						Title: "分析页",
						Icon:  "lucide:area-chart",
					},
				},
				{
					Path:      "/workspace",
					Name:      "Workspace",
					Component: "/dashboard/workspace/index",
					Meta: &v1.RouteMeta{
						Title: "工作台",
						Icon:  "carbon:workspace",
					},
				},
			},
		},
	}
	return
}
