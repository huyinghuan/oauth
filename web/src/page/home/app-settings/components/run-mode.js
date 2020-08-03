import React from "react"
import { get as GetData } from '../../../../service'
import { Button, Modal, notification, Radio, Form, Input, Select} from 'antd';

import { UserAddOutlined} from '@ant-design/icons';

const {Option} = Select

export default class Components extends React.Component{
    constructor(props){
        super(props)
        this.moduleNameMap ={
            "white": "白名单",
            "black": "黑名单"
        }
        this.formRef = React.createRef()
    }
    componentDidMount(){
    }

    updateAppRunMode(m){
        GetData(`/app/${this.props.appId}/user_mode/${m}`,{method:"PATCH"}).then(()=>{
            notification.success({
                message: '操作成功',
                placement: 'topRight',
                duration: 3,
            });
            let app = this.state.app
            app.model = m
            this.setState({app: app})
        }).catch((e)=>{})
    }
    switchAppRunMode = (e)=>{
        let m = e.target.value
        Modal.confirm({
            title: `是否将应用运行模式变更为:${this.moduleNameMap[m]}?`,
            //icon: <ExclamationCircleOutlined />,
            content: `黑名单模式: 禁止名单内用户访问应用；白名单模式: 仅允许名单用户访问应用`,
            okText: '确认',
            cancelText: '取消',
            destroyOnClose: true,
            onOk: ()=>{
                this.updateAppRunMode(m)
            }
        });
    }
    addUserToCategoryList(v){
        console.log(v)
        GetData(`/app/${this.props.appId}/user`,{method: 'POST'}, v).then(()=>{
            notification.success({
                message: '操作成功',
                placement: 'topRight',
                duration: 3,
            });
            this.formRef.current.resetFields()
        }).catch((e)=>{
            console.log(e);
        })
    }
    render(){
        return(<>
              <Form>
                    <Form.Item label="切换模式">
                        <Radio.Group onChange={this.switchAppRunMode} value={this.props.appModel}>
                            <Radio value="black">黑名单</Radio>
                            <Radio value="white">白名单</Radio>
                        </Radio.Group>
                    </Form.Item>
                </Form>
                <Form layout="inline" 
                    onFinish={(v)=>{this.addUserToCategoryList(v)}}
                    initialValues={{category:"white", role_id:0}}
                    ref={this.formRef}
                    >
                    <Form.Item label="用户" name="username" rules={[{ required: true, message: '用户不能为空' }]}>
                        <Input />
                    </Form.Item>
                    <Form.Item  name="category">
                        <Select style={{ width: 120 }}>
                            <Option value="white">白名单</Option>
                            <Option value="black">黑名单</Option>
                        </Select>
                    </Form.Item>
                    <Form.Item  name="role_id">
                        <Select  style={{ width: 120 }} >
                            <Option value={0}>默认角色</Option>
                            {this.props.roleList.map((item)=>{
                                return (<Option value={item.id} key={item.id}>{item.name}</Option>)
                            })}
                        </Select>
                    </Form.Item>
                    <Form.Item>
                        <Button type="primary" htmlType="submit" icon={<UserAddOutlined/>}>添加</Button>
                    </Form.Item>
                </Form>
        </>)
    }
}