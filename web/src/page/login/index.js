import React from 'react';
import { get as GetData } from '../../service'

import { Form, Input, Button, Checkbox, Row, Col } from 'antd';

import { NavLink, withRouter} from "react-router-dom"

import "./index.css"

class NormalLoginForm extends React.Component {
  render() {
    return (
      <Form onFinish={this.props.onFinish}
        onFinishFailed={this.props.onFinishFailed}
      >
        <Form.Item name="username" rules={[{ required: true, message: '请输入用户名' }]}>
          <Input placeholder="用户名" autoComplete="username"/>
        </Form.Item>
        <Form.Item name="password" rules={[{ required: true, message: '请输入密码' }]}>
          <Input.Password placeholder="密码" autoComplete="current-password"/>
        </Form.Item>
        <Form.Item name="remember" valuePropName="checked">
          <Checkbox>Remember me</Checkbox>
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit">登录</Button>
          <Button type="link"><NavLink to="/users">忘记密码?</NavLink></Button>
          <Button type="link"><NavLink to="/users">注册</NavLink></Button>
        </Form.Item>
      </Form>
    );
  }
}
// const WrappedNormalLoginForm = Form.create({ name: 'normal_login' })(NormalLoginForm);

class Page extends React.Component {
  onFinish(values){
    GetData('/user/login', {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(values)
    }).then((resp) => {
      console.log(this.props)
      this.props.history.push("/home/apps")
    }).catch((e)=>{console.log(e)})
  };
  onFinishFailed(errorInfo) {
    console.log('Failed:', errorInfo);
  }
  render() {
    return (
      <div className="page-login">
        <Row>
          <Col span={10}></Col>
          <Col span={4}>
            <NormalLoginForm onFinish={(values)=>{this.onFinish(values)}} onFinishFailed={(errorInfo)=>{this.onFinishFailed(errorInfo)}}/>
          </Col>
          <Col span={10}></Col>
        </Row>
      </div>

    )
  }
}

const p = withRouter(Page)

export default p