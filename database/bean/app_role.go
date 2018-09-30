package bean

import (
	"oauth/database"
	"oauth/database/schema"
)

type role struct{}

var Role role

func (r *role) Add(name string, clientID string) error {
	engine := database.GetDriver()
	_, err := engine.InsertOne(schema.AppRole{Name: name, ClientID: clientID})
	return err
}

func (r *role) GetRoleList(clientID string) (list []schema.AppRole, err error) {
	engine := database.GetDriver()
	err = engine.Where("client_id = ?", clientID).Find(&list)
	return
}

func (r *role) GetPromise(roleID int64) (list []schema.AppRolePromise, err error) {
	engine := database.GetDriver()
	engine.Where("role_id = ?", roleID).Find(&list)
	return
}

func (r *role) Delete(id int64) error {
	engine := database.GetDriver()
	sess := engine.NewSession()
	defer sess.Close()
	if err := sess.Begin(); err != nil {
		return err
	}
	_, err := engine.ID(id).Delete(&schema.AppRole{})
	return err
}
