package database

import (
	"log"
	"oauth/config"
	"oauth/database/iredis"
	"oauth/database/schema"
	"oauth/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

//初始化appid-clientID映射
func initAppIDMapToClientID() error {
	list := []schema.Application{}
	if err := engine.Find(&list); err != nil {
		return err
	}
	for _, app := range list {
		iredis.AppCache.SetMap(app.ID, app.ClientID)
		iredis.AppCache.SetAll(app.ID, app.PrivateKey, app.Callback, app.Mode)
	}
	return nil
}

//初始化管理员账户
func initAdmin(username string, password string, resetOnRestart bool) error {
	user := schema.User{
		Name: username,
	}
	exist, err := engine.Get(&user)
	if err != nil {
		return err
	}
	if !exist {
		newUser := schema.User{
			Name:     username,
			Password: utils.Encrypt(password),
		}
		if _, err = engine.InsertOne(&newUser); err != nil {
			return err
		} else {
			log.Println("admin账户初始化完成")
		}
		return nil
	}
	if !resetOnRestart {
		return nil
	}
	user.Password = utils.Encrypt(password)
	if _, err := engine.Id(user.ID).Update(user); err != nil {
		return err
	}
	log.Println("更新Amin密码完成")
	return nil
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
	engine.Sync(new(schema.User), new(schema.Application), new(schema.AppUserList), new(schema.AppRole), new(schema.AppRolePermission))

	log.Printf("数据库连接成功")

	account := conf.Account
	if err := initAdmin(account.User, account.Pass, account.ResetOnRestart); err != nil {
		log.Fatal(err)
	}
	if conf.RedisCacheFromDB {
		initAppIDMapToClientID()
	}
}

// GetDriver 获取数据库链接
func GetDriver() *xorm.Engine {
	return engine
}
