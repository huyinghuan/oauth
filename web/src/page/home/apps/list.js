import React from 'react';
import { get as GetData } from '../../../service'
import {Route, Switch, withRouter, Link} from 'react-router-dom';
import { Table, Button, Popconfirm } from 'antd';

import { TeamOutlined, FormOutlined, DeleteOutlined, EyeOutlined} from '@ant-design/icons';


class Page extends React.Component {
    constructor(props){
        super(props)
        this.state = {
            dataSource: [],
            loading: true
        }
        this.loadData()
    }
    tableColumns = [
        {
            title: '应用名称',
            dataIndex: ["application", "name"],
            key: "application_name"
        },
        {
            title: '创建人',
            dataIndex: ["user", "name"],
            key: "user_name"
        },{
            title: 'Client ID',
            dataIndex: ["application", "client_id"],
            key:"application_client_id"
        },{
            title: "Private Key",
            dataIndex: ["application", "private_key"],
            key:"application_client_id"
        },{
            title: '回调地址',
            dataIndex: ["application", "callback"],
            key:"callback"
        },{
            title: '操作',
            key:"action",
            width: 360,
            render:(text, record)=>{
                let editHref = [this.props.match.path, record.application.id].join("/")
                let userManagerHref = [this.props.match.path, record.application.id, "userManager"].join("/")
                return (
                    <div className="custom-btn-group">
                        <Popconfirm placement="topLeft" title="确认删除该应用?" onConfirm={()=>{this.delApp(record.application.id)}} okText="Yes" cancelText="No">
                            <Button danger icon={<DeleteOutlined />}  >删除</Button>
                        </Popconfirm>
                        <Button icon={<FormOutlined />} type="primary" onClick={()=>{this.goto(editHref)}}>编辑</Button>
                        <Button icon={<TeamOutlined/>} onClick={()=>{this.goto(userManagerHref)}}>用户管理</Button>
                    </div>
                )
            }
        }
    ]
    delApp(key){
        console.log(key)
        GetData(`/app/${key}`,{ method: "DELETE" }).then(()=>{
            this.loadData()
        })
    }
    goto(href){
        this.props.history.push(href)
    }
    loadData(){
        GetData("/app").then((data)=>{
            this.setState({dataSource: data, loading: false})
        }).catch((e)=>{
            this.setState({loading: false})
        })
    }

    render() {
        return (
            <Table loading={this.state.loading}
                dataSource={this.state.dataSource}
                columns={this.tableColumns}
                rowKey={(record)=>{return record.application.id}}
                pagination={{
                    defaultPageSize: 15,
                    hideOnSinglePage: true,
                    showSizeChanger: true,
                    pageSizeOptions: [15,30,50]
                }}
                scroll={{
                    y: window.innerHeight - 290,
                    // x: "100%",
                    // scrollToFirstRowOnChange: true
                }}
            />
        )
    }
}

export default withRouter(Page)