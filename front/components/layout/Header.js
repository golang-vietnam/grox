import React from 'react';
import {gcx, requireServerCss} from 'utils';

const CSS = __CLIENT__ ? require('components/layout/layout.less') : requireServerCss('components/layout/layout.less');

export default
class CounterButton extends React.Component {
  render() {
    return (
      <div className={gcx('ui main pointing menu')} style={{borderLeft: 'none !important', borderRight: 'none !important', borderRadius: 0}}>
        <div className={gcx('ui container')}>
          <div href='#' className={gcx('header item')}>
            <span className={CSS.logo}/>
            Grox
          </div>
          <a href='/' className={gcx('active item')}>Front</a>
          <a href='/newest' className={gcx('item')}>New</a>
          <div className={gcx('right menu')}>
            <a href='#' className={gcx('item')}>
              <i className={gcx('fa fa-plus-circle')}/>
              &nbsp;Submit
            </a>
            <a href='#' className={gcx('item')} tabIndex='0'>
              <span className={CSS.githubLogo}/>
              &nbsp;Login
            </a>
            <div className={gcx('ui floated simple dropdown item')} tabIndex='0'>
              <i className={gcx('fa fa-caret-down')}/>
              <div className={gcx('menu')} tabIndex='-1'>
                <div className={gcx('item')}>Send feedback!</div>
                <a className={gcx('item')} href='https://github.com/golang-vietnam/grox'>Grox on GitHub</a>
                <div className={gcx('divider')}></div>
                <div className={gcx('header')}>Customize</div>
                <div className={gcx('item')}>
                  <i className={gcx('dropdown icon')}></i>
                  Themes
                  <div className={gcx('menu')} style={{left: 'auto !important', right: '97% !important', top: '4px !important'}}>
                    <a className={gcx('item')} href='?theme=default'>
                      <span className={CSS.iconS} />
                      Default
                    </a>
                    <a className={gcx('item')} id={CSS.hackerNews} href='?theme=hacker-news'>
                      <span className={CSS.iconY} />
                      Hacker News
                    </a>
                  </div>
                </div>
                <div className={gcx('item')}>
                  <i className={gcx('united states flag')}></i>
                  English
                </div>
                <div className={gcx('item')}>
                  <i className={gcx('vietnam flag')}></i>
                  Vietnamese
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  }
}
