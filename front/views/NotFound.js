import React, {Component} from 'react';
import {gcx} from 'utils';

export default class NotFound extends Component {
  render() {
    return (
      <div className={gcx('ui container')}>
        <h1>404</h1>
        <p>This page does not exist!</p>
      </div>
    );
  }
}
