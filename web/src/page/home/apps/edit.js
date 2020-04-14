import React from 'react';
import { get as GetData } from '../../../service'
import {Route, Switch, withRouter, Link} from 'react-router-dom';
import { Table, Button, Form, Layout, Input, Typography} from 'antd';

import { TeamOutlined, FormOutlined, DeleteOutlined, EyeOutlined} from '@ant-design/icons';

const { Title } = Typography;
const {Content} = Layout



class Page extends React.Component{

    render(){
        return (
            <Layout>
                <Content style={{background: '#fff'}}>
                    <Title level={3}>应用修改</Title>
                    <Form
                        name="basic"
                        // initialValues={{ remember: true }}
                        // onFinish={onFinish}
                        // onFinishFailed={onFinishFailed}
                        >
                        <Form.Item
                            label="应用名称"
                            name="name"
                            rules={[{ required: true, message: '应用名称不能为空' }]}
                        >
                            <Input />
                        </Form.Item>

                        <Form.Item
                            label="回调地址"
                            name="callback"
                            rules={[{ required: true, message: '回调地址不能为空' }]}
                        >
                            <Input />
                        </Form.Item>

                        <Form.Item>
                            <Button type="primary" htmlType="submit">
                            Submit
                            </Button>
                        </Form.Item>
                        </Form>
                </Content>
          </Layout>
        )
    }
}


export default withRouter(Page)