package bean

import (
	"fmt"
	"oauth/database"
	"oauth/database/schema"
	"oauth/utils"
)

func FindUser(name string, password string) (schema.User, bool, error) {
	engine := database.GetDriver()
	user := schema.User{
		Name:     name,
		Password: utils.Encrypt(password),
	}
	exist, err := engine.Get(&user)
	return user, exist, err
}

func RegisterUser(name string, password string) error {
	engine := database.GetDriver()
	exist, err := engine.Get(&schema.User{
		Name: name,
	})
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("已存在用户名")
	}

	user := schema.User{
		Name:     name,
		Password: utils.Encrypt(password),
	}
	_, err = engine.InsertOne(&user)
	if err != nil {
		return err
	}
	return nil
}
