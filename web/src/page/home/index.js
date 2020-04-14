import React from 'react';
import { Layout, Breadcrumb } from 'antd';
import {Route, Switch, withRouter, Link} from 'react-router-dom';
import "./index.scss"

import TopNav from "../../components/top-nav"
import SideNav from "../../components/side-nav"

import App from "./apps/list"
import Role from "./role"


const { Content } = Layout;

const breadcrumbNameMap = {
  "appList": "应用列表",
  "home": "首页",
  "appRegister": "应用注册",
  "usersList":"用户列表",
  "usersRegister":"用户注册",
  "userManager": "用户管理"
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
                        { index !== 0 && ~~pathname == 0 ? (<Link to={url}>{getBreadcrumbName(pathname)}</Link>) : breadcrumbNameMap[pathname]}
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
                        <Route exact path={`${path}/appList`} component={App} />
                        <Route exact path={`${path}/role`} component={Role} />
                    </Switch>
                </Content>
              </Layout>
            </Layout>
          </Layout>)
    }
}

const p = withRouter(Page)

export default p