import React from 'react';
import { get as GetData } from '../../../service'
import { withRouter, Route,  Switch as RouteSwitch} from 'react-router-dom';
import { Typography, Modal,  Tabs, Divider, Tag} from 'antd';

import { UserOutlined, ControlOutlined, TeamOutlined, UsergroupDeleteOutlined} from '@ant-design/icons';
import RunModel from "./components/run-mode"
import UserList from "./components/user-list"
import RoleList from "./components/role-list"

import RoleEdit from "./components/role-edit"

const { TabPane } = Tabs;
const { Title } = Typography;

class Page extends React.Component{
    constructor(props){
        super(props)
        this.state = {app: {}, roleList:[]}
        this.modeNameMap ={
            "white": "白名单",
            "black": "黑名单"
        }
    }

    getModeNameMap(key){
        return this.modeNameMap[key]
    }

    componentDidMount(){
        this.loadApp()
        this.loadRoles()
    }
    componentWillUnmount(){
        Modal.destroyAll()
    }

    loadApp(){
        GetData(`/app/${this.props.match.params.appId}`).then((data)=>{
            this.setState({app: data})
        }).catch(()=>{})
    }

    loadRoles(){
        GetData(`/app/${this.props.match.params.appId}/role`).then((data)=>{
            data = data || []
            this.setState({roleList: data})
        }).catch(()=>{})
    }

    tabsChange(key){
        let currentURL = this.props.match.url
        if(currentURL[currentURL.length - 1] === "/"){
            currentURL = currentURL.substring(0, currentURL.length - 1)
        }
        this.props.history.push(currentURL+key)
    }
    render(){
        let appId = this.props.match.params.appId
        let childPath = this.props.location.pathname.replace(this.props.match.url, "")
        let activeKey = childPath || "/"

        let path = this.props.match.url

        return (
        <>   
        <span>应用: {this.state.app.name}<Tag color="cyan" style={{marginLeft:"20px"}}>{this.getModeNameMap(this.state.app.model)}</Tag></span>
        <Divider />
        <RouteSwitch>
            <Route exact path={`${path}/app-role/:roleId`}>
                <RoleEdit/>
            </Route>
            <Route path={path}>
                <Tabs tabPosition="left" animated={false} onChange={(key)=>{this.tabsChange(key)}} activeKey={activeKey}>
                    <TabPane tab={<span><ControlOutlined />运行模式</span>} key="/">
                        <RunModel appId={appId} appModel={this.state.app.model} roleList={this.state.roleList}/>
                    </TabPane>
                    <TabPane tab={ <span><UsergroupDeleteOutlined />黑名单</span>} key="/app-black">
                        <UserList appId={appId} type="black" roleList={this.state.roleList} />
                    </TabPane>
                    <TabPane tab={ <span><TeamOutlined />白名单</span>} key="/app-white">
                        <UserList appId={appId} type="white" roleList={this.state.roleList} />
                    </TabPane>
                    <TabPane tab={ <span><UserOutlined />角色列表</span>} key="/app-role">
                        <RoleList appId={appId} type="role" roleList={this.state.roleList} />
                    </TabPane>
                </Tabs>
            </Route>
        </RouteSwitch>
        
       
        </>)
    }
}

export default withRouter(Page)