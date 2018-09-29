package bean

import (
	"oauth/database"
	"oauth/database/iredis"
	"oauth/database/schema"
)

// category: "white", "black"
func AddUserToApp(userID int64, clientID string, category string) error {
	engine := database.GetDriver()
	appUser := schema.AppUserList{
		UserID:   userID,
		ClientID: clientID,
		Category: category,
	}
	if exist, err := engine.Get(&appUser); err != nil {
		return err
	} else if exist {
		return nil
	}
	_, err := engine.InsertOne(schema.AppUserList{
		UserID:   userID,
		ClientID: clientID,
		Category: category,
	})
	return err
}

type AppUseGroup struct {
	AppUserList schema.AppUserList `xorm:"extends" json:"appUserList"`
	User        schema.User        `xorm:"extends" json:"user"`
}

func GetAppUserList(clientID string, category string) (list []AppUseGroup, err error) {
	engine := database.GetDriver()
	err = engine.Table("app_user_list").Join("LEFT", "user", "app_user_list.user_id = user.id").Where("app_user_list.category = ?", category).Find(&list)
	return
}

func DeleteUserFromAppUserList(id int64) error {
	engine := database.GetDriver()
	_, err := engine.ID(id).Delete(schema.AppUserList{})
	return err
}

//拥有登陆权限
func HaveEnterPromise(uid int64, clientID string) (bool, error) {
	mode, err := iredis.AppCache.GetMode(clientID)
	//如果获取列表失败，才用严格的白名单模式
	if err != nil {
		mode = "white"
	}
	engine := database.GetDriver()
	exist, err := engine.Table("app_user_list").Where("user_id = ? and client_id = ? and category = ?", uid, clientID, mode).Exist()
	if err != nil {
		return false, err
	}
	if exist && mode == "white" {
		return true, nil
	}
	if !exist && mode == "black" {
		return true, nil
	}
	return false, nil

}
