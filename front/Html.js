import React, {Component, PropTypes} from 'react';
import serialize from 'serialize-javascript';

/**
 * Wrapper component containing HTML metadata and boilerplate tags.
 * Used in server-side code only to wrap the string output of the
 * rendered route component.
 *
 * The only thing this component doesn't (and can't) include is the
 * HTML doctype declaration, which is added to the rendered output
 * by the server.js file.
 */
export default class Html extends Component {
  static propTypes = {
    webpackStats: PropTypes.object,
    component: PropTypes.object,
    store: PropTypes.object
  }

  render() {
    const {webpackStats, component, store} = this.props;
    const title = 'Grox';
    const reactHtml = React.renderToString(component);
    const {html, script} = cleanReactid('content', reactHtml);

    return (
      <html lang='en-us'>
        <head>
          <meta charSet='utf-8'/>
          <title>{title}</title>
          <link rel='shortcut icon' href='/favicon.ico' />
          {webpackStats.css.files.map((css, i) =>
            <link href={css} key={i} media='screen, projection'
                  rel='stylesheet' type='text/css'/>)}
        </head>
        <body>
          <div id='content' dangerouslySetInnerHTML={{__html: html}}/>
          <script dangerouslySetInnerHTML={{__html: `window.__data=${serialize(store.getState())};` + script}} />
          <script src={webpackStats.script[0]}/>
        </body>
      </html>
    );
  }
}

function cleanReactid(elemId, reactHtml) {
  const regex = /\ data-reactid="((\.[0-9A-z:$]+)+)"/g;
  let lastId = [];
  const tree = [];

  const html = reactHtml.replace(regex, (match, reactId) => {
    const id = reactId.split('.').slice(1);
    if (id.length > lastId.length) {
      const extra = id.slice(lastId.length);
      tree.push(extra.map(e => '.' + e));

    } else if (id.length === lastId.length) {
      tree.push(',' + id[id.length - 1]);

    } else {
      for (let i = 0, l = lastId.length - id.length; i < l; i++) {
        tree.push(';');
      }
      tree.push(id[id.length - 1]);
    }

    lastId = id;
    return '';
  });

  const script = `(function(){var h="${tree.join('')}".split(/(?=[.,;]+)/),i=[],k=document.getElementById("${elemId}").querySelectorAll("*");for(var l=0,m=k.length,q=0;l<m;l++,q++){var n=h[q],o=n[0],p=n.slice(1);if(o==".")i.push(p);else if(o==",")i[i.length-1]=p;else{i.pop();while(!p){i.pop();q++;p=h[q].slice(1)};i[i.length-1]=p};k[l].dataset.reactid="."+i.join(".")}})();`;

  return {html, script};
  // return {html: reactHtml, script: ''};
}
