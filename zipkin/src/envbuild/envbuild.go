package envbuild

import (
	"database/sql"
	"flag"
	"time"

	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/msbranco/goconfig"
)

var pool *redis.Pool
var db *sql.DB
var config_file_path string
var timeout int64

func init() {
	flag.StringVar(&config_file_path, "path", "../etc/config.ini", "Use -c <path>")

	if err := envBuild(); err != nil {
		panic(err)
	}

}
func Getdb() *sql.DB {
	return db
}
func GetRedisPool() *redis.Pool {
	return pool
}
func GetTimeout() int64 {
	return timeout
}
func envBuild() error {

	cf, err := goconfig.ReadConfigFile(config_file_path)

	if err != nil {
		return err
	}
	timeout, err = cf.GetInt64("LOCAL", "TIMEOUT")

	ip, _ := cf.GetString("DBCONN", "IP")
	user, _ := cf.GetString("DBCONN", "USERID")
	pwd, _ := cf.GetString("DBCONN", "USERPWD")
	name, _ := cf.GetString("DBCONN", "DBNAME")

	dbstr := user + ":" + pwd + "@tcp(" + ip + ")/" + name + "?charset=utf8"

	db, err = sql.Open("mysql", dbstr)

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.Ping()

	redisip, _ := cf.GetString("REDISCONN", "REDISADDR")
	pool = &redis.Pool{
		MaxIdle:     20,
		MaxActive:   200,
		IdleTimeout: 1 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisip)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}
