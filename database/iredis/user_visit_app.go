package iredis

type userVisitApp struct {
	Format    string
	MapFormat string
}

func (cache *userVisitApp) Set() {

}

// UserVisitApp 缓存用户使用过的app
var UserVisitApp = userVisitApp{Format: "user:%d:apps"}
