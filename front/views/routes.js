import React from 'react';
import {Route} from 'react-router';
import App from 'views/App';
import Home from 'views/Home';
import Newest from 'views/Newest';
import Submit from 'views/Submit';
import NotFound from 'views/NotFound';

export default (
  <Route component={App}>
    <Route path='/' component={Home}/>
    <Route path='/newest' component={Newest}/>
    <Route path='/submit' component={Submit}/>
    <Route path='*' component={NotFound}/>
  </Route>
);
