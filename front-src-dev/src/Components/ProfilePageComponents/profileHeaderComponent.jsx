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
                <div className="left floated column">
                    <div className="item ltr" id="exitButton">
                        <button className="ui negative button" onClick={() => logOut()}>
                            خروج
                        </button>
                    </div>
                </div>

            </div>
        );
    }
}

export default ProfileHeader;
