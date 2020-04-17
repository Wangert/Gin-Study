package tool

import (
	"github.com/gin-contrib/sessions/redis"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//初始化session操作
func InitSession(engine *gin.Engine)  {

	config := GetConfig().Redis
	store, err := redis.NewStore(10, "tcp", config.Addr + ":" + config.Port, "", []byte("secert"))
	if err != nil {
		fmt.Println(err.Error())
	}

	//使用中间件
	engine.Use(sessions.Sessions("mysession", store))
}

//set方法
func SessionSet(context *gin.Context, key, value interface{}) error {
	session := sessions.Default(context)
	if session == nil {
		return nil
	}

	session.Set(key,value)
	return session.Save()
}

//get方法
func SessionGet(context *gin.Context, key interface{}) interface{} {
	session := sessions.Default(context)
	return session.Get(key)
}
