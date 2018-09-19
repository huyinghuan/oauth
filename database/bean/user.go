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

func GetAllUser() ([]schema.User, error) {
	engine := database.GetDriver()
	list := make([]schema.User, 0)
	err := engine.Find(&list)
	return list, err
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

func FindUserByUsername(name string) (*schema.User, error) {
	engine := database.GetDriver()
	user := schema.User{
		Name: name,
	}
	exist, err := engine.Get(&user)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("不存在用户")
	}
	return &user, nil
}
