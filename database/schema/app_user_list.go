package schema

type AppUserList struct {
	ID       int64  `xorm:"id unique autoincr index pk" json:"id"`
	UserID   int64  `xorm:"user_id" json:"user_id" `          //用户ID
	ClientID string `xorm:"client_id index" json:"client_id"` //应用ID
	Category string `xorm:"category" json:"category"`         //黑,白名单
	RoleID   int64  `xorm:"role_id" json:"role_id"`           //角色

}

type AppRole struct {
	ID       int64  `xorm:"id unique autoincr index pk" json:"id"`
	ClientID string `xorm:"client_id index" json:"client_id"` //应用ID
	Name     string `xorm:"name" json:"name"`
}

type AppRolePermission struct {
	ID      int64  `xorm:"id unique autoincr index pk" json:"id"`
	RoleID  int64  `xorm:"role_id index"`
	Name    string `xorm:"name" json:"name"`       //权限别称
	Method  string `xorm:"method" json:"method"`   //http method
	Pattern string `xorm:"pattern" json:"pattern"` // http url reg pattern
}

func (arp *AppRolePermission) TableName() string {
	return "app_role_permission"
}

func (ar *AppRole) TableName() string {
	return "app_role"
}

func (au *AppUserList) TableName() string {
	return "app_user_list"
}
