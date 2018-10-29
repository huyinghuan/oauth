package bean

import (
	"oauth/database"
	"oauth/database/schema"
)

type perssion struct{}

var Perssion perssion

func (p *perssion) Add(data *schema.AppRolePermission) error {
	engine := database.GetDriver()
	_, err := engine.InsertOne(data)
	return err
}

func (p *perssion) Remove(id int64) error {
	engine := database.GetDriver()
	_, err := engine.ID(id).Delete(&schema.AppRolePermission{})
	return err
}
