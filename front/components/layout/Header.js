import React, {PropTypes} from 'react';
import {Link} from 'react-router';
import {gcx, styles, requireServerCss} from 'utils';

const CSS = __CLIENT__ ? require('components/layout/layout.less') : requireServerCss('components/layout/layout.less');

export default
class CounterButton extends React.Component {
  static propTypes = {
    location: PropTypes.object.isRequired
  }

  render() {
    const {location} = this.props;
    const path = location.pathname;

    return (
      <div>
        <div className={gcx('ui fixed pointing menu')} style={{borderLeft: 'none !important', borderRight: 'none !important', borderRadius: 0}}>
          <div className={gcx('ui container')}>
            <a href='/' className={gcx('header item')}>
              <span className={CSS.logo}/>
              Grox
            </a>
            <Link to='/' className={gcx('item', {[styles.active]: path === '/'})}>Front</Link>
            <Link to='/newest' className={gcx('item')} activeClassName={gcx('active')}>New</Link>
            <div className={gcx('right menu')}>
              <Link to='/submit' className={gcx('item')} activeClassName={gcx('active')}>
                <i className={gcx('fa fa-plus-circle')}/>
                &nbsp;Submit
              </Link>
              <a href='#' className={gcx('item')} tabIndex='0'>
                <span className={CSS.githubLogo}/>
                &nbsp;Login
              </a>
              <div className={gcx('ui floated simple dropdown item')} tabIndex='0'>
                <i className={gcx('fa fa-caret-down')}/>
                <div className={gcx('menu')} tabIndex='-1'>
                  <div className={gcx('ui compact basic icon buttons')} id={CSS.iconGroup}>
                    <button className={gcx('ui button active', CSS.iconFlex)}>
                      <i className={gcx('fa fa-list')}></i>
                    </button>
                    <button className={gcx('ui button', CSS.iconFlex)}>
                      <i className={gcx('fa fa-th-list')}></i>
                    </button>
                    <button className={gcx('ui button', CSS.iconFlex)}>
                      <i className={gcx('fa fa-th')}></i>
                    </button>
                    <button className={gcx('ui active button', CSS.iconFlex)}>
                      <i className={gcx('fa fa-thumb-tack')}></i>
                    </button>
                  </div>
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
                  <div className={gcx('divider')}></div>
                  <div className={gcx('header')}>Grox</div>
                  <div className={gcx('item')}>Send feedback!</div>
                  <a className={gcx('item')} href='https://github.com/golang-vietnam/grox'>Grox on GitHub</a>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div style={{height: '50px'}}/>
      </div>
    );
  }
}
