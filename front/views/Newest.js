import React, {Component} from 'react';
import {gcx} from 'utils';

const items = Array.apply(null, Array(10)).map((x, i) => ({
  title: 'Sample',
  url: 'http://grox.golang.vn',
  score: 10 * i,
  comments: 5 * i
}));

export default
class Newest extends Component {
  render() {
    return (
      <div className={gcx('ui container')}>
        <div className={gcx('ui divided very relaxed list')}>
          {
            items.map((item, i) => (
              <div className={gcx('item')} key={i}>
                <i className={gcx('chevron up icon')}></i>
                <div className={gcx('content')}>
                  <div className={gcx('header')}>
                    {item.title}
                  </div>
                  {item.score} points | {item.comments} comments
                </div>
              </div>
            ))
          }
          <div className={gcx('item')}/>
        </div>
      </div>
    );
  }
}
