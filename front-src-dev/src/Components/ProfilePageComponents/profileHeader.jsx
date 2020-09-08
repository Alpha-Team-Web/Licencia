import React, { Component } from 'react';
import '../../CSS Designs/Header.css';
import '../../CSS Designs/basic.css';
import '../../CSS Designs/extra-css.css'

class Menu extends Component {
    state = {  }
    render() {
        return (
            <div className=" header">
                <div className="ui teal four borderless menu">
                    <div className="item">
                        <button className="ui labeled icon negative basic huge button rightAligned">
                            <i class="arrow alternate circle left icon"></i>
                            خروج
                        </button>
                    </div>
                    <div className="ui borderless menu menu-item">
                        <a className="item ">دنبال کننده ها</a>
                        <a className="item">دنبال شونده ها</a>
                        <a className="item">مشاهده</a>
                        <a className="item">دسته بندی ها</a>
                    </div>
                </div>
            </div>
        );
    }
}1

export default Menu;