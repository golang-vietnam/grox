import classNames from 'classnames';
export const cx = classNames;

let __stats;

export function readStats() {
  // don't cache the `webpack-stats.json` on dev so we read the file on each request.

  if (!__DEVELOPMENT__ && __stats) {
    return __stats;
  }

  try {
    const fs = require('fs');
    const path = require('path');

    const data = fs.readFileSync(path.join(__dirname, '../webpack-stats.json')).toString();
    __stats = JSON.parse(data);
    return __stats;

  } catch (e) {
    console.log(e);
    console.log('You might see the error: webpack-stats.json does not exist. Don\'t worry, just wait until webpack finish running then make a change to reload the node server.');
    return {
      css: { modules: {} }
    };
  }
}

export function requireServerCss(cssPath) {
  if (__CLIENT__) {
    throw new Error('server-side only css resolver called on client');
  }
  return readStats().css.modules[cssPath];
}

export function requireServerImage(imagePath) {
  if (!imagePath) {
    return '';
  }
  if (__CLIENT__) {
    throw new Error('server-side only image resolver called on client');
  }
  const images = readStats().images;
  if (!images) {
    return '';
  }

  // Find the correct image
  const regex = new RegExp(`${imagePath}$`);
  const image = images.find(img => regex.test(img.original));

  // Serve image.
  if (image) return image.compiled;

  // Serve a not-found asset maybe?
  return '';
}

const vendorStyles = __CLIENT__ ?
  require('styles/vendor.less') :
  requireServerCss('styles/vendor.less');

const appStyles = __CLIENT__ ?
  require('styles/main.less') :
  requireServerCss('styles/main.less');

export const styles = Object.assign({}, vendorStyles, appStyles);

export function gcx(classes, ...args) {
  return cx(classes.split(' ').map(i => styles[i] || i), ...args);
}
