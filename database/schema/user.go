package schema

type User struct {
	ID       int64  `xorm:"id unique autoincr index pk" json:"id"`
	Name     string `xorm:"name" json:"name"`
	Password string `xorm:"password" json:"-"`
}

func (u *User) TableName() string { return "user" }
