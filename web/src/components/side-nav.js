import React from 'react';
import { Menu, Layout } from 'antd';
import { AppstoreOutlined, UnorderedListOutlined, AppstoreAddOutlined,
    TeamOutlined, UserOutlined, UserAddOutlined } from '@ant-design/icons';

import { withRouter } from 'react-router-dom';

const {Sider} = Layout;
const {SubMenu} = Menu;


class Component extends React.Component{

    menuClick(item){
        let queue = [this.props.match.path, item.key]
        this.props.history.push(queue.join("/"))
    }
    render(){
        return (
            <Sider width={200} style={{ background: '#fff' }}>
                <Menu mode="inline"
                    onClick={(item, key)=>{this.menuClick(item, key)}}
                    defaultSelectedKeys={['appList']}
                    defaultOpenKeys={['apps']}
                    style={{ height: '100%', borderRight: 0 }}>
                  <SubMenu key="apps" title={<span><span><AppstoreOutlined/></span>应用管理</span>}>
                    <Menu.Item key="appList"><span><span><UnorderedListOutlined/></span>列表</span></Menu.Item>
                    <Menu.Item key="appRegister"><span><span><AppstoreAddOutlined/></span>注册</span></Menu.Item>
                  </SubMenu>
                  <SubMenu key="users" title={<span><span> <UserOutlined /></span>用户管理</span>}>
                    <Menu.Item key="usersList"><span><span><TeamOutlined/></span>列表</span></Menu.Item>
                    <Menu.Item key="usersRegister"><span><span><UserAddOutlined/></span>注册</span></Menu.Item>
                  </SubMenu>
                </Menu>
              </Sider>
        )
    }
}

export default withRouter(Component)