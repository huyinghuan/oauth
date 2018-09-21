package schema

type AppUserList struct {
	ID       int64  `xorm:"id unique autoincr index pk" json:"id"`
	UserID   int64  `xorm:"user_id" json:"user_id" `
	ClientID string `xorm:"client_id" json:"client_id"`
	Category string `xorm:"category" json:"category"`
}

func (au *AppUserList) TableName() string {
	return "app_user_list"
}
