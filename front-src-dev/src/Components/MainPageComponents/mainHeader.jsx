import React, {Component} from 'react';
import '../../CSS Designs/Header.css';
import '../../CSS Designs/basic.css';
import {Button, Image} from "semantic-ui-react";
import HeaderListItem from "./HeaderListItem";
// import '../../CSS Designs/MainPage';

class MainHeader extends Component {
    render() {
        return (
            <div className="header" id="Header">
                <Button className="loginButton" onClick="openLogSinMenu()">ورود / ثبت نام</Button>
                <ul id='Header-UnOrderedList'>
                    <HeaderListItem href='KireKhar' value='دسته بندی ها'/>
                    <HeaderListItem href='KireKhar' value='مشاهده'/>
                    <HeaderListItem href='KireKhar' value='دنبال کننده ها'/>
                    <HeaderListItem href='KireKhar' value='دنبال شونده ها'/>
                </ul>
                {/*<a href="#badan" className="Header-Link"><img src="../../Pics/Licencia-Logo.png" alt="Licencia" id="Logo"/></a>*/}
                <Image as='a' href="#badan" className='Header-Link' src='../../Pics/Licencia-Logo.png' alt="Licencia" id="Logo"/>
            </div>
        );
    }
}

export default MainHeader;
