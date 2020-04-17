import React from 'react';


import { Layout,  Row, Col} from 'antd';
import RegisterFrom from "../../../components/register-form"
import { withRouter} from "react-router-dom"

const {Content} = Layout

class Page extends React.Component {
  onFinish(){
    // this.props.history.push("/login")
  };

  render() {
    return (
        <Layout>
            <Content style={{background: '#fff'}}>
            <Row>
                 <Col span={6}>
                    
                    <RegisterFrom onFinish={()=>{this.onFinish()}} />
                 </Col>
            </Row>
          </Content>
        </Layout>
    )
  }
}

const p = withRouter(Page)

export default p