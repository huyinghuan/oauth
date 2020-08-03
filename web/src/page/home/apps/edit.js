import React from 'react';
import { get as GetData } from '../../../service'
import { withRouter } from 'react-router-dom';
import { Button, Form, Layout, Input, Typography, Switch, notification} from 'antd';


const { Title } = Typography;
const {Content} = Layout



class Page extends React.Component{
    formRef = React.createRef();
    constructor(props){
        super(props)
        this.state = this.getStateData()   
        this.loadData(this.state.isEdit)
    }
    getStateData(){
        let isEdit = false
        if(this.props.match.params.appId){
            isEdit = true
        }
        return {
            appId: this.props.match.params.appId,
            title: isEdit ? "应用修改" : "应用注册",
            btnTitle: isEdit ? "保存" : "提交",
            isEdit: isEdit
        }
    }
    loadData(isEdit){
        if(!isEdit){
            return
        }
        GetData(`/app/${this.props.match.params.appId}`, {method:"GET"}).then((data)=>{
           this.formRef.current.setFieldsValue(data)
           console.log(data);
        }).catch((e)=>{
            this.formRef.current.resetFields()
        })
    }
    onFinish(v){
        let api = "/app"
        let method = "POST"
        if(this.state.isEdit){
            api = api + "/" + this.props.match.params.appId
            method = "PUT"
        }else{
            api = api + "/register"
        }
        GetData(api, {
            method: method,
            headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
            },
            body: JSON.stringify(v)
        }).then(()=>{
            notification.success({
                message: '操作成功',
                placement: 'topRight',
                duration: 3,
            });
            if(!this.state.isEdit){
                this.formRef.current.resetFields()
            }
        }).catch((e)=>{})
    }
    componentDidUpdate(){
        console.log(this.props.match.url)
        let data = this.getStateData()
        // 有编辑页面 -> 创建页面
        if(!data.appId && this.state.appId !== data.appId){
            this.formRef.current.resetFields()
            this.setState(data)
            // 编辑A -> 编辑 B
        }else if(data.appId && this.state.appId !== data.appId){
            this.loadData(data.isEdit)
        }
    }
    render(){
        return (
            <Layout>
                <Content style={{background: '#fff'}}>
                    <Title level={3}>{this.state.title}</Title>
                    <Form
                        ref={this.formRef}
                        name="basic"
                        // initialValues={this.state.app}
                        onFinish={(values)=>{this.onFinish(values)}}
                        // onFinishFailed={onFinishFailed}
                        wrapperCol={{span: 6}} >
                        <Form.Item label="应用名称" name="name"
                            rules={[{ required: true, message: '应用名称不能为空' }]} >
                            <Input />
                        </Form.Item>

                        <Form.Item label="回调地址" name="callback"
                            rules={[{ required: true, message: '回调地址不能为空' }]}>
                            <Input />
                        </Form.Item>
                        <Form.Item label="是否公开" name="open" valuePropName="checked">
                            <Switch />
                        </Form.Item>
                        <Form.Item>
                            <Button type="primary" htmlType="submit">{this.state.btnTitle}</Button>
                            { this.state.isEdit ? null : (<Button type="link" href="https://github.com/huyinghuan/oauth/blob/master/README.md" >如何使用?</Button>)}
                        </Form.Item>
                    </Form>
                </Content>
          </Layout>
        )
    }
}


export default withRouter(Page)