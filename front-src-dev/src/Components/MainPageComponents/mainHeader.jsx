import React, {Component} from 'react';
import '../../CSS Designs/Header.css';
import '../../CSS Designs/basic.css';
import {Image} from "semantic-ui-react";
import HeaderListItem from "./HeaderListItem";
import LicenciaLogo from '../../Pics/Licencia-Logo.png';
import ModalLogSin from "./LogSinMenuModal";

class MainHeader extends Component {
    render() {
        return (
            <div className="header" id="Header">
                <ModalLogSin/>
                <ul id='Header-UnOrderedList'>
                    <HeaderListItem href='KireKhar' value='دسته بندی ها'/>
                    <HeaderListItem href='KireKhar' value='مشاهده'/>
                    <HeaderListItem href='KireKhar' value='دنبال کننده ها'/>
                    <HeaderListItem href='KireKhar' value='دنبال شونده ها'/>
                </ul>
                <Image as='a' href="#badan" className='Header-Link' src={LicenciaLogo} alt="Licencia" id="Logo"/>
            </div>
        );
    }
}

export default MainHeader;
