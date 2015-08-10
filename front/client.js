/* global __DEVTOOLS__ */
require('./styles/main.less');

import React from 'react';
import BrowserHistory from 'react-router/lib/BrowserHistory';
import Location from 'react-router/lib/Location';
import createStore from './redux/create';
import ApiClient from './ApiClient';
import universalRouter from './universalRouter';

const history = new BrowserHistory();
const client = new ApiClient();

const dest = document.getElementById('content');
const store = createStore(client, window.__data);
const location = new Location(document.location.pathname, document.location.search);
universalRouter(location, history, store)
  .then(({component}) => {
    if (__DEVTOOLS__) {
      window.__store = store;

      const initState = JSON.stringify(store.getState());
      const states = window.__states = [JSON.parse(initState)];
      console.log('state', initState);

      store.subscribe(() => {
        const state = JSON.stringify(store.getState());
        states.push(JSON.parse(state));
        console.log('state', state);
      });
    }

    React.render(component, dest);
  }, (error) => {
    console.error(error);
  });


if (process.env.NODE_ENV !== 'production') {
  window.React = React; // enable debugger
}
