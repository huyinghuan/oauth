package schema

type User struct {
	ID       int64  `xorm:"id unique autoincr index pk" json:"id" formam:"-"`
	Name     string `xorm:"name" json:"name" formam:"name"`
	Password string `xorm:"password" json:"password" formam:"password"`
}

type Group struct {
	ID   int64  `xorm:"id unique autoincr index pk" json:"id" formam:"-"`
	Name string `xorm:"name" json:"name" formam:"name"`
}

type UserGroupMap struct {
	ID  int64 `xorm:"id unique autoincr index pk" json:"id" formam:"-"`
	UID int64 `xorm:"uid"`
	GID int64 `xorm:"gid"`
}

func (u *User) TableName() string          { return "user" }
func (g *Group) TableName() string         { return "group_name" }
func (ug *UserGroupMap) TableName() string { return "user_group" }
