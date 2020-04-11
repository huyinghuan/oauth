import React from 'react';
import { HashRouter as Router, Route, Redirect } from 'react-router-dom';

import Login from './page/login'
import Home from './page/home'
function App() {
  return (
    <Router>
      <Route exact path="/login">
        <Login />
      </Route>
      <Route path="/home">
        <Home />
      </Route>
      {/* <Route strict path="/">
        <Redirect to="/home"/>
      </Route> */}
    </Router>
  );
}

export default App;