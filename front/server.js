import Express from 'express';
import React from 'react';
import Location from 'react-router/lib/Location';
import config from './config';
import favicon from 'serve-favicon';
import compression from 'compression';
import httpProxy from 'http-proxy';
import path from 'path';
import createStore from './redux/create';
import ApiClient from './ApiClient';
import universalRouter from './universalRouter';
import Html from './Html';
import PrettyError from 'pretty-error';

const pretty = new PrettyError();
const app = new Express();
const proxy = httpProxy.createProxyServer({
  target: 'http://localhost:' + config.apiPort + config.apiPath
});

app.disable('x-powered-by');
app.use(compression());
app.use(favicon(path.join(__dirname, 'static', 'favicon.ico')));

let webpackStats;

if (!__DEVELOPMENT__) {
  webpackStats = require('./webpack-stats.json');
}

app.use(require('serve-static')(path.join(__dirname, 'static')));

// Proxy to API server
app.use(config.apiPath, (req, res) => {
  proxy.web(req, res);
});

app.use((req, res) => {
  if (__DEVELOPMENT__) {
    webpackStats = require('./webpack-stats.json');
    // Do not cache webpack stats: the script file would change since
    // hot module replacement is enabled in the development env
    delete require.cache[require.resolve('./webpack-stats.json')];
  }
  const client = new ApiClient(req);
  const store = createStore(client);
  const location = new Location(req.path, req.query);
  if (__DISABLE_SSR__) {
    res.send('<!doctype html>\n' +
      React.renderToStaticMarkup(<Html webpackStats={webpackStats} store={store}/>));
  } else {
    universalRouter(location, undefined, store)
      .then(({component, transition, isRedirect}) => {
        if (isRedirect) {
          res.redirect(transition.redirectInfo.pathname);
          return;
        }
        res.send('<!doctype html>\n' +
          React.renderToStaticMarkup(<Html webpackStats={webpackStats} component={component} store={store}/>));
      })
      .catch((error) => {
        console.error('ROUTER ERROR:', pretty.render(error));
        res.status(500).send({error: error.stack});
      });
  }
});

if (config.port) {
  app.listen(config.port, (err) => {
    if (err) {
      console.error(err);
    } else {
      console.info('==> Server is listening');
      console.info('==>   %s running on port %s', config.app.name, config.port);
      console.info('==>   %s proxy to http://localhost:%s', config.apiPath, config.apiPort);

      if (__DISABLE_SSR__) {
        console.info('==>   Server side rendering is disabled');
      } else {
        console.info('==>   `export DISABLE_SSR=1` to disable server side rendering');
      }
    }
  });
} else {
  console.error('==> ERROR: No FRONT_PORT environment variable has been specified');
}
