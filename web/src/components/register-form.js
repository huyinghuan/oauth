
import React from "react"
import { Form, Input, Button,  notification} from 'antd';
import { get as GetData } from '../service'

export default class Component extends React.Component {
    onFinish(values){
        GetData('/user/register', {method: 'POST'}, values).then((resp)=>{
          notification.success({
            message: '注册作成功',
            placement: 'topRight',
            duration: 3,
          });
          // this.props.history.push("/login")
          this.props.onFinish()
        }).catch((e)=>{})
      };
    render() {
        return (
        <Form onFinish={(values) => { this.onFinish(values) }} labelCol={{ span: 8 }} wrapperCol={{ span: 16 }}>
            <Form.Item label="用户名" name="username" rules={[{ required: true, message: '请输入用户名' }]}>
                <Input autoComplete="username" />
            </Form.Item>
            <Form.Item label="密码" name="password" rules={[{ required: true, message: '请输入密码' }]}>
                <Input.Password autoComplete="current-password" />
            </Form.Item>
            <Form.Item label="确认密码" name="confirmPassword" dependencies={['password']}
                rules={[
                    { required: true, message: '确认密码不能为空' },
                    ({ getFieldValue }) => ({
                        validator(rule, value) {
                            if (!value || getFieldValue('password') === value) {
                                return Promise.resolve();
                            }
                            return Promise.reject('两次密码不一致');
                        },
                    }),
                ]}>
                <Input.Password />
            </Form.Item>
            <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
                <Button type="primary" htmlType="submit">注册</Button>
            </Form.Item>
        </Form>)
    }
}