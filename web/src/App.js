import React from 'react';
import { HashRouter as Router, Route } from 'react-router-dom';

import Login from './page/login'
import Home from './page/home'
function App() {
  return (
    <Router>
      <Route exact path="/login" component={Login}></Route>
      <Route path="/home">
        <Home />
      </Route>
    </Router>
  );
}

export default App;