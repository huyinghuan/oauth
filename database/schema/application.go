package schema

type Application struct {
	ID         int64  `xorm:"id unique autoincr index pk" json:"id" `
	UserID     int64  `xorm:"user_id" json:"user_id" `
	Name       string `xorm:"name" json:"name"`
	ClientID   string `xorm:"client_id" json:"client_id"`
	PrivateKey string `xorm:"private_key" json:"private_key" `
	Callback   string `xorm:"callback" json:"callback"`
	Mode       string `xoram:"mode" json:"mode"` //运行模式
}
