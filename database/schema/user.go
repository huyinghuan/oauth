package schema

type User struct {
	ID       int64  `xorm:"id unique autoincr index pk" json:"id" formam:"-"`
	Name     string `xorm:"name" json:"name" formam:"name"`
	Password string `xorm:"password" json:"password" formam:"password"`
}

func (u *User) TableName() string { return "user" }
