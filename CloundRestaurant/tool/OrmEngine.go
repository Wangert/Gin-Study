package tool

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"CloundRestaurant/model"
)

var DBEngine *Orm

type Orm struct {
	*xorm.Engine
}

//初始化xorm引擎
func OrmEngine(cfg *Config) (*Orm, error){

	dbConfig := cfg.Database
	//设置数据库连接参数
	conn := dbConfig.User + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.DBName

	engine, err := xorm.NewEngine(dbConfig.Driver, conn)
	if err != nil {
		return nil, err
	}
	//设置是否输出sql语言
	engine.ShowSQL(dbConfig.ShowSql)
	//同步表到数据库
	err = engine.Sync2(new(model.Member))
	if err != nil {
		return nil, err
	}

	orm := new(Orm)
	orm.Engine = engine
	DBEngine = orm

	return orm, nil
}

