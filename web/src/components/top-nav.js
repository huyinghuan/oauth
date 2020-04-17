import React from 'react';
import { Menu, Layout, Typography } from 'antd';
import { UserOutlined,  LogoutOutlined, ReloadOutlined} from '@ant-design/icons';
import { get as GetData } from '../service'
import { withRouter } from 'react-router-dom';

const { Header} = Layout;
const { Title } = Typography;
const {SubMenu} = Menu;


class Component extends React.Component{
    
    constructor(props){
        super(props)
        this.state = {
            isLogin: false,
            username: "未登录"
        }
    }
    async componentDidMount(){
        try{
            let data = await GetData("/user-status")
            this.setState({
                username: data.username,
                isLogin: true
            })
        }catch(e){
            const {location} = this.props.history
            this.props.history.push(`/login?goback=${encodeURIComponent(location.pathname+location.search)}`)
        }
    }
    signOut(){
        GetData("/user-status", {method: 'Delete'})
            .then(()=>{
                this.props.history.push("/login")
            }).catch((e)=>{})
    }
    menuClick(item){
        switch(item.key){
            case "signOut":
                this.signOut()
                break
            case "resetPassword":
                this.props.history.push("/home/password-reset")
                break
            default:
        }
    }
    render(){
        return (
        <Header className="header">
            <div className="logo"><Title level={4}>MGTV OAuth</Title></div>
            <Menu
            // theme="dark"
            mode="horizontal"
            style={{ lineHeight: '64px', float: "right" }}
            onClick={(item, key)=>{this.menuClick(item, key)}}
            >
            { !this.state.isLogin ? (<Menu.Item key="3">未登录</Menu.Item>) : 
                <SubMenu key="sub1" title= {
                        <span><UserOutlined/><span>{this.state.username}</span></span>
                    }>
                    <Menu.Item key="signOut"><span><LogoutOutlined/><span>退出系统</span></span></Menu.Item>
                    <Menu.Item key="resetPassword"><span><ReloadOutlined/><span>修改密码</span></span></Menu.Item>
                </SubMenu>
            }
            </Menu>
        </Header>)
    }
}

export default withRouter(Component)