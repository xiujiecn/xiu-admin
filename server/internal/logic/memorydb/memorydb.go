package memorydb

import (
	"context"
	"errors"
	"time"
	"xiuadmin/internal/cmd/inithttp"
	"xiuadmin/internal/model"
	"xiuadmin/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

var db gdb.DB

type memDB interface {
	Init(ctx context.Context) error
	EventHandler(ctx context.Context) error
	TableName(ctx context.Context) string
	CreateTable(ctx context.Context) error
	LoadData(ctx context.Context, pk int64) error
	DeleteData(ctx context.Context, ids []int64) error
}

var memDBMap = make(map[string]memDB)

func RegisterMemDB(name string, memDB memDB) {
	memDBMap[name] = memDB
}

type sMemoryDB struct {
	memDBMap map[string]interface{}
}

func NewMemoryDB() *sMemoryDB {
	return &sMemoryDB{}
}

func init() {
	service.RegisterMemoryDB(NewMemoryDB())
	inithttp.RegisterHttpInitFunc("memorydb", func(ctx context.Context) error {
		service.MemoryDB().Init(ctx)
		return nil
	})
}

func (s *sMemoryDB) Init(ctx context.Context) error {
	link := g.Cfg().MustGet(ctx, "database.memorydb.link", "").String()
	debug := g.Cfg().MustGet(ctx, "database.memorydb.debug", "false").Bool()
	if link == "" {
		g.Log().Warning(ctx, "sMemoryDB.Init link is empty")
		return nil
	}
	g.Log().Debug(ctx, "sMemoryDB.Init link", link)
	var err error
	db, err = gdb.New(gdb.ConfigNode{
		Link:             link,
		MaxIdleConnCount: 10,
		MaxOpenConnCount: 1,
		MaxConnLifeTime:  24 * 3650 * time.Hour,
		Debug:            debug,
	})
	if err != nil {
		g.Log().Errorf(ctx, "sMemoryDB.Init get db err: %v", err)
		return err
	}
	if db == nil {
		g.Log().Errorf(ctx, "sMemoryDB.Init not found memorydb")
		return errors.New("sMemoryDB.Init get db err")
	}
	for _, memDB := range memDBMap {
		err = memDB.Init(ctx)
		if err != nil {
			return err
		}
	}
	g.Log().Debug(ctx, "sMemoryDB.Init end.", db.GetConfig())
	return nil
}

func (s *sMemoryDB) EventHandler(ctx context.Context, chg *model.BroadcastDbChgCache) error {
	memDB, ok := memDBMap[chg.TableName]
	if !ok {
		return nil
	}

	if chg.Type == "add" || chg.Type == "edit" {
		for _, pk := range chg.PK {
			memDB.LoadData(ctx, pk)
		}
	} else if chg.Type == "del" {
		memDB.DeleteData(ctx, chg.PK)
	}
	return nil
}

func (s *sMemoryDB) DB(ctx context.Context) gdb.DB {
	link := g.Cfg().MustGet(ctx, "database.memorydb.link", "").String()
	if link == "" {
		return g.DB()
	}
	return db
}

func (s *sMemoryDB) Ctx(ctx context.Context) gdb.DB {
	return s.DB(ctx).Ctx(ctx)
}
