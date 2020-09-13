import React, { Component } from 'react';
import '../../CSS Designs/Header.css';
import '../../CSS Designs/basic.css';
import '../../CSS Designs/extra-css.css'
import {logOut} from "../../Js Functionals/ProfilePage/profilePageContent";

class ProfileHeader extends Component {
    state = {  }
    render() {
        return (
            <div className="ui grid ui-rtl" style={{margin: '10px',}}>
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
                        <button className="ui labeled icon negative basic huge button" onClick={() => logOut()}>
                            <i className="arrow alternate circle left icon"></i>
                            خروج
                        </button>
                    </div>
                </div>

            </div>
        );
    }
}

export default ProfileHeader;
