import React, { Component } from 'react';
import '../../CSS Designs/Header.css';
import '../../CSS Designs/basic.css';
import '../../CSS Designs/extra-css.css'
import GridColumn, {Grid} from "semantic-ui-react";

class Menu extends Component {
    state = {  }
    render() {
        return (
        <div className="ui grid ui-rtl">
            <div className="right floated five wide column">
                <div className="ui borderless menu menu-item">
                    <a className="item ">دنبال کننده ها</a>
                    <a className="item">دنبال شونده ها</a>
                    <a className="item">مشاهده</a>
                    <a className="item">دسته بندی ها</a>
                </div>
            </div>
            <div className="left floated five wide column">
                <div className="item">
                    <button className="ui labeled icon negative basic huge button rightAligned">
                        <i className="arrow alternate circle left icon"></i>
                        خروج
                    </button>
                </div>
            </div>

        </div>

        );
    }
}

export default Menu;