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

type sMemDBSysRole struct {
}

var instanceMemDBSysRole *sMemDBSysRole

func NewMemDBSysRole() *sMemDBSysRole {
	if instanceMemDBSysRole != nil {
		return instanceMemDBSysRole
	}
	instanceMemDBSysRole = &sMemDBSysRole{}
	initLoadMemoryDB := g.Cfg().MustGet(context.TODO(), "server.initLoadMemoryDB").Bool()
	if initLoadMemoryDB {
		instanceMemDBSysRole.Init(context.TODO())
	}
	return instanceMemDBSysRole
}

func init() {
	s := NewMemDBSysRole()
	service.RegisterMemDBSysRole(s)
	RegisterMemDB(s.TableName(context.TODO()), s)
}

func (s *sMemDBSysRole) Init(ctx context.Context) error {
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

func (s *sMemDBSysRole) EventHandler(ctx context.Context) error {
	eventList := []string{
		consts.EventKeyDBSysRoleCreate,
		consts.EventKeyDBSysRoleUpdate,
		consts.EventKeyDBSysRoleDelete,
	}
	event.EventsInstance().Register(eventList, func(ctx context.Context, eventKey string, args ...interface{}) {
		if len(args) < 1 {
			g.Log().Errorf(ctx, "sMemDBSysRole.EventHandler event args error. eventKey: %s, args: %+v", eventKey, args)
			return
		}
		pkList, ok := args[0].([]int64)
		if !ok {
			pk, ok := args[0].(int64)
			if !ok {
				g.Log().Errorf(ctx, "sMemDBSysRole.EventHandler event args error. eventKey: %s, args: %+v", eventKey, args)
				return
			}
			pkList = []int64{pk}
		}
		if len(pkList) == 0 {
			g.Log().Errorf(ctx, "sMemDBSysRole.EventHandler event args error. eventKey: %s, args: %+v", eventKey, args)
			return
		}
		if eventKey == consts.EventKeyDBSysRoleDelete {
			s.DeleteData(ctx, pkList)
		} else {
			for _, pk := range pkList {
				s.LoadData(ctx, pk)
			}
		}
		dbType := "edit"
		if eventKey == consts.EventKeyDBSysRoleCreate {
			dbType = "add"
		} else if eventKey == consts.EventKeyDBSysRoleDelete {
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

func (s *sMemDBSysRole) TableName(ctx context.Context) string {
	return dao.SysRole.Table()
}

func (s *sMemDBSysRole) CreateTable(ctx context.Context) error {
	_, err := db.Exec(ctx, "DROP TABLE IF EXISTS "+s.TableName(ctx))
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysRole.CreateTable delete table err: %v", err)
		return err
	}
	total, err := db.Model("sqlite_master").Where("type", "table").Where("name", dao.SysRole.Table()).Count()
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysRole.CreateTable count err: %v", err)
		return err
	}
	if total == 0 {
		g.Log().Debug(ctx, "sMemDBSysRole.List create table ", s.TableName(ctx))
		res, err := db.Exec(ctx, `
			CREATE TABLE "sys_role" (
				"role_id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				"tenant_id" TEXT DEFAULT '000000' CHECK(length("tenant_id") <= 20),
				"dept_id" INTEGER DEFAULT 0,
				"role_name" TEXT NOT NULL CHECK(length("role_name") <= 30),
				"role_key" TEXT NOT NULL CHECK(length("role_key") <= 100),
				"role_sort" INTEGER NOT NULL,
				"data_scope" TEXT DEFAULT '1',
				"menu_check_strictly" INTEGER DEFAULT 1,
				"dept_check_strictly" INTEGER DEFAULT 1,
				"status" TEXT NOT NULL ,
				"created_dept" INTEGER,
				"created_by"          INTEGER NOT NULL,
				"created_at"          DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				"updated_by"          INTEGER,
				"updated_at"          DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				"deleted_by"          INTEGER,
				"deleted_at"          DATETIME,
				"remark" TEXT CHECK(length("remark") <= 500)
			);
		`)
		if err != nil {
			g.Log().Errorf(ctx, "sMemDBSysRole.CreateTable create table err: %v", err)
			return err
		}
		g.Log().Debug(ctx, "sMemDBSysRole.CreateTable create table success", res)
		s.LoadData(ctx, 0)
		return nil
	}
	return nil
}

func (s *sMemDBSysRole) LoadData(ctx context.Context, pk int64) error {
	count := 0
	if pk == 0 {
		maxRoleId := int64(0)
		for {
			page := 1
			pageSize := 500
			dataList := make([]*entity.SysRole, 0)
			dao.SysRole.Ctx(ctx).Page(page, pageSize).WhereGT(dao.SysRole.Columns().RoleId, maxRoleId).
				OrderAsc(dao.SysRole.Columns().RoleId).Scan(&dataList)
			if len(dataList) == 0 {
				break
			}
			count += len(dataList)
			_, err := db.Model(dao.SysRole.Table()).Data(dataList).OnConflict(dao.SysRole.Columns().RoleId).Save()
			if err != nil {
				g.Log().Errorf(ctx, "sMemDBSysRole.Load data insert err: %v", err)
				return err
			}
			maxRoleId = dataList[len(dataList)-1].RoleId
			if len(dataList) < pageSize {
				break
			}
		}
	} else {
		dataList := make([]*entity.SysRole, 0)
		dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().RoleId, pk).Scan(&dataList)
		_, err := db.Model(dao.SysRole.Table()).Data(dataList).OnConflict(dao.SysRole.Columns().RoleId).Save()
		if err != nil {
			g.Log().Errorf(ctx, "sMemDBSysRole.Load data insert err: %v", err)
		}
	}
	g.Log().Debugf(ctx, "sMemDBSysRole.LoadData load data success, total: %d, pk: %d", count, pk)
	return nil
}

func (s *sMemDBSysRole) DeleteData(ctx context.Context, ids []int64) error {
	g.Log().Debugf(ctx, "sMemDBSysRole.DeleteData called with ids: %+v, len: %d", ids, len(ids))

	if len(ids) == 0 {
		g.Log().Warningf(ctx, "sMemDBSysRole.DeleteData called with empty ids")
		return nil
	}

	// 使用 Unscoped() 来确保删除操作被执行
	_, err := db.Model(dao.SysRole.Table()).WhereIn(dao.SysRole.Columns().RoleId, ids).Unscoped().Delete()
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysRole.DeleteData err: %v", err)
		return err
	}
	g.Log().Debugf(ctx, "sMemDBSysRole.DeleteData delete data success, total: %d, ids: %+v", len(ids), ids)
	return nil
}
