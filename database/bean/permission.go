package bean

import (
	"oauth/database"
	"oauth/database/schema"
)

type permission struct{}

var Perssion permission

func (p *permission) Add(data *schema.AppRolePermission) error {
	engine := database.GetDriver()
	_, err := engine.InsertOne(data)
	return err
}

func (p *permission) Remove(id int64) error {
	engine := database.GetDriver()
	_, err := engine.ID(id).Delete(&schema.AppRolePermission{})
	return err
}

func (p *permission) GetPermision(roleID int64) (list []schema.AppRolePermission, err error) {
	engine := database.GetDriver()
	err = engine.Where("role_id = ?", roleID).Find(&list)
	return
}
