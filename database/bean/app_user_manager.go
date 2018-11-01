package bean

import (
	"oauth/database"
	"oauth/database/iredis"
	"oauth/database/schema"
)

// category: "white", "black"
func AddUserToApp(userID int64, appID int64, category string, roleID int64) error {
	engine := database.GetDriver()
	appUser := schema.AppUserList{
		UserID:   userID,
		AppID:    appID,
		Category: category,
	}
	if exist, err := engine.Get(&appUser); err != nil {
		return err
	} else if exist {
		return nil
	}
	_, err := engine.InsertOne(schema.AppUserList{
		UserID:   userID,
		AppID:    appID,
		Category: category,
		RoleID:   roleID,
	})
	return err
}

type AppUseGroup struct {
	AppUserList schema.AppUserList `xorm:"extends" json:"appUser"`
	User        schema.User        `xorm:"extends" json:"user"`
}

func GetAppUserList(appID int64, category string) (list []AppUseGroup, err error) {
	engine := database.GetDriver()

	session := engine.Table("app_user_list").
		Join("LEFT", "user", "app_user_list.user_id = user.id").
		Where("app_id = ?", appID)

	if category != "" {
		session.Where("app_user_list.category = ?", category)
	}

	err = session.Find(&list)
	return
}

func DeleteUserFromAppUserList(id int64) error {
	engine := database.GetDriver()
	_, err := engine.ID(id).Delete(schema.AppUserList{})
	return err
}

//拥有登陆权限
func HaveEnterPromise(uid int64, appID int64) (bool, error) {
	mode, err := iredis.AppCache.GetMode(appID)
	//如果获取列表失败，才用严格的白名单模式
	if err != nil {
		mode = "white"
	}
	engine := database.GetDriver()
	exist, err := engine.Table("app_user_list").Where("user_id = ? and app_id = ? and category = ?", uid, appID, mode).Exist()
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
