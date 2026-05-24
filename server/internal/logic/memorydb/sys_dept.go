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

type sMemDBSysDept struct {
}

var instanceMemDBSysDept *sMemDBSysDept

func NewMemDBSysDept() *sMemDBSysDept {
	if instanceMemDBSysDept != nil {
		return instanceMemDBSysDept
	}
	instanceMemDBSysDept = &sMemDBSysDept{}
	initLoadMemoryDB := g.Cfg().MustGet(context.TODO(), "server.initLoadMemoryDB").Bool()
	if initLoadMemoryDB {
		instanceMemDBSysDept.Init(context.TODO())
	}
	return instanceMemDBSysDept
}

func init() {
	s := NewMemDBSysDept()
	service.RegisterMemDBSysDept(s)
	RegisterMemDB(s.TableName(context.TODO()), s)
}

func (s *sMemDBSysDept) Init(ctx context.Context) error {
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

func (s *sMemDBSysDept) EventHandler(ctx context.Context) error {
	eventList := []string{
		consts.EventKeyDBSysDeptCreate,
		consts.EventKeyDBSysDeptUpdate,
		consts.EventKeyDBSysDeptDelete,
	}
	event.EventsInstance().Register(eventList, func(ctx context.Context, eventKey string, args ...interface{}) {
		if len(args) < 1 {
			g.Log().Errorf(ctx, "sMemDBSysDept.EventHandler event args error. eventKey: %s, args: %+v", eventKey, args)
			return
		}
		pkList, ok := args[0].([]int64)
		if !ok {
			pk, ok := args[0].(int64)
			if !ok {
				g.Log().Errorf(ctx, "sMemDBSysDept.EventHandler event args error. eventKey: %s, args: %+v", eventKey, args)
				return
			}
			pkList = []int64{pk}
		}
		if len(pkList) == 0 {
			g.Log().Errorf(ctx, "sMemDBSysDept.EventHandler event args error. eventKey: %s, args: %+v", eventKey, args)
			return
		}
		if eventKey == consts.EventKeyDBSysDeptDelete {
			s.DeleteData(ctx, pkList)
		} else {
			for _, pk := range pkList {
				s.LoadData(ctx, pk)
			}
		}

		dbType := "edit"
		if eventKey == consts.EventKeyDBSysDeptCreate {
			dbType = "add"
		} else if eventKey == consts.EventKeyDBSysDeptDelete {
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

func (s *sMemDBSysDept) TableName(ctx context.Context) string {
	return dao.SysDept.Table()
}

func (s *sMemDBSysDept) CreateTable(ctx context.Context) error {
	_, err := db.Exec(ctx, "DROP TABLE IF EXISTS "+s.TableName(ctx))
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysDept.CreateTable delete table err: %v", err)
		return err
	}

	total, err := db.Model("sqlite_master").Where("type", "table").Where("name", dao.SysDept.Table()).Count()
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysDept.CreateTable count err: %v", err)
		return err
	}
	if total == 0 {
		g.Log().Debug(ctx, "sMemDBSysDept.List create table ", s.TableName(ctx))
		res, err := db.Exec(ctx, `
			CREATE TABLE "sys_dept" (
			"dept_id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"tenant_id" TEXT DEFAULT '000000' CHECK(length("tenant_id") <= 20),
			"parent_id" INTEGER DEFAULT 0,
			"ancestors" TEXT DEFAULT '' CHECK(length("ancestors") <= 500),
			"dept_name" TEXT DEFAULT '' CHECK(length("dept_name") <= 30),
			"dept_type" INTEGER DEFAULT 0,
			"dept_category" TEXT CHECK(length("dept_category") <= 100),
			"order_num" INTEGER DEFAULT 0,
			"leader" INTEGER,
			"phone" TEXT CHECK(length("phone") <= 11),
			"email" TEXT CHECK(length("email") <= 50),
			"status" TEXT DEFAULT '0' ,
			"created_dept" INTEGER,
			"created_by"          INTEGER NOT NULL,
			"created_at"          DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			"updated_by"          INTEGER,
			"updated_at"          DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			"deleted_by"          INTEGER,
			"deleted_at"          DATETIME
			);
		`)
		if err != nil {
			g.Log().Errorf(ctx, "sMemDBSysDept.CreateTable create table err: %v", err)
			return err
		}
		g.Log().Debug(ctx, "sMemDBSysDept.CreateTable create table success", res)
		// 加载数据
		s.LoadData(ctx, 0)
	}
	return nil
}

func (s *sMemDBSysDept) LoadData(ctx context.Context, pk int64) error {
	count := 0
	if pk == 0 {
		maxDeptId := int64(0)
		for {
			page := 1
			pageSize := 100
			dataList := make([]*entity.SysDept, 0)
			dao.SysDept.Ctx(ctx).Page(page, pageSize).WhereGT(dao.SysDept.Columns().DeptId, maxDeptId).
				OrderAsc(dao.SysDept.Columns().DeptId).Scan(&dataList)
			if len(dataList) == 0 {
				break
			}
			count += len(dataList)
			_, err := db.Model(dao.SysDept.Table()).Data(dataList).OnConflict(dao.SysDept.Columns().DeptId).Save()
			if err != nil {
				g.Log().Errorf(ctx, "sMemDBSysDept.Load data insert err: %v", err)
				return err
			}
			maxDeptId = dataList[len(dataList)-1].DeptId
			if len(dataList) < pageSize {
				break
			}
		}
	} else {
		dataList := make([]*entity.SysDept, 0)
		dao.SysDept.Ctx(ctx).Where(dao.SysDept.Columns().DeptId, pk).Scan(&dataList)
		_, err := db.Model(dao.SysDept.Table()).Data(dataList).OnConflict(dao.SysDept.Columns().DeptId).Save()
		if err != nil {
			g.Log().Errorf(ctx, "sMemDBSysDept.Load data insert err: %v", err)
		}
	}

	g.Log().Debugf(ctx, "sMemDBSysDept.LoadData load data success, total: %d, pk: %d", count, pk)
	return nil
}

func (s *sMemDBSysDept) DeleteData(ctx context.Context, ids []int64) error {
	_, err := db.Model(dao.SysDept.Table()).WhereIn(dao.SysDept.Columns().DeptId, ids).Unscoped().Delete()
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysDept.DeleteData err: %v", err)
		return err
	}
	g.Log().Debugf(ctx, "sMemDBSysDept.DeleteData delete data success, total: %d, ids: %+v", len(ids), ids)
	return nil
}
