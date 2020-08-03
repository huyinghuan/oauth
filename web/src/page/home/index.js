import React from 'react';
import { Layout, Breadcrumb } from 'antd';
import {Route, Switch, withRouter, Link} from 'react-router-dom';
import "./index.scss"

import TopNav from "../../components/top-nav"
import SideNav from "../../components/side-nav"

import App from "./apps/list"
import AppEdit from "./apps/edit"
import AppSetting from './app-settings/index'
import User from "./user/list"
import PasswordReset from "./user/password-reset"
import Register from './user/register'

const { Content } = Layout;

const breadcrumbNameMap = {
  "app": "应用列表",
  "home": "首页",
  "app-register": "应用注册",
  "user":"用户列表",
  "user-register":"用户注册",
  "app-settings": "权限配置",
  "password-reset":"密码重置",
  "app-role":"角色",
  "app-black":"黑名单",
  "app-white":"白名单",
  "role-settings":"权限分配"
}

function getBreadcrumbName(name){
  return breadcrumbNameMap[name] || name
}

class Page extends React.Component {
    render() {
        let path = this.props.match.path
        let pathList = this.props.history.location.pathname.split("/").filter(i => i)
        return (
        <Layout className="page-home">
            <TopNav />
            <Layout>
              <SideNav />
              <Layout style={{ padding: '0 24px 24px' }}>
                <Breadcrumb style={{ margin: '16px 0' }}>
                  {pathList.map((pathname, index)=>{
                    const url = `/${pathList.slice(0, index + 1).join('/')}`;
                    return (
                      <Breadcrumb.Item key={url}>
                        { index !== 0 && ~~pathname === 0 ? (<Link to={url}>{getBreadcrumbName(pathname)}</Link>) : getBreadcrumbName(pathname)}
                      </Breadcrumb.Item>
                    )
                  })}
                </Breadcrumb>
                <Content
                  style={{
                    background: '#fff',
                    padding: 24,
                    margin: 0,
                    height: "100%",
                    overflow: "hidden"
                  }}
                >
                    <Switch>
                        <Route exact path={`${path}/app`} component={App} />
                        <Route exact path={`${path}/app/:appId`} component={AppEdit}/>
                        <Route path={`${path}/app/:appId/app-settings`} component={AppSetting}/>
                        <Route exact path={`${path}/app/:appId/role`} component={AppEdit}/>
                        <Route exact path={`${path}/app-register`}  component={AppEdit} />
                        <Route exact path={`${path}/user`} component={User} />
                        <Route exact path={`${path}/user-register`} component={Register} />
                        <Route exact path={`${path}/user/:id/password-reset`} component={PasswordReset}/>
                        <Route exact path={`${path}/password-reset`} component={PasswordReset}/>
                        {/* <Route exact path={`${path}/role`} component={Role} /> */}
                    </Switch>
                </Content>
              </Layout>
            </Layout>
          </Layout>)
    }
}

const p = withRouter(Page)

export default p