package gfredis

import (
	"context"
	"crypto/tls"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/redis/go-redis/v9"
)

const (
	DefaultGroupName = "default" // Default configuration group name.
)

var (
	lock      sync.RWMutex
	instances map[string]*Redis
)

func init() {
	instances = make(map[string]*Redis)
}

type Redis struct {
	config *Config
	rdb    *redis.Client
}

// Config is redis configuration.
type Config struct {
	// Address It supports single and cluster redis server. Multiple addresses joined with char ','. Eg: 192.168.1.1:6379, 192.168.1.2:6379.
	Address         string        `json:"address"`
	Db              int           `json:"db"`              // Redis db.
	User            string        `json:"user"`            // Username for AUTH.
	Pass            string        `json:"pass"`            // Password for AUTH.
	SentinelUser    string        `json:"sentinel_user"`   // Username for sentinel AUTH.
	SentinelPass    string        `json:"sentinel_pass"`   // Password for sentinel AUTH.
	MinIdle         int           `json:"minIdle"`         // Minimum number of connections allowed to be idle (default is 0)
	MaxIdle         int           `json:"maxIdle"`         // Maximum number of connections allowed to be idle (default is 10)
	MaxActive       int           `json:"maxActive"`       // Maximum number of connections limit (default is 0 means no limit).
	MaxConnLifetime time.Duration `json:"maxConnLifetime"` // Maximum lifetime of the connection (default is 30 seconds, not allowed to be set to 0)
	IdleTimeout     time.Duration `json:"idleTimeout"`     // Maximum idle time for connection (default is 10 seconds, not allowed to be set to 0)
	WaitTimeout     time.Duration `json:"waitTimeout"`     // Timed out duration waiting to get a connection from the connection pool.
	DialTimeout     time.Duration `json:"dialTimeout"`     // Dial connection timeout for TCP.
	ReadTimeout     time.Duration `json:"readTimeout"`     // Read timeout for TCP. DO NOT set it if not necessary.
	WriteTimeout    time.Duration `json:"writeTimeout"`    // Write timeout for TCP.
	MasterName      string        `json:"masterName"`      // Used in Redis Sentinel mode.
	TLS             bool          `json:"tls"`             // Specifies whether TLS should be used when connecting to the server.
	TLSSkipVerify   bool          `json:"tlsSkipVerify"`   // Disables server name verification when connecting over TLS.
	TLSConfig       *tls.Config   `json:"-"`               // TLS Config to use. When set TLS will be negotiated.
	SlaveOnly       bool          `json:"slaveOnly"`       // Route all commands to slave read-only nodes.
	Cluster         bool          `json:"cluster"`         // Specifies whether cluster mode be used.
	Protocol        int           `json:"protocol"`        // Specifies the RESP version (Protocol 2 or 3.)
}

func New(groupName string, config *Config) *Redis {
	instance := &Redis{
		config: config,
	}
	if config == nil {
		instance.rdb = redis.NewClient(&redis.Options{})
		return instance
	}
	poolSize := config.MaxActive
	if poolSize == 0 {
		poolSize = 10
		config.MaxActive = 10
	}

	if config.IdleTimeout < 1*time.Minute {
		config.IdleTimeout = 30 * time.Minute
	}
	if config.MaxConnLifetime <= 1 {
		config.MaxConnLifetime = 24 * 365 * 10 * time.Hour
	}
	if config.WaitTimeout < 1*time.Second {
		config.WaitTimeout = 6 * time.Second
	}
	if config.DialTimeout < 1*time.Second {
		config.DialTimeout = 5 * time.Second
	}
	if config.ReadTimeout < 1*time.Second {
		config.ReadTimeout = 5 * time.Second
	}
	if config.WriteTimeout < 1*time.Second {
		config.WriteTimeout = 5 * time.Second
	}
	opt := &redis.Options{
		Addr:            config.Address,
		Password:        config.Pass,
		DB:              config.Db,
		MinIdleConns:    config.MinIdle,
		MaxIdleConns:    config.MaxIdle,
		MaxActiveConns:  config.MaxActive,
		ConnMaxIdleTime: config.IdleTimeout,
		PoolSize:        poolSize,
		PoolTimeout:     config.WaitTimeout,
		DialTimeout:     config.DialTimeout,
		ReadTimeout:     config.ReadTimeout,
		WriteTimeout:    config.WriteTimeout,
		TLSConfig:       config.TLSConfig,
		// IdentitySuffix:  config.IdentitySuffix,
		// ConnMaxLifetime: config.MaxConnLifetime,
	}
	instance.rdb = redis.NewClient(opt)
	g.Log().Noticef(context.Background(), "groupName: %s, opt: %+v", groupName, opt)
	return instance
}

func GetConfig(name string) *Config {
	cfg, err := g.Cfg().Get(context.Background(), "redis."+name)
	if err != nil {
		panic(err)
	}
	if cfg == nil || cfg.IsNil() {
		panic("redis config not found: " + name)
	}
	config := &Config{}
	if err = gconv.Scan(cfg.Map(), config); err != nil {
		panic(`invalid redis configuration: ` + cfg.String())
	}
	return config
}

func Instance(name ...string) *Redis {
	groupName := DefaultGroupName
	if len(name) > 0 {
		groupName = name[0]
		if groupName == "" {
			groupName = DefaultGroupName
		}
	}
	lock.RLock()
	instance, ok := instances[groupName]
	lock.RUnlock()
	if !ok {
		lock.Lock()
		instance, ok = instances[groupName]
		if !ok {
			instance = New(groupName, GetConfig(groupName))
			instances[groupName] = instance
		}
		lock.Unlock()
		return instance
	}
	return instance
}
