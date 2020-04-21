import React from "react"

import {Table,  Button, Modal, notification, Popconfirm, Form, Input, Divider} from "antd"

import { KeyOutlined, FormOutlined, DeleteOutlined } from '@ant-design/icons';

import { get as GetData } from '../../../../service'


const EditableCell = ({
    editing,
    dataIndex,
    title,
    inputType,
    record,
    index,
    children,
    ...restProps
  }) => {
    return (
      <td {...restProps}>
        {editing ? (
          <Form.Item
            name={dataIndex}
            style={{
              margin: 0,
            }}
            rules={[
              {
                required: true,
                message: `Please Input ${title}!`,
              },
            ]}
          >
            <Input />
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
        this.formRef = React.createRef();
        this.cellFormRef = React.createRef()

        this.state = {
            loading: true,
            list: [],
            editingKey:""
        }
    }
    isEditing = record => record.id === this.state.editingKey;
    edit = record => {
        this.cellFormRef.current.setFieldsValue({
          name: '',
          ...record,
        });
        this.setState({
            editingKey: record.id
        })  
    };
    cancel = () => {
        this.setState({
            editingKey: ""
        })
    };
    
    save = async key => {
        try {
            // TODO
            const row = await this.cellFormRef.current.validateFields();
            console.log(row)
            

        } catch (errInfo) { }
    };
    tableColumns = [
        {
            title: '角色名',
            dataIndex: ["name"],
            key: "name",
            align: "center",
            editable: true,
        },
        {
            title: '操作',
            key:"action",
            align: "center",
            render:(_, record)=>{
                let editable = this.isEditing(record);

                return !editable ? (
                    <div>
                        <Popconfirm placement="topLeft" title="确认删除该角色?" onConfirm={()=>{this.del(record.id)}} okText="Yes" cancelText="No">
                            <Button danger icon={<DeleteOutlined />} type="link"  >删除</Button>
                        </Popconfirm>
                        {/* <Button icon={<FormOutlined />} type="link" onClick={()=>{this.edit(record)}}>修改</Button> */}
                        <Button icon={<KeyOutlined />} type="link" onClick={()=>{}}>权限分配</Button>
                    </div>
                ):(
                    <span>
                        <Button icon={<FormOutlined />} type="link" onClick={()=>{this.save(record.key)}}>保存</Button>
                        <Popconfirm title="Sure to cancel?" onConfirm={this.cancel}>
                            <Button icon={<FormOutlined />} type="link" >取消</Button>
                        </Popconfirm>
                    </span>
                )
            }
        }
    ]

    del(id){
        GetData(`/app/${this.props.appId}/role/${id}`,{ method: "DELETE" }).then(()=>{
            notification.success({
                message: '操作成功',
                placement: 'topRight',
                duration: 3,
            });
            this.loadData()
        })
    }

    componentDidMount(){
        this.loadData()
    }

    loadData(){
        let appId = this.props.appId
        GetData(`/app/${appId}/role`).then((data)=>{
            data = data || []
            this.setState({list: data, loading: false})
        })
    }

    onFinish(v){
        GetData(`/app/${this.props.appId}/role`, {method:"POST"}, v).then(()=>{
            this.loadData()
            this.formRef.current.resetFields()
        })

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
                inputType: "",
                dataIndex: col.dataIndex,
                title: col.title,
                editing: this.isEditing(record),
              }),
            };
          });
        return (
            <>
            <Form layout="inline" ref={this.formRef}
                        name="basic"
                        onFinish={(values)=>{this.onFinish(values)}}>
                <Form.Item label="角色" name="role" rules={[{ required: true, message: '角色名称不能为空' }]}>
                    <Input />
                </Form.Item>
                <Form.Item>
                    <Button type="primary" htmlType="submit">添加</Button>
                </Form.Item>
            </Form>
            <Divider />
            <Form ref={this.cellFormRef} component={false}>
            <Table loading={this.state.loading}
                bordered
                rowClassName="custom-row-strict"
                dataSource={this.state.list}
                columns={mergedColumns}
                rowKey={(record)=>{return record.id}}
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
            </>
        )
       
    }

}