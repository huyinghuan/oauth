import React from "react"

import {Table, notification, Popconfirm, Button} from "antd"

import { get as GetData } from '../../../../service'
import { FormOutlined, DeleteOutlined } from '@ant-design/icons';

export default class Components extends React.Component{
    constructor(props){
        super(props)
        this.initRoleNameMap()
        this.state = {
            loading: true,
            list: []
        }
    }
    initRoleNameMap(){
        let o = {0:"默认角色"}
        this.props.roleList.forEach((item)=>{
            o[item.id] = item.name
        })
        this.roleNameMap = o
    }
    del(id){
        GetData(`/app/${this.props.appId}/user/${id}`,{ method: "DELETE" }).then(()=>{
            notification.success({
                message: '操作成功',
                placement: 'topRight',
                duration: 3,
            });
            this.loadData()
        })
    }
    tableColumns = [
        {
            title: '用户名',
            dataIndex: ["user", "name"],
            key: "application_name",
            align: "center"
        },
        {
            title: '角色',
            dataIndex: ["appUser", "role_id"],
            key: "role_id",
            align: "center",
            render: (text)=>{
                return this.roleNameMap[text]
            }
        },
        {
            title: '操作',
            key:"action",
            align: "center",
            render:(text, record)=>{
              
                return (
                    <div>
                        <Popconfirm placement="topLeft" title="确认从名单中删除该用户?" onConfirm={()=>{this.del(record.appUser.id)}} okText="Yes" cancelText="No">
                            <Button danger icon={<DeleteOutlined />} type="link"  >删除</Button>
                        </Popconfirm>
                        <Button icon={<FormOutlined />} type="link" onClick={()=>{}}>分配角色</Button>
                        <Button icon={<FormOutlined />} type="link" onClick={()=>{}}>权限详情</Button>
                    </div>
                )
            }
        }
    ]
    componentDidMount(){
        this.loadData()
    }
    loadData(){
        let appId = this.props.appId
        let listType = this.props.type
        GetData(`/app/${appId}/user`).then((data)=>{
            data = data || []
            let targetList = []
            data.forEach((item)=>{
                if(item.appUser.category === listType){
                    // item.appUser["roleName"] = this.roleNameMap[item.appUser.role_id]
                    targetList.push(item)
                }
            })
            this.setState({list: targetList, loading: false})
        })
    }

    render(){
        return (
            <Table loading={this.state.loading}
                rowClassName="custom-row-strict"
                dataSource={this.state.list}
                columns={this.tableColumns}
                rowKey={(record)=>{return record.appUser.id}}
                pagination={{
                    defaultPageSize: 10,
                    showSizeChanger: true,
                    pageSizeOptions: [10,20,50]
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