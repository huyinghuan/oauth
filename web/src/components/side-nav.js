import React from 'react';
import { Menu, Layout } from 'antd';
import { AppstoreOutlined, UnorderedListOutlined, AppstoreAddOutlined,
    TeamOutlined, UserOutlined, UserAddOutlined, HomeOutlined } from '@ant-design/icons';

import { withRouter } from 'react-router-dom';
import { get as GetData } from '../service'
const {Sider} = Layout;
const {SubMenu} = Menu;


class Component extends React.Component{
  constructor(props){
    super(props)
    this.state = {isAdmin: false}
  }
  componentDidMount(){
    GetData("/user-status").then((data)=>{
      this.setState({isAdmin: data.uid === 0})
    }).catch((e)=>{})
  }
    menuClick(item){
        let queue = [this.props.match.path]
        if(item.key !== "home"){
          queue.push(item.key)
        }
        this.props.history.push(queue.join("/"))
    }
    render(){
        let arr =  this.props.history.location.pathname.split("/")
        let key = ""
        if(arr.length === 2){
          key = arr[1]
        }else if(arr.length > 1){
          key = arr[2]
        }
        return (
            <Sider width={200} style={{ background: '#fff' }}>
                <Menu mode="inline"
                    onClick={(item, key)=>{this.menuClick(item, key)}}
                    defaultSelectedKeys={[key]}
                    defaultOpenKeys={['apps', "users"]}
                    style={{ height: '100%', borderRight: 0 }}>
                  <Menu.Item key="home"><span><span><HomeOutlined/></span>主页</span></Menu.Item>
                  <SubMenu key="apps" title={<span><span><AppstoreOutlined/></span>应用管理</span>}>
                    <Menu.Item key="app"><span><span><UnorderedListOutlined/></span>应用列表</span></Menu.Item>
                    <Menu.Item key="app-register"><span><span><AppstoreAddOutlined/></span>应用注册</span></Menu.Item>
                  </SubMenu>
                  {this.state.isAdmin ? (<SubMenu key="users" title={<span><span> <UserOutlined /></span>用户管理</span>}>
                    <Menu.Item key="user"><span><span><TeamOutlined/></span>用户列表</span></Menu.Item>
                    <Menu.Item key="user-register"><span><span><UserAddOutlined/></span>用户注册</span></Menu.Item>
                  </SubMenu>) : null }
                  
                </Menu>
              </Sider>
        )
    }
}

export default withRouter(Component)