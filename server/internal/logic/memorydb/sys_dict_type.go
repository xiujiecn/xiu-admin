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

type sMemDBSysDictType struct {
}

var instanceMemDBSysDictType *sMemDBSysDictType

func NewMemDBSysDictType() *sMemDBSysDictType {
	if instanceMemDBSysDictType != nil {
		return instanceMemDBSysDictType
	}
	instanceMemDBSysDictType = &sMemDBSysDictType{}
	initLoadMemoryDB := g.Cfg().MustGet(context.TODO(), "server.initLoadMemoryDB").Bool()
	if initLoadMemoryDB {
		instanceMemDBSysDictType.Init(context.TODO())
	}
	return instanceMemDBSysDictType
}

func init() {
	s := NewMemDBSysDictType()
	service.RegisterMemDBSysDictType(s)
	RegisterMemDB(s.TableName(context.TODO()), s)
}

func (s *sMemDBSysDictType) Init(ctx context.Context) error {
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

func (s *sMemDBSysDictType) EventHandler(ctx context.Context) error {
	eventList := []string{
		consts.EventKeyDBSysDictTypeCreate,
		consts.EventKeyDBSysDictTypeUpdate,
		consts.EventKeyDBSysDictTypeDelete,
	}
	event.EventsInstance().Register(eventList, func(ctx context.Context, eventKey string, args ...interface{}) {
		if len(args) < 1 {
			g.Log().Errorf(ctx, "sMemDBSysDictType.EventHandler event args error. eventKey: %s, args: %+v", eventKey, args)
			return
		}
		pkList, ok := args[0].([]int64)
		if !ok {
			pk, ok := args[0].(int64)
			if !ok {
				g.Log().Errorf(ctx, "sMemDBSysDictType.EventHandler event args error. eventKey: %s, args: %+v", eventKey, args)
				return
			}
			pkList = []int64{pk}
		}
		if len(pkList) == 0 {
			g.Log().Errorf(ctx, "sMemDBSysDictType.EventHandler event args error. eventKey: %s, args: %+v", eventKey, args)
			return
		}
		if eventKey == consts.EventKeyDBSysDictTypeDelete {
			s.DeleteData(ctx, pkList)
		} else {
			for _, pk := range pkList {
				s.LoadData(ctx, pk)
			}
		}
		dbType := "edit"
		if eventKey == consts.EventKeyDBSysDictTypeCreate {
			dbType = "add"
		} else if eventKey == consts.EventKeyDBSysDictTypeDelete {
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

func (s *sMemDBSysDictType) TableName(ctx context.Context) string {
	return dao.SysDictType.Table()
}

func (s *sMemDBSysDictType) CreateTable(ctx context.Context) error {
	_, err := db.Exec(ctx, "DROP TABLE IF EXISTS "+s.TableName(ctx))
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysDictType.CreateTable delete table err: %v", err)
		return err
	}

	total, err := db.Model("sqlite_master").Where("type", "table").Where("name", dao.SysDictType.Table()).Count()
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysDictType.CreateTable count err: %v", err)
		return err
	}
	if total == 0 {
		g.Log().Debug(ctx, "sMemDBSysDictType.List create table ", s.TableName(ctx))
		res, err := db.Exec(ctx, `
			CREATE TABLE "sys_dict_type" (
				"dict_id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				"tenant_id" TEXT DEFAULT '000000' CHECK(length("tenant_id") <= 20),
				"dict_name" TEXT DEFAULT '' CHECK(length("dict_name") <= 100),
				"dict_type" TEXT DEFAULT '' CHECK(length("dict_type") <= 100),
				"is_sys" TEXT DEFAULT '1'  ,
				"created_dept" INTEGER,
				"created_by" INTEGER,
				"created_at" TEXT,
				"updated_by" INTEGER,
				"updated_at" TEXT,
				"remark" TEXT CHECK(length("remark") <= 500)
			);
		`)
		if err != nil {
			g.Log().Errorf(ctx, "sMemDBSysDictType.CreateTable create table err: %v", err)
			return err
		}
		g.Log().Debug(ctx, "sMemDBSysDictType.CreateTable create table success", res)
		// 加载数据
		s.LoadData(ctx, 0)
	}
	return nil
}

func (s *sMemDBSysDictType) LoadData(ctx context.Context, pk int64) error {
	count := 0
	if pk == 0 {
		maxDictId := int64(0)
		for {
			page := 1
			pageSize := 500
			dataList := make([]*entity.SysDictType, 0)
			dao.SysDictType.Ctx(ctx).Page(page, pageSize).WhereGT(dao.SysDictType.Columns().DictId, maxDictId).
				OrderAsc(dao.SysDictType.Columns().DictId).Scan(&dataList)
			if len(dataList) == 0 {
				break
			}
			count += len(dataList)
			_, err := db.Model(dao.SysDictType.Table()).Data(dataList).OnConflict(dao.SysDictType.Columns().DictId).Save()
			if err != nil {
				g.Log().Errorf(ctx, "sMemDBSysDictType.Load data insert err: %v", err)
				return err
			}
			maxDictId = dataList[len(dataList)-1].DictId
			if len(dataList) < pageSize {
				break
			}
		}
	} else {
		dataList := make([]*entity.SysDictType, 0)
		dao.SysDictType.Ctx(ctx).Where(dao.SysDictType.Columns().DictId, pk).Scan(&dataList)
		_, err := db.Model(dao.SysDictType.Table()).Data(dataList).OnConflict(dao.SysDictType.Columns().DictId).Save()
		if err != nil {
			g.Log().Errorf(ctx, "sMemDBSysDictType.Load data insert err: %v", err)
		}
	}

	g.Log().Debugf(ctx, "sMemDBSysDictType.LoadData load data success, total: %d, pk: %d", count, pk)
	return nil
}

func (s *sMemDBSysDictType) DeleteData(ctx context.Context, ids []int64) error {
	_, err := db.Model(dao.SysDictType.Table()).WhereIn(dao.SysDictType.Columns().DictId, ids).Delete()
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysDictType.DeleteData err: %v", err)
		return err
	}
	g.Log().Debugf(ctx, "sMemDBSysDictType.DeleteData delete data success, total: %d, ids: %+v", len(ids), ids)
	return nil
}
