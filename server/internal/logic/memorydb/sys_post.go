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

type sMemDBSysPost struct {
}

var instanceMemDBSysPost *sMemDBSysPost

func NewMemDBSysPost() *sMemDBSysPost {
	if instanceMemDBSysPost != nil {
		return instanceMemDBSysPost
	}
	instanceMemDBSysPost = &sMemDBSysPost{}
	initLoadMemoryDB := g.Cfg().MustGet(context.TODO(), "server.initLoadMemoryDB").Bool()
	if initLoadMemoryDB {
		instanceMemDBSysPost.Init(context.TODO())
	}
	return instanceMemDBSysPost
}

func init() {
	s := NewMemDBSysPost()
	service.RegisterMemDBSysPost(s)
	RegisterMemDB(s.TableName(context.TODO()), s)
}

func (s *sMemDBSysPost) Init(ctx context.Context) error {
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

func (s *sMemDBSysPost) EventHandler(ctx context.Context) error {
	eventList := []string{
		consts.EventKeyDBSysPostCreate,
		consts.EventKeyDBSysPostUpdate,
		consts.EventKeyDBSysPostDelete,
	}
	event.EventsInstance().Register(eventList, func(ctx context.Context, eventKey string, args ...interface{}) {
		if len(args) < 1 {
			g.Log().Errorf(ctx, "sMemDBSysPost.EventHandler event args error. eventKey: %s, args: %+v", eventKey, args)
			return
		}
		pkList, ok := args[0].([]int64)
		if !ok {
			pk, ok := args[0].(int64)
			if !ok {
				g.Log().Errorf(ctx, "sMemDBSysPost.EventHandler event args error. eventKey: %s, args: %+v", eventKey, args)
				return
			}
			pkList = []int64{pk}
		}
		if len(pkList) == 0 {
			g.Log().Errorf(ctx, "sMemDBSysPost.EventHandler event args error. eventKey: %s, args: %+v", eventKey, args)
			return
		}
		if eventKey == consts.EventKeyDBSysPostDelete {
			s.DeleteData(ctx, pkList)
		} else {
			for _, pk := range pkList {
				s.LoadData(ctx, pk)
			}
		}
		dbType := "edit"
		if eventKey == consts.EventKeyDBSysPostCreate {
			dbType = "add"
		} else if eventKey == consts.EventKeyDBSysPostDelete {
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

func (s *sMemDBSysPost) TableName(ctx context.Context) string {
	return dao.SysPost.Table()
}

func (s *sMemDBSysPost) CreateTable(ctx context.Context) error {
	_, err := db.Exec(ctx, "DROP TABLE IF EXISTS "+s.TableName(ctx))
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysPost.CreateTable delete table err: %v", err)
		return err
	}

	total, err := db.Model("sqlite_master").Where("type", "table").Where("name", dao.SysPost.Table()).Count()
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysPost.CreateTable count err: %v", err)
		return err
	}
	if total == 0 {
		g.Log().Debug(ctx, "sMemDBSysPost.List create table ", s.TableName(ctx))
		res, err := db.Exec(ctx, `
			CREATE TABLE "sys_post" (
			"post_id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"tenant_id" TEXT DEFAULT '000000' CHECK(length("tenant_id") <= 20),
			"dept_id" INTEGER NOT NULL,
			"post_code" TEXT NOT NULL CHECK(length("post_code") <= 64),
			"post_category" TEXT CHECK(length("post_category") <= 100),
			"post_name" TEXT NOT NULL CHECK(length("post_name") <= 50),
			"post_sort" INTEGER NOT NULL,
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
			g.Log().Errorf(ctx, "sMemDBSysPost.CreateTable create table err: %v", err)
			return err
		}
		g.Log().Debug(ctx, "sMemDBSysPost.CreateTable create table success", res)
		// 加载数据
		s.LoadData(ctx, 0)
	}
	return nil
}

func (s *sMemDBSysPost) LoadData(ctx context.Context, pk int64) error {
	count := 0
	if pk == 0 {
		maxPostId := int64(0)
		for {
			page := 1
			pageSize := 500
			dataList := make([]*entity.SysPost, 0)
			dao.SysPost.Ctx(ctx).Page(page, pageSize).WhereGT(dao.SysPost.Columns().PostId, maxPostId).
				OrderAsc(dao.SysPost.Columns().PostId).Scan(&dataList)
			if len(dataList) == 0 {
				break
			}
			count += len(dataList)
			_, err := db.Model(dao.SysPost.Table()).Data(dataList).OnConflict(dao.SysPost.Columns().PostId).Save()
			if err != nil {
				g.Log().Errorf(ctx, "sMemDBSysPost.Load data insert err: %v", err)
				return err
			}
			maxPostId = dataList[len(dataList)-1].PostId
			if len(dataList) < pageSize {
				break
			}
		}
	} else {
		dataList := make([]*entity.SysPost, 0)
		dao.SysPost.Ctx(ctx).Where(dao.SysPost.Columns().PostId, pk).Scan(&dataList)
		_, err := db.Model(dao.SysPost.Table()).Data(dataList).OnConflict(dao.SysPost.Columns().PostId).Save()
		if err != nil {
			g.Log().Errorf(ctx, "sMemDBSysPost.Load data insert err: %v", err)
		}
	}
	g.Log().Debugf(ctx, "sMemDBSysPost.LoadData load data success, total: %d, pk: %d", count, pk)
	return nil
}

func (s *sMemDBSysPost) DeleteData(ctx context.Context, ids []int64) error {
	_, err := db.Model(dao.SysPost.Table()).WhereIn(dao.SysPost.Columns().PostId, ids).Delete()
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysPost.DeleteData err: %v", err)
		return err
	}
	g.Log().Debugf(ctx, "sMemDBSysPost.DeleteData delete data success, total: %d, ids: %+v", len(ids), ids)
	return nil
}
