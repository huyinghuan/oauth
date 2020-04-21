import React from 'react';
import { HashRouter as Router, Route, Redirect } from 'react-router-dom';
import Login from './page/login'
import Register from './page/register'
import Home from './page/home'

function App() {
  return (
    <Router>
      <Route exact path="/login" component={Login} />
      <Route exact path="/register" component={Register} />
      <Route path="/home" component={Home}/>

      <Route exact path="/">
        <Redirect to="/home"/>
      </Route>
    </Router>
  );
}


export default App;