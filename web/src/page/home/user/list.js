import React from 'react';
import { get as GetData } from '../../../service'
import { withRouter } from 'react-router-dom';
import { Table, Button, Popconfirm } from 'antd';

import { FormOutlined, DeleteOutlined } from '@ant-design/icons';


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
            title: '用户名',
            dataIndex: ["name"],
            key: "application_name",
            align: "center"
        },
        {
            title: '操作',
            key:"action",
            align: "center",
            render:(text, record)=>{
                let editHref = [this.props.match.path, record.id, "password-reset"].join("/")
                return (
                    <div>
                        <Popconfirm placement="topLeft" title="确认删除该用户?" onConfirm={()=>{this.del(record.id)}} okText="Yes" cancelText="No">
                            <Button danger icon={<DeleteOutlined />} type="link"  >删除</Button>
                        </Popconfirm>
                        <Button icon={<FormOutlined />} type="link" onClick={()=>{this.goto(editHref)}}>修改密码</Button>
                    </div>
                )
            }
        }
    ]
    del(key){
        GetData(`/user/${key}`,{ method: "DELETE" }).then(()=>{
            this.loadData()
        })
    }
    goto(href){
        this.props.history.push(href)
    }
    loadData(){
        GetData("/user").then((data)=>{
            this.setState({dataSource: data, loading: false})
        }).catch((e)=>{
            this.setState({loading: false})
        })
    }

    render() {
        return (
            <Table loading={this.state.loading}
                rowClassName="custom-row-strict"
                dataSource={this.state.dataSource}
                columns={this.tableColumns}
                rowKey={(record)=>{return record.id}}
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

export default withRouter(Page)