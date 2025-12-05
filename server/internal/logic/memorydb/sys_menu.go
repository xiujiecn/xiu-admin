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

type sMemDBSysMenu struct {
}

var instanceMemDBSysMenu *sMemDBSysMenu

func NewMemDBSysMenu() *sMemDBSysMenu {
	if instanceMemDBSysMenu != nil {
		return instanceMemDBSysMenu
	}
	instanceMemDBSysMenu = &sMemDBSysMenu{}
	initLoadMemoryDB := g.Cfg().MustGet(context.TODO(), "server.initLoadMemoryDB").Bool()
	if initLoadMemoryDB {
		instanceMemDBSysMenu.Init(context.TODO())
	}
	return instanceMemDBSysMenu
}

func init() {
	s := NewMemDBSysMenu()
	service.RegisterMemDBSysMenu(s)
	RegisterMemDB(s.TableName(context.TODO()), s)
}

func (s *sMemDBSysMenu) Init(ctx context.Context) error {
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

func (s *sMemDBSysMenu) EventHandler(ctx context.Context) error {
	eventList := []string{
		consts.EventKeyDBSysMenuCreate,
		consts.EventKeyDBSysMenuUpdate,
		consts.EventKeyDBSysMenuDelete,
	}
	event.EventsInstance().Register(eventList, func(ctx context.Context, eventKey string, args ...interface{}) {
		if len(args) < 1 {
			g.Log().Errorf(ctx, "sMemDBSysMenu.EventHandler event args error. eventKey: %s, args: %+v", eventKey, args)
			return
		}
		pkList, ok := args[0].([]int64)
		if !ok {
			pk, ok := args[0].(int64)
			if !ok {
				g.Log().Errorf(ctx, "sMemDBSysMenu.EventHandler event args error. eventKey: %s, args: %+v", eventKey, args)
				return
			}
			pkList = []int64{pk}
		}
		if len(pkList) == 0 {
			g.Log().Errorf(ctx, "sMemDBSysMenu.EventHandler event args error. eventKey: %s, args: %+v", eventKey, args)
			return
		}
		if eventKey == consts.EventKeyDBSysMenuDelete {
			s.DeleteData(ctx, pkList)
		} else {
			for _, pk := range pkList {
				s.LoadData(ctx, pk)
			}
		}
		dbType := "edit"
		if eventKey == consts.EventKeyDBSysMenuCreate {
			dbType = "add"
		} else if eventKey == consts.EventKeyDBSysMenuDelete {
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

func (s *sMemDBSysMenu) TableName(ctx context.Context) string {
	return dao.SysMenu.Table()
}

func (s *sMemDBSysMenu) CreateTable(ctx context.Context) error {
	_, err := db.Exec(ctx, "DROP TABLE IF EXISTS "+s.TableName(ctx))
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysMenu.CreateTable delete table err: %v", err)
		return err
	}

	total, err := db.Model("sqlite_master").Where("type", "table").Where("name", dao.SysMenu.Table()).Count()
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysMenu.CreateTable count err: %v", err)
		return err
	}
	if total == 0 {
		g.Log().Debug(ctx, "sMemDBSysMenu.List create table ", s.TableName(ctx))
		res, err := db.Exec(ctx, `
			CREATE TABLE "sys_menu" (
				"menu_id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				"menu_name" TEXT NOT NULL CHECK(length("menu_name") <= 50),
				"parent_id" INTEGER DEFAULT 0,
				"level" INTEGER DEFAULT 1,
				"tree" TEXT CHECK(length("tree") <= 255),
				"order_num" INTEGER DEFAULT 0,
				"path" TEXT DEFAULT '' CHECK(length("path") <= 200),
				"component" TEXT CHECK(length("component") <= 255),
				"query_param" TEXT CHECK(length("query_param") <= 255),
				"is_frame" INTEGER DEFAULT 1 ,
				"is_cache" INTEGER DEFAULT 0 ,
				"menu_type" TEXT DEFAULT '' ,
				"visible" TEXT DEFAULT '0',
				"status" TEXT DEFAULT '0' ,
				"perms" TEXT CHECK(length("perms") <= 100),
				"icon" TEXT DEFAULT '#' CHECK(length("icon") <= 100),
				"created_dept" INTEGER,
				"created_by" INTEGER,
				"created_at" TEXT,
				"updated_by" INTEGER,
				"updated_at" TEXT,
				"remark" TEXT DEFAULT '' CHECK(length("remark") <= 500)
				);
		`)
		if err != nil {
			g.Log().Errorf(ctx, "sMemDBSysMenu.CreateTable create table err: %v", err)
			return err
		}
		g.Log().Debug(ctx, "sMemDBSysMenu.CreateTable create table success", res)
		s.LoadData(ctx, 0)
		return nil
	}
	return nil
}

func (s *sMemDBSysMenu) LoadData(ctx context.Context, pk int64) error {
	count := 0
	if pk == 0 {
		maxMenuId := int64(0)
		for {
			page := 1
			pageSize := 500
			dataList := make([]*entity.SysMenu, 0)
			dao.SysMenu.Ctx(ctx).Page(page, pageSize).WhereGT(dao.SysMenu.Columns().MenuId, maxMenuId).
				OrderAsc(dao.SysMenu.Columns().MenuId).Scan(&dataList)
			if len(dataList) == 0 {
				break
			}
			count += len(dataList)
			_, err := db.Model(dao.SysMenu.Table()).Data(dataList).OnConflict(dao.SysMenu.Columns().MenuId).Save()
			if err != nil {
				g.Log().Errorf(ctx, "sMemDBSysMenu.Load data insert err: %v", err)
				return err
			}
			maxMenuId = dataList[len(dataList)-1].MenuId
			if len(dataList) < pageSize {
				break
			}
		}
	} else {
		dataList := make([]*entity.SysMenu, 0)
		dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().MenuId, pk).Scan(&dataList)
		_, err := db.Model(dao.SysMenu.Table()).Data(dataList).OnConflict(dao.SysMenu.Columns().MenuId).Save()
		if err != nil {
			g.Log().Errorf(ctx, "sMemDBSysMenu.Load data insert err: %v", err)
		}
	}
	g.Log().Debugf(ctx, "sMemDBSysMenu.LoadData load data success, total: %d, pk: %d", count, pk)
	return nil
}

func (s *sMemDBSysMenu) DeleteData(ctx context.Context, ids []int64) error {
	_, err := db.Model(dao.SysMenu.Table()).WhereIn(dao.SysMenu.Columns().MenuId, ids).Delete()
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysMenu.DeleteData err: %v", err)
		return err
	}
	g.Log().Debugf(ctx, "sMemDBSysMenu.DeleteData delete data success, total: %d, ids: %+v", len(ids), ids)
	return nil
}
