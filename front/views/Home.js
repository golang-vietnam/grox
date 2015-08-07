import React, {Component} from 'react';
import CounterButton from '../components/CounterButton';

import {styles, gcx} from 'utils';

export default class Home extends Component {
  render() {
    return (
      <div className={gcx('ui container')}>
        <h1>Home</h1>

        <CounterButton/>
      </div>
    );
  }
}
