package redisson

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

// Redisson 自测试reids命令
type Redisson struct {
	pool *redis.Pool
	defaultTimeout time.Duration
}

// NewRedisson 创建连接对象 这里需要一个配置对象
func NewRedisson(host string, port int,db int) *Redisson{
	pool := &redis.Pool{
		MaxActive: 64,
		MaxIdle: 32,
		IdleTimeout: time.Duration(100) * time.Second,
		Wait: true,
		Dial: func()(redis.Conn,error) {
			conn,err := redis.Dial("tcp",fmt.Sprintf("%s:%d",host,port),redis.DialPassword("123456"), redis.DialDatabase(db))
			
			if err != nil {
				return nil, err
			}
			return conn,nil
		},
	}

	return &Redisson{pool: pool,defaultTimeout: time.Duration(1) * time.Second}
}

// Do proxy of method
func (r *Redisson) Do(commandName string, args ...interface{}) (reply interface{}, err error){
	conn := r.pool.Get()
	defer conn.Close()
	// 参数后需添加... 代表这是个数组，而不是一个interface
	return  conn.Do(commandName,args...)
}


// func (r *Redisson) do2(a func()(reply interface{}, err error)) (reply interface{}, err error){
// 	return a();
// }