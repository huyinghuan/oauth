package bean

import (
	"fmt"
	"oauth/database"
	"oauth/database/schema"
)

type role struct{}

var Role role

func (r *role) Add(name string, appID int64) error {
	engine := database.GetDriver()
	_, err := engine.InsertOne(schema.AppRole{Name: name, AppID: appID})
	return err
}

func (r *role) GetRoleList(appID int64) (list []schema.AppRole, err error) {
	engine := database.GetDriver()
	err = engine.Where("app_id = ?", appID).Find(&list)
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
		sess.Rollback()
		return err
	}
	if _, err := sess.Where("role_id = ?", id).Delete(&schema.AppRolePermission{}); err != nil {
		sess.Rollback()
		return err
	}
	return sess.Commit()
}

func (r *role) AppHaveRole(id int64, appID int64) (bool, error) {
	engine := database.GetDriver()
	return engine.ID(id).Where("app_id = ? ", appID).Exist(&schema.AppRole{})
}

func (r *role) GetRoleIDByUserIDInApp(appID int64, userID int64) (roleID int64, err error) {
	engine := database.GetDriver()
	appUser := schema.AppUserList{}
	if _, err = engine.Where("user_id = ? and app_id = ?", userID, appID).Get(&appUser); err != nil {
		return
	}
	roleID = appUser.RoleID
	return
}

func (r *role) Update(id int64, name string) error{
	engine := database.GetDriver()
	_, err:= engine.Where("id = ?", id).Cols("name").Update(schema.AppRole{Name: name})
	return err
}