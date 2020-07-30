import React from "react"

import {Table, notification, Popconfirm, Button, Form, Select} from "antd"

import { get as GetData } from '../../../../service'
import { FormOutlined, DeleteOutlined, KeyOutlined } from '@ant-design/icons';


const EditableCell = ({
    roleList,
    editing,
    dataIndex,
    title,
    record,
    index,
    children,
    ...restProps
  }) => {
    return (
      <td {...restProps}>
        {editing ? (
          <Form.Item name={dataIndex} style={{margin: 0}}>
            <Select>
                <Select.Option value={0}>默认角色</Select.Option>
                { roleList.map((item)=>{
                    return  <Select.Option key={item.id} value={item.id}>{item.name}</Select.Option>
                })}
               
            </Select>
          </Form.Item>
        ) : (
          children
        )}
      </td>
    );
};

export default class Components extends React.Component{
    constructor(props){
        super(props)
        this.initRoleNameMap()
        this.state = {
            loading: true,
            list: [],
            editingKey: ""
        }
        this.cellFormRef = React.createRef()

    }
    isCellEditing = record => record.id === this.state.editingKey;
    cellEdit = (record) => {
        this.cellFormRef.current.setFieldsValue({
          name: '',
          ...record,
        });
        this.setState({
            editingKey: record.id
        })  
    };
    cellCancel = () => {
        this.setState({
            editingKey: ""
        })
    };
    
    cellSave = async (key) => {
        const row = await this.cellFormRef.current.validateFields();
        console.log(row)
    };

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
            },
            editable: true,
        },
        {
            title: '操作',
            key:"action",
            align: "center",
            render:(text, record)=>{
                let editable = this.isCellEditing(record);
                return !editable ?  (
                    <div>
                        <Popconfirm placement="topLeft" title="确认从名单中删除该用户?" onConfirm={()=>{this.del(record.appUse.id)}} okText="Yes" cancelText="No">
                            <Button danger icon={<DeleteOutlined />} type="link"  >删除</Button>
                        </Popconfirm>
                        <Button icon={<FormOutlined />} type="link" onClick={()=>{this.cellEdit(record)}}>分配角色</Button>
                        <Button icon={<FormOutlined />} type="link" onClick={()=>{}}>权限详情</Button>
                    </div>
                ) : (
                    <span>
                    <Button icon={<FormOutlined />} type="link" onClick={()=>{this.cellSave(record.key)}}>保存</Button>
                    <Popconfirm title="Sure to cancel?" onConfirm={this.cellCancel}>
                        <Button icon={<FormOutlined />} type="link" >取消</Button>
                    </Popconfirm>
                </span>
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
        }).catch(()=>{})
    }

    render(){

        const mergedColumns = this.tableColumns.map(col => {
            if (!col.editable) {
              return col;
            }
        
            return {
              ...col,
              onCell: record => ({
                record,
                roleList: this.props.roleList,
                dataIndex: col.dataIndex,
                title: col.title,
                editing: this.isCellEditing(record),
              }),
            };
        });


        return (
            <Form ref={this.cellFormRef} component={false}>
            <Table loading={this.state.loading}
                rowClassName="custom-row-strict"
                dataSource={this.state.list}
                columns={mergedColumns}
                rowKey={(record)=>{return record.appUser.id}}
                components={{
                    body: {
                      cell: EditableCell,
                    },
                }}
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
            </Form>
        )
       
    }

}