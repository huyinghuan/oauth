package schema

type Application struct {
	ID         int64  `xorm:"id unique autoincr index pk" json:"id" formam:"-"`
	Name       string `xorm:"name" json:"name" formam:"name"`
	ClientID   string `xorm:"client_id" json:"name" formam:"client_id"`
	PrivateKey string `xorm:"private_key" json:"private_key" formam:"private_key"`
	Password   string `xorm:"password" json:"password" formam:"password"`
}
