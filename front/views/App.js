import React, {Component, PropTypes} from 'react';
import {createTransitionHook} from '../universalRouter';
import Header from '../components/layout/Header';

import {gcx} from 'utils';

class App extends Component {
  static propTypes = {
    children: PropTypes.element.isRequired
  }

  static contextTypes = {
    router: PropTypes.object.isRequired,
    store: PropTypes.object.isRequired
  };

  componentWillMount() {
    const {router, store} = this.context;
    router.addTransitionHook(createTransitionHook(store));
  }

  render() {
    return (
      <div className={gcx('container')}>
        <Header {...this.props}/>
        {this.props.children}
      </div>
    );
  }
}

export default
class AppContainer {
  static propTypes = {
    children: PropTypes.element.isRequired
  }

  render() {
    return (
      <App {...this.props}>
        {this.props.children}
      </App>
    );
  }
}
