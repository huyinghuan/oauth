package bean

import (
	"fmt"
	"oauth/database"
	"oauth/database/schema"
	"oauth/utils"
)

func GetUserByID(id int64) (schema.User, error) {
	engine := database.GetDriver()
	user := schema.User{}
	_, err := engine.ID(id).Get(&user)
	return user, err
}

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

func UpdateUserPassword(uid int64, oldPass string, newPass string) error {
	encrypOldPass := utils.Encrypt(oldPass)
	encrypNewPass := utils.Encrypt(newPass)
	engine := database.GetDriver()
	user := schema.User{}
	exists, err := engine.ID(uid).Where("password = ?", encrypOldPass).Get(&user)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("密码错误")
	}
	_, err = engine.ID(uid).Cols("password").Update(&schema.User{
		Password: encrypNewPass,
	})
	return err
}

func UpdateUserPasswordNoOld(uid int64, password string) error {
	password = utils.Encrypt(password)
	engine := database.GetDriver()
	_, err := engine.ID(uid).Cols("password").Update(&schema.User{
		Password: password,
	})
	return err
}

func DeleteUser(uid int64) error {
	engine := database.GetDriver()
	_, e := engine.ID(uid).Delete(&schema.User{})
	return e
}
