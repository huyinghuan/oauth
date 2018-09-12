package bean

import (
	"fmt"
	"oauth/database"
	"oauth/database/schema"
	"oauth/utils"
)

func FindApplicationByClientID(clientID string) (schema.Application, error) {
	app := schema.Application{
		ClientID: clientID,
	}
	engine := database.GetDriver()
	_, err := engine.Get(&app)
	return app, err
}

func GetApplictionList() ([]schema.Application, error) {
	engine := database.GetDriver()
	list := make([]schema.Application, 0)
	err := engine.Find(&list)
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
	app.Password = utils.Encrypt(app.Password)
	_, err = engine.InsertOne(app)
	if err != nil {
		return err
	}
	return nil
}
