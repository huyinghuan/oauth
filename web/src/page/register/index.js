import React from 'react';


import { Row, Col } from 'antd';
import RegisterFrom from "../../components/register-form"
import { withRouter} from "react-router-dom"

import "./index.css"

class Page extends React.Component {
  onFinish(){
    this.props.history.push("/login")
  };

  render() {
    return (
      <div className="page-register">
        <Row>
          <Col span={10}></Col>
          <Col span={4}>
            <RegisterFrom onFinish={()=>{this.onFinish()}} />
          </Col>
          <Col span={10}></Col>
        </Row>
      </div>

    )
  }
}

const p = withRouter(Page)

export default p