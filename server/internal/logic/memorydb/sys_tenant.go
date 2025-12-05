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

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sMemDBSysTenant struct {
}

var instanceMemDBSysTenant *sMemDBSysTenant

func NewMemDBSysTenant() *sMemDBSysTenant {
	if instanceMemDBSysTenant != nil {
		return instanceMemDBSysTenant
	}
	instanceMemDBSysTenant = &sMemDBSysTenant{}
	initLoadMemoryDB := g.Cfg().MustGet(context.TODO(), "server.initLoadMemoryDB").Bool()
	if initLoadMemoryDB {
		instanceMemDBSysTenant.Init(context.TODO())
	}
	return instanceMemDBSysTenant
}

func init() {
	s := NewMemDBSysTenant()
	service.RegisterMemDBSysTenant(s)
	RegisterMemDB(s.TableName(context.TODO()), s)
}

func (s *sMemDBSysTenant) Init(ctx context.Context) error {
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

func (s *sMemDBSysTenant) EventHandler(ctx context.Context) error {
	eventList := []string{
		consts.EventKeyDBSysTenantCreate,
		consts.EventKeyDBSysTenantUpdate,
		consts.EventKeyDBSysTenantDelete,
	}
	event.EventsInstance().Register(eventList, func(ctx context.Context, eventKey string, args ...interface{}) {
		if len(args) < 1 {
			g.Log().Errorf(ctx, "sMemDBSysTenant.EventHandler event args error. eventKey: %s, args: %+v", eventKey, args)
			return
		}
		pkList, ok := args[0].([]int64)
		if !ok {
			pk, ok := args[0].(int64)
			if !ok {
				g.Log().Errorf(ctx, "sMemDBSysTenant.EventHandler event args error. eventKey: %s, args: %+v", eventKey, args)
				return
			}
			pkList = []int64{pk}
		}
		if len(pkList) == 0 {
			g.Log().Errorf(ctx, "sMemDBSysTenant.EventHandler event args error. eventKey: %s, args: %+v", eventKey, args)
			return
		}
		if eventKey == consts.EventKeyDBSysTenantDelete {
			s.DeleteData(ctx, pkList)
		} else {
			for _, pk := range pkList {
				s.LoadData(ctx, pk)
			}
		}
		dbType := "edit"
		if eventKey == consts.EventKeyDBSysTenantCreate {
			dbType = "add"
		} else if eventKey == consts.EventKeyDBSysTenantDelete {
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

func (s *sMemDBSysTenant) DB(ctx context.Context) gdb.DB {
	return db
}

func (s *sMemDBSysTenant) TableName(ctx context.Context) string {
	return dao.SysTenant.Table()
}

func (s *sMemDBSysTenant) CreateTable(ctx context.Context) error {
	_, err := db.Exec(ctx, "DROP TABLE IF EXISTS "+s.TableName(ctx))
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysTenant.CreateTable delete table err: %v", err)
		return err
	}

	total, err := db.Model("sqlite_master").Where("type", "table").Where("name", dao.SysTenant.Table()).Count()
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysTenant.CreateTable count err: %v", err)
		return err
	}
	if total == 0 {
		g.Log().Debug(ctx, "sMemDBSysTenant.List create table ", s.TableName(ctx))
		res, err := db.Exec(ctx, `
			CREATE TABLE "sys_tenant" (
			"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"tenant_id" TEXT NOT NULL CHECK(length("tenant_id") <= 20),
			"contact_user_name" TEXT CHECK(length("contact_user_name") <= 20),
			"contact_phone" TEXT CHECK(length("contact_phone") <= 20),
			"company_name" TEXT CHECK(length("company_name") <= 50),
			"license_number" TEXT CHECK(length("license_number") <= 30),
			"address" TEXT CHECK(length("address") <= 200),
			"intro" TEXT CHECK(length("intro") <= 200),
			"domain" TEXT CHECK(length("domain") <= 200),
			"remark" TEXT CHECK(length("remark") <= 200),
			"package_id" INTEGER,
			"expire_time" TEXT,
			"account_count" INTEGER DEFAULT -1,
			"status" TEXT DEFAULT '0',
			"created_dept" INTEGER,
				"created_by"          INTEGER NOT NULL,
				"created_at"          DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				"updated_by"          INTEGER,
				"updated_at"          DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				"deleted_by"          INTEGER,
				"deleted_at"          DATETIME,
			CONSTRAINT "sys_tenant_tenant_id_unique" UNIQUE ("tenant_id")
			);
		`)
		if err != nil {
			g.Log().Errorf(ctx, "sMemDBSysTenant.CreateTable create table err: %v", err)
			return err
		}
		g.Log().Debug(ctx, "sMemDBSysTenant.CreateTable create table success", res)
		// 加载数据
		s.LoadData(ctx, 0)
	}
	return nil
}

func (s *sMemDBSysTenant) LoadData(ctx context.Context, pk int64) error {
	count := 0
	if pk == 0 {
		maxId := int64(0)
		for {
			page := 1
			pageSize := 500
			dataList := make([]*entity.SysTenant, 0)
			dao.SysTenant.Ctx(ctx).Page(page, pageSize).WhereGT(dao.SysTenant.Columns().Id, maxId).
				OrderAsc(dao.SysTenant.Columns().Id).Scan(&dataList)
			if len(dataList) == 0 {
				break
			}
			count += len(dataList)
			_, err := db.Model(dao.SysTenant.Table()).Data(dataList).OnConflict(dao.SysTenant.Columns().Id).Save()
			if err != nil {
				g.Log().Errorf(ctx, "sMemDBSysTenant.Load data insert err: %v, dataList: %+v", err, dataList)
				return err
			}
			maxId = dataList[len(dataList)-1].Id
			if len(dataList) < pageSize {
				break
			}
		}
	} else {
		dataList := make([]*entity.SysTenant, 0)
		dao.SysTenant.Ctx(ctx).Where(dao.SysTenant.Columns().Id, pk).Scan(&dataList)
		_, err := db.Model(dao.SysTenant.Table()).Data(dataList).OnConflict(dao.SysTenant.Columns().Id).Save()
		if err != nil {
			g.Log().Errorf(ctx, "sMemDBSysTenant.Load data insert err: %v", err)
		}
	}
	g.Log().Debugf(ctx, "sMemDBSysTenant.LoadData load data success, total: %d, pk: %d", count, pk)
	return nil
}

func (s *sMemDBSysTenant) DeleteData(ctx context.Context, ids []int64) error {
	_, err := db.Model(dao.SysTenant.Table()).WhereIn(dao.SysTenant.Columns().Id, ids).Delete()
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysTenant.DeleteData err: %v, ids: %+v", err, ids)
		return err
	}
	g.Log().Debugf(ctx, "sMemDBSysTenant.DeleteData delete data success, total: %d, ids: %+v", len(ids), ids)
	return nil
}
