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

type sMemDBSysDictData struct {
}

var instanceMemDBSysDictData *sMemDBSysDictData

func NewMemDBSysDictData() *sMemDBSysDictData {
	if instanceMemDBSysDictData != nil {
		return instanceMemDBSysDictData
	}
	instanceMemDBSysDictData = &sMemDBSysDictData{}
	initLoadMemoryDB := g.Cfg().MustGet(context.TODO(), "server.initLoadMemoryDB").Bool()
	if initLoadMemoryDB {
		instanceMemDBSysDictData.Init(context.TODO())
	}
	return instanceMemDBSysDictData
}

func init() {
	s := NewMemDBSysDictData()
	service.RegisterMemDBSysDictData(s)
	RegisterMemDB(s.TableName(context.TODO()), s)
}

func (s *sMemDBSysDictData) Init(ctx context.Context) error {
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

func (s *sMemDBSysDictData) EventHandler(ctx context.Context) error {
	eventList := []string{
		consts.EventKeyDBSysDictDataCreate,
		consts.EventKeyDBSysDictDataUpdate,
		consts.EventKeyDBSysDictDataDelete,
	}
	event.EventsInstance().Register(eventList, func(ctx context.Context, eventKey string, args ...interface{}) {
		if len(args) < 1 {
			g.Log().Errorf(ctx, "sMemDBSysDictData.EventHandler event args error. eventKey: %s, args: %+v", eventKey, args)
			return
		}
		pkList, ok := args[0].([]int64)
		if !ok {
			pk, ok := args[0].(int64)
			if !ok {
				g.Log().Errorf(ctx, "sMemDBSysDictData.EventHandler event args error. eventKey: %s, args: %+v", eventKey, args)
				return
			}
			pkList = []int64{pk}
		}
		if len(pkList) == 0 {
			g.Log().Errorf(ctx, "sMemDBSysDictData.EventHandler event args error. eventKey: %s, args: %+v", eventKey, args)
			return
		}
		if eventKey == consts.EventKeyDBSysDictDataDelete {
			s.DeleteData(ctx, pkList)
		} else {
			for _, pk := range pkList {
				s.LoadData(ctx, pk)
			}
		}
		dbType := "edit"
		if eventKey == consts.EventKeyDBSysDictDataCreate {
			dbType = "add"
		} else if eventKey == consts.EventKeyDBSysDictDataDelete {
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

func (s *sMemDBSysDictData) TableName(ctx context.Context) string {
	return dao.SysDictData.Table()
}

func (s *sMemDBSysDictData) CreateTable(ctx context.Context) error {
	_, err := db.Exec(ctx, "DROP TABLE IF EXISTS "+s.TableName(ctx))
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysDictData.CreateTable delete table err: %v", err)
		return err
	}

	total, err := db.Model("sqlite_master").Where("type", "table").Where("name", dao.SysDictData.Table()).Count()
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysDictData.CreateTable count err: %v", err)
		return err
	}
	if total == 0 {
		g.Log().Debug(ctx, "sMemDBSysDictData.List create table ", s.TableName(ctx))
		res, err := db.Exec(ctx, `
			CREATE TABLE "sys_dict_data" (
				"dict_code" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				"tenant_id" TEXT DEFAULT '000000' CHECK(length("tenant_id") <= 20),
				"dict_sort" INTEGER DEFAULT 0,
				"dict_label" TEXT DEFAULT '' CHECK(length("dict_label") <= 100),
				"dict_value" TEXT DEFAULT '' CHECK(length("dict_value") <= 100),
				"dict_type" TEXT DEFAULT '' CHECK(length("dict_type") <= 100),
				"css_class" TEXT CHECK(length("css_class") <= 100),
				"list_class" TEXT CHECK(length("list_class") <= 100),
				"is_default" TEXT DEFAULT 'N' ,
				"created_dept" INTEGER,
				"created_by" INTEGER,
				"created_at" TEXT,
				"updated_by" INTEGER,
				"updated_at" TEXT,
				"remark" TEXT CHECK(length("remark") <= 500)
				);
		`)
		if err != nil {
			g.Log().Errorf(ctx, "sMemDBSysDictData.CreateTable create table err: %v", err)
			return err
		}
		g.Log().Debug(ctx, "sMemDBSysDictData.CreateTable create table success", res)
		// 加载数据
		s.LoadData(ctx, 0)
	}
	return nil
}

func (s *sMemDBSysDictData) LoadData(ctx context.Context, pk int64) error {
	count := 0
	if pk == 0 {
		maxDictCode := int64(0)
		for {
			page := 1
			pageSize := 500
			dataList := make([]*entity.SysDictData, 0)
			dao.SysDictData.Ctx(ctx).Page(page, pageSize).WhereGT(dao.SysDictData.Columns().DictCode, maxDictCode).
				OrderAsc(dao.SysDictData.Columns().DictCode).Scan(&dataList)
			if len(dataList) == 0 {
				break
			}
			count += len(dataList)
			_, err := db.Model(dao.SysDictData.Table()).Data(dataList).OnConflict(dao.SysDictData.Columns().DictCode).Save()
			if err != nil {
				g.Log().Errorf(ctx, "sMemDBSysDictData.Load data insert err: %v", err)
				return err
			}
			maxDictCode = dataList[len(dataList)-1].DictCode
			if len(dataList) < pageSize {
				break
			}
		}
	} else {
		dataList := make([]*entity.SysDictData, 0)
		dao.SysDictData.Ctx(ctx).Where(dao.SysDictData.Columns().DictCode, pk).Scan(&dataList)
		_, err := db.Model(dao.SysDictData.Table()).Data(dataList).OnConflict(dao.SysDictData.Columns().DictCode).Save()
		if err != nil {
			g.Log().Errorf(ctx, "sMemDBSysDictData.Load data insert err: %v", err)
		}
	}

	g.Log().Debugf(ctx, "sMemDBSysDictData.LoadData load data success, total: %d, pk: %d", count, pk)
	return nil
}

func (s *sMemDBSysDictData) DeleteData(ctx context.Context, ids []int64) error {
	_, err := db.Model(dao.SysDictData.Table()).WhereIn(dao.SysDictData.Columns().DictCode, ids).Delete()
	if err != nil {
		g.Log().Errorf(ctx, "sMemDBSysDictData.DeleteData err: %v", err)
		return err
	}
	g.Log().Debugf(ctx, "sMemDBSysDictData.DeleteData delete data success, total: %d, ids: %+v", len(ids), ids)
	return nil
}
