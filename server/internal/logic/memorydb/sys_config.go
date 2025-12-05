package memorydb

import (
	"context"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/dao"
	"xiuadmin/internal/library/event"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/entity"
	"xiuadmin/internal/service"

	"xiuadmin/utility/queue"

	"github.com/gogf/gf/v2/frame/g"
)

type sMemDBSysConfig struct {
}

var instanceMemDBSysConfig *sMemDBSysConfig

func NewMemDBSysConfig() *sMemDBSysConfig {
	if instanceMemDBSysConfig != nil {
		return instanceMemDBSysConfig
	}
	instanceMemDBSysConfig = &sMemDBSysConfig{}
	initLoadMemoryDB := g.Cfg().MustGet(context.TODO(), "server.initLoadMemoryDB").Bool()
	if initLoadMemoryDB {
		instanceMemDBSysConfig.Init(context.TODO())
	}
	return instanceMemDBSysConfig
}

func init() {
	s := NewMemDBSysConfig()
	service.RegisterMemDBSysConfig(s)
	RegisterMemDB(s.TableName(context.TODO()), s)
}

func (s *sMemDBSysConfig) Init(ctx context.Context) error {
	err := s.CreateTable(ctx)
	if err != nil {
		return err
	}
	err = s.EventHandler(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *sMemDBSysConfig) EventHandler(ctx context.Context) error {
	eventList := []string{
		consts.EventKeyDBSysConfigCreate,
		consts.EventKeyDBSysConfigUpdate,
		consts.EventKeyDBSysConfigDelete,
	}
	event.EventsInstance().Register(eventList, func(ctx context.Context, eventKey string, args ...interface{}) {
		if len(args) < 1 {
			g.Log().Errorf(ctx, "memorydb/sys_config.go event args error. eventKey: %s, args: %+v", eventKey, args)
			return
		}
		pkList, ok := args[0].([]int64)
		if !ok {
			pk, ok := args[0].(int64)
			if !ok {
				g.Log().Errorf(ctx, "sMemDBSysConfig.EventHandler event args error. eventKey: %s, args: %+v", eventKey, args)
				return
			}
			pkList = []int64{pk}
		}
		if len(pkList) == 0 {
			g.Log().Errorf(ctx, "sMemDBSysConfig.EventHandler event args error. eventKey: %s, args: %+v", eventKey, args)
			return
		}
		if eventKey == consts.EventKeyDBSysConfigDelete {
			s.DeleteData(ctx, pkList)
		} else {
			for _, pk := range pkList {
				s.LoadData(ctx, pk)
			}
		}

		dbType := "edit"
		if eventKey == consts.EventKeyDBSysConfigCreate {
			dbType = "add"
		} else if eventKey == consts.EventKeyDBSysConfigDelete {
			dbType = "del"
		} else {
			dbType = "edit"
		}
		queue.Push(ctx, consts.BroadcastDbChg, &model.BroadcastDbChgCache{
			TableName: s.TableName(ctx),
			PK:        pkList,
			Type:      dbType,
			Data:      nil,
		})
	})
	return nil
}

func (s *sMemDBSysConfig) TableName(ctx context.Context) string {
	return dao.SysConfig.Table()
}

func (s *sMemDBSysConfig) CreateTable(ctx context.Context) error {
	_, err := db.Exec(ctx, "DROP TABLE IF EXISTS "+s.TableName(ctx))
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysConfig.CreateTable delete table err: %v", err)
		return err
	}

	total, err := db.Model("sqlite_master").Where("type", "table").Where("name", dao.SysConfig.Table()).Count()
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysConfig.CreateTable count err: %v", err)
		return err
	}
	if total == 0 {
		g.Log().Debug(ctx, "sMemDBSysConfig.List create table ", s.TableName(ctx))
		res, err := db.Exec(ctx, `
			CREATE TABLE "sys_config" (
				"config_id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				"tenant_id" TEXT DEFAULT '000000' CHECK(length("tenant_id") <= 20),
				"config_name" TEXT DEFAULT '' CHECK(length("config_name") <= 100),
				"config_key" TEXT DEFAULT '' CHECK(length("config_key") <= 100),
				"config_value" TEXT DEFAULT '' CHECK(length("config_value") <= 1024),
				"config_type" TEXT DEFAULT 'N',
				"created_dept" INTEGER,
				"created_by" INTEGER,
				"created_at" TEXT,
				"updated_by" INTEGER,
				"updated_at" TEXT,
				"remark" TEXT CHECK(length("remark") <= 500)
				);
		`)
		if err != nil {
			g.Log().Errorf(ctx, "sMemDBSysConfig.CreateTable create table err: %v", err)
			return err
		}
		g.Log().Debug(ctx, "sMemDBSysConfig.CreateTable create table success", res)
		// 加载数据
		s.LoadData(ctx, 0)
	}
	return nil
}

func (s *sMemDBSysConfig) LoadData(ctx context.Context, pk int64) error {
	count := 0
	if pk == 0 {
		maxConfigId := int64(0)
		for {
			page := 1
			pageSize := 500
			dataList := make([]*entity.SysConfig, 0)
			dao.SysConfig.Ctx(ctx).Page(page, pageSize).WhereGT(dao.SysConfig.Columns().ConfigId, maxConfigId).
				OrderAsc(dao.SysConfig.Columns().ConfigId).Scan(&dataList)
			if len(dataList) == 0 {
				break
			}
			count += len(dataList)
			_, err := db.Model(dao.SysConfig.Table()).Data(dataList).OnConflict(dao.SysConfig.Columns().ConfigId).Save()
			if err != nil {
				g.Log().Errorf(ctx, "sMemDBSysConfig.Load data insert err: %v", err)
				return err
			}
			maxConfigId = dataList[len(dataList)-1].ConfigId
			if len(dataList) < pageSize {
				break
			}
		}
	} else {
		dataList := make([]*entity.SysConfig, 0)
		dao.SysConfig.Ctx(ctx).Where(dao.SysConfig.Columns().ConfigId, pk).Scan(&dataList)
		_, err := db.Model(dao.SysConfig.Table()).Data(dataList).OnConflict(dao.SysConfig.Columns().ConfigId).Save()
		if err != nil {
			g.Log().Errorf(ctx, "sMemDBSysConfig.Load data insert err: %v", err)
			return err
		}
	}

	g.Log().Debugf(ctx, "sMemDBSysConfig.LoadData load data success, total: %d, pk: %d", count, pk)
	return nil
}

func (s *sMemDBSysConfig) DeleteData(ctx context.Context, ids []int64) error {
	_, err := db.Model(dao.SysConfig.Table()).WhereIn(dao.SysConfig.Columns().ConfigId, ids).Delete()
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysConfig.DeleteData err: %v", err)
		return err
	}
	g.Log().Debugf(ctx, "sMemDBSysConfig.DeleteData delete data success, total: %d, ids: %+v", len(ids), ids)
	return nil
}
