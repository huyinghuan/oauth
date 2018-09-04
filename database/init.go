package database

import (
	"log"
	"oauth/config"
	"oauth/database/schema"
	"oauth/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func initAdmin(user string, password string) error {
	exist, err := engine.Get(&schema.User{
		Name: user,
	})
	if err != nil {
		return err
	}
	if !exist {
		user := schema.User{
			Name:     user,
			Password: utils.Encrypt(password),
		}
		if _, err = engine.InsertOne(&user); err != nil {
			return err
		} else {
			log.Println("admin账户初始化完成")
		}
		return nil
	} else {
		user := schema.User{
			Name:     user,
			Password: utils.Encrypt(password),
		}
		if _, err := engine.Update(user); err != nil {
			return err
		}
		log.Println("更新Amin密码完成")
		return nil
	}
}

// InitDriver 初始化数据库链接
func init() {
	conf := config.Get()
	var connectErr error
	engine, connectErr = xorm.NewEngine(conf.Db.Driver, conf.Db.Connect)
	if conf.Dev {
		engine.ShowSQL(true)
	}
	if connectErr != nil {
		log.Fatal(connectErr)
	}
	if err := engine.Ping(); err != nil {
		log.Fatal("数据库链接失败", err)
	}
	engine.Sync(new(schema.User), new(schema.Group), new(schema.UserGroupMap))

	log.Printf("数据库连接成功")

	account := conf.Account

	if err := initAdmin(account.User, account.Pass); err != nil {
		log.Fatal(err)
	}
}

// GetDriver 获取数据库链接
func GetDriver() *xorm.Engine {
	return engine
}
