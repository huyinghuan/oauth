package bean

import (
	"oauth/database"
	"oauth/database/schema"
)

// category: "white", "black"
func AddUserToApp(userID int64, clientID string, category string) error {
	engine := database.GetDriver()
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

func GetAppUserList(clientID string, category string) {
	engine := database.GetDriver()
	list := []schema.AppUserList{}

	engine.Table("app_user_list").Join("LEFT", "user", "app_user_list.user_id = user.id").Find(&list)
}
