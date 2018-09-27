package bean

import (
	"fmt"
	"oauth/database"
	"oauth/database/schema"
	"oauth/utils"
)

func FindApplicationByID(id int64) (schema.Application, error) {
	app := schema.Application{}
	engine := database.GetDriver()
	_, err := engine.Id(id).Get(&app)
	return app, err
}

func FindApplicationByClientID(clientID string) (schema.Application, error) {
	app := schema.Application{
		ClientID: clientID,
	}
	engine := database.GetDriver()
	_, err := engine.Get(&app)
	return app, err
}

// func GetApplictionList() ([]schema.Application, error) {
// 	engine := database.GetDriver()
// 	list := make([]schema.Application, 0)
// 	err := engine.Find(&list)
// 	return list, err
// }

type ApplicationUserGroup struct {
	Appliction schema.Application `xorm:"extends" json:"application"`
	User       schema.User        `xorm:"extends" json:"user"`
}

func GetApplictionList(userID int64) ([]ApplicationUserGroup, error) {

	engine := database.GetDriver()

	session := engine.Table("application").
		Join("LEFT", "user", "application.user_id = user.id")

	if userID != -1 {
		session.Where("user.id = ?", userID)
	}
	list := make([]ApplicationUserGroup, 0)
	err := session.Find(&list)
	return list, err
}

func RegisterAppliction(app *schema.Application) error {
	engine := database.GetDriver()
	exist, err := engine.Get(&schema.Application{
		Name: app.Name,
	})
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("已存在应用名")
	}
	app.ClientID = utils.RandomString(24)
	app.PrivateKey = utils.RandomString(24)
	_, err = engine.InsertOne(app)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAppliction(id int64) error {
	engine := database.GetDriver()
	app := new(schema.Application)
	_, err := engine.Id(id).Delete(app)
	return err
}

func GetAppliction(id int64, uid int64) (schema.Application, error) {
	engine := database.GetDriver()
	app := schema.Application{}
	session := engine.ID(id)
	if uid != -1 {
		session.Where("user_id = ?", uid)
	}
	_, err := session.Get(&app)
	return app, err
}

func UpdateApplication(id int64, uid int64, app *schema.Application) (*schema.Application, error) {
	engine := database.GetDriver()
	findApp := schema.Application{
		Name: app.Name,
	}
	exist, err := engine.Get(&findApp)
	if err != nil {
		return nil, err
	}

	if exist && findApp.ID != id {
		return nil, fmt.Errorf("已存在应用名")
	}
	findApp.Callback = app.Callback
	findApp.Name = app.Name
	_, err = engine.ID(id).Update(&findApp)
	return &findApp, err
}
