import React, { Component } from 'react'
import { withRouter } from 'react-router-dom';
import { Button, Form,  Input, Table, Popconfirm, Divider, notification} from 'antd';
import { get as GetData } from '../../../service'
import { DeleteOutlined } from '@ant-design/icons';
class Page extends Component {
    constructor(props){
        super(props)
        // this.formRef = React.createRef();
        // this.cellFormRef = React.createRef()
        this.state = {
            LimitList: [],
            loading:true
        }
    }
    del(id){
        GetData(`/app/${this.props.appId}/role/${this.props.match.params.roleId}/permission/${id}`,{ method: "DELETE" }).then(()=>{
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
            title: '权限',
            dataIndex: ["pattern"],
            key: "pattern",
            align: "center"
        },
        {
            title: 'HTTP Method',
            dataIndex: ["method"],
            key: "method",
            align: "center"
        },
        {
            title: '备注',
            dataIndex: ["name"],
            key: "name",
            align: "center"
        },
        {
            title: '操作',
            key:"action",
            align: "center",
            render:(_, record)=>{
                return (
                    <div>
                        <Popconfirm placement="topLeft" title="确认删除该权限?" onConfirm={()=>{this.del(record.id)}} okText="Yes" cancelText="No">
                            <Button danger icon={<DeleteOutlined />} type="link"  >删除</Button>
                        </Popconfirm>
                    </div>
                )
            }
        }
    ]
    loadData(){
        let appId = this.props.appId
        let roleId = this.props.match.params.roleId
        GetData(`/app/${appId}/role/${roleId}/permission`).then((data)=>{
            data = data || []
            this.setState({LimitList: data,loading:false})
        })
    }
    onFinish(v){
        let roleId = this.props.match.params.roleId
        GetData(`/app/${this.props.appId}/role/${roleId}/permission`, {method:"POST"}, v).then(()=>{
            this.loadData()
        })
    }
    componentDidMount(){
        this.loadData()
    }
    render() {
        return (
            <div>
                <Form
                    name="AddLimit"
                    layout="inline"
                    onFinish={(values)=>{this.onFinish(values)}}
                    size='middle'
                    
                >
                    <Form.Item label="URL正则" name="pattern" labelCol={{span: 8}}>
                        <Input placeholder="url正则或字符串"/>
                    </Form.Item>
                    <Form.Item label="HTTP Method" name="method" labelCol={{span: 10}}>
                        <Input placeholder="HTTP Method多个用','隔开"/>
                    </Form.Item>
                    <Form.Item label="备注" name="name" labelCol={{span: 8}}>
                        <Input placeholder="备注"/>
                    </Form.Item>
                    <Form.Item>
                        <Button type="primary" htmlType="submit">
                        Submit
                        </Button>
                    </Form.Item>
                </Form>
                <Divider 
                    orientation='left'
                >权限列表</Divider>
                <Table 
                    loading={this.state.loading}
                    columns={this.tableColumns}
                    dataSource={this.state.LimitList}
                    rowKey={(record)=>{return record.id}}
                />
            </div>
        )
    }
}
export default withRouter(Page)