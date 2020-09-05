import React, {Component, useState} from 'react'
import { Button, Modal } from 'semantic-ui-react';
import MainLoginMenu from "./mainLoginMenu";
import MainSignUpMenu from "./mainSignUpMenu";
import licenciaImg from '../../Pics/Licencia-Logo.png';
import {loginMenu, signUpMenu} from "../../Js Functionals/MainPage/Login SignUp Show";

export default class LogSinMenu extends Component {
    render() {
        let open, setOpen;
        [open, setOpen] = useState(false);

        return (
            <Modal
                onClose={() => setOpen(false)}
                onOpen={() => {
                    setOpen(true);
                    loginMenu();
                }}
                open={open}
                trigger={<Button className="loginButton" onClick="openLogSinMenu()">ورود / ثبت نام</Button>}
            >
                <Modal.Content>
                    <div className="header" id="Login-Menu-Header">
                        <div id="Signup-Login">
                            <div className="Signup-login-text" id="LoginMenuButton" onClick={() => loginMenu()}>ورود</div>
                            <div className="Signup-login-text" id="SignUpMenuButton" onClick={() => signUpMenu()}>ثبت نام</div>
                        </div>
                        <div className="image content">
                            <img src={licenciaImg} id="logoImage" alt="logoLicencia"/>
                        </div>
                        <h3 id="welcomeHeader">Welcome To Licencia</h3>
                    </div>
                    <MainLoginMenu/>
                    <MainSignUpMenu/>
                </Modal.Content>
            </Modal>
        )
    }
}
