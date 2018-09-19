package schema

type Application struct {
	ID         int64  `xorm:"id unique autoincr index pk" json:"id" formam:"-"`
	UserID     int64  `xorm:"user_id" json:"user_id" formam:"-"`
	Name       string `xorm:"name" json:"name" formam:"name"`
	ClientID   string `xorm:"client_id" json:"client_id" formam:"client_id"`
	PrivateKey string `xorm:"private_key" json:"private_key" formam:"private_key"`
	Callback   string `xorm:"callback" json:"callback" formam:"callback"`
}
