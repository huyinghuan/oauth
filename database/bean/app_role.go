package bean

import (
	"fmt"
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

func (r *role) Get(roleID int64) (role schema.AppRole, err error) {
	engine := database.GetDriver()
	_, err = engine.ID(roleID).Get(&role)
	return
}

func (r *role) GetPermission(roleID int64) (list []schema.AppRolePermission, err error) {
	engine := database.GetDriver()
	engine.Where("role_id = ?", roleID).Find(&list)
	return
}

func (r *role) Delete(id int64) error {
	engine := database.GetDriver()

	exist, err := engine.Where("role_id = ?", id).Exist(&schema.AppUserList{})
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("该角色已分配给用户，请修改用户的角色分配后再尝试删除")
	}
	sess := engine.NewSession()
	defer sess.Close()
	if err := sess.Begin(); err != nil {
		return err
	}
	if _, err := sess.ID(id).Delete(&schema.AppRole{}); err != nil {
		return err
	}
	if _, err := sess.Where("role_id = ?", id).Delete(&schema.AppRolePermission{}); err != nil {
		return err
	}
	return sess.Commit()
}

func (r *role) AppHaveRole(id int64, clientID string) (bool, error) {
	engine := database.GetDriver()
	return engine.ID(id).Where("client_id = ? ", clientID).Exist(&schema.AppRole{})
}
