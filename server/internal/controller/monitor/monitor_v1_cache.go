package monitor

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	v1 "server/api/monitor/v1"
	"server/utility"
)

func (c *ControllerV1) RedisInfo(ctx context.Context, req *v1.RedisInfoReq) (res *v1.RedisInfoRes, err error) {
	info, err := g.Redis().Do(ctx, "INFO")
	if err != nil {
		return nil, err
	}

	infoMap := make(map[string]string)
	for _, line := range strings.Split(info.String(), "\n") {
		if strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			infoMap[parts[0]] = parts[1]
		}
	}
	dbSize := gconv.Int(infoMap["db0"])

	info, err = g.Redis().Do(ctx, "INFO", "COMMANDSTATS")
	if err != nil {
		return nil, err
	}
	commandStats := make([]*v1.CommandStats, 0)
	for _, line := range strings.Split(info.String(), "\n") {
		if strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			if strings.HasPrefix(parts[0], "cmdstat_") {
				commandStats = append(commandStats, &v1.CommandStats{
					Name:  strings.TrimPrefix(parts[0], "cmdstat_"),
					Value: utility.GetBetweenStr(parts[1], "calls=", ",usec"),
				})
			}
		}
	}

	return &v1.RedisInfoRes{
		Info:         infoMap,
		DBSize:       dbSize,
		CommandStats: commandStats,
	}, nil
}
