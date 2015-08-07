import React from 'react';
import {gcx} from 'utils';

export default
class CounterButton extends React.Component {
  render() {
    return (
      <div className={gcx('ui main menu')}>
        <div className={gcx('ui container')}>
          <div href='#' className={gcx('header item')}>
            <img className={gcx('logo')}/>
            Grox
          </div>
          <a href='#' className={gcx('item')}>Blog</a>
          <a href='#' className={gcx('item')}>Articles</a>
          <a href='#' className={gcx('ui right floated dropdown item')} tabIndex='0'>
            Dropdown <i className={gcx('dropdown icon')}></i>
            <div className={gcx('menu')} tabIndex='-1'>
              <div className={gcx('item')}>Link Item</div>
              <div className={gcx('item')}>Link Item</div>
              <div className={gcx('divider')}></div>
              <div className={gcx('header')}>Header Item</div>
              <div className={gcx('item')}>
                <i className={gcx('dropdown icon')}></i>
                Sub Menu
                <div className={gcx('menu')}>
                  <div className={gcx('item')}>Link Item</div>
                  <div className={gcx('item')}>Link Item</div>
                </div>
              </div>
              <div className={gcx('item')}>Link Item</div>
            </div>
          </a>
        </div>
      </div>
    );
  }
}
