import React from 'react';
import { get as GetData } from '../../../service'
import { withRouter } from 'react-router-dom';
import { Button, Form, Layout, Input, Modal, notification} from 'antd';

import { ExclamationCircleOutlined } from '@ant-design/icons';

const {Content} = Layout



class Page extends React.Component{
    formRef = React.createRef();

    componentWillUnmount(){
        Modal.destroyAll()
    }
    onFinish(v){
        let api = "/user/password"
        let uid = this.props.match.params.id
        let needRelogin = true
        if(uid){
            api = `/user/${uid}/password`
            needRelogin = false
        }
        GetData(api, {
            method: "PUT",
            headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
            },
            body: JSON.stringify(v)
        }).then(()=>{
            if(!needRelogin){
                notification.success({
                    message: '操作成功',
                    placement: 'topRight',
                    duration: 3,
                });
                return
            }
            Modal.confirm({
                title: '密码重置成功',
                icon: <ExclamationCircleOutlined />,
                content: '请重新登录',
                okText: '确认',
                cancelText: '重新登录',
                destroyOnClose: true,
                onOk: ()=>{
                    const {location} = this.props.history
                    this.props.history.push(`/login?goback=${encodeURIComponent(location.pathname+location.search)}`)
                },
                onCancel: ()=>{
                    const {location} = this.props.history
                    this.props.history.push(`/login?goback=${encodeURIComponent(location.pathname+location.search)}`)
                }
            });
        }).catch((e)=>{})
    }
    render(){
        let uid = this.props.match.params.id
        console.log(this.props.match.params)
        let setForOther = false
        if(uid){
            setForOther = true
        }
        
        return (
            <Layout>
                <Content style={{background: '#fff'}}>
                    <Form
                        ref={this.formRef}
                        name="basic"
                        onFinish={(values)=>{this.onFinish(values)}}
                        labelCol={{ span: 2 }}
                        wrapperCol={{span: 4}} >

                        { setForOther ? null : (
                            <Form.Item label="旧密码" name="oldPassword"
                                rules={[{ required: true, message: '旧密码不能为空' }]} >
                                <Input.Password />
                            </Form.Item>
                        )}
                        

                        <Form.Item label="新密码" name="newPassword"
                            rules={[{ required: true, message: '新密码不能为空' }]}>
                            <Input.Password />
                        </Form.Item>
                        <Form.Item label="确认密码" name="newPassword2" dependencies={['newPassword']}
                            rules={[
                                { required: true, message: '确认密码不能为空' },
                                ({ getFieldValue }) => ({
                                    validator(rule, value) {
                                      if (!value || getFieldValue('newPassword') === value) {
                                        return Promise.resolve();
                                      }
                                      return Promise.reject('两次密码不一致');
                                    },
                                }),
                            ]}>
                            <Input.Password />
                        </Form.Item>
                        <Form.Item wrapperCol={{ offset: 2, span: 4 }}>
                            <Button type="primary" htmlType="submit">保存</Button>
                        </Form.Item>
                    </Form>
                </Content>
          </Layout>
        )
    }
}


export default withRouter(Page)