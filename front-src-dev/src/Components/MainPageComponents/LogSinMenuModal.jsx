import React, {Component, useState} from 'react'
import {Button, Modal} from 'semantic-ui-react';
import MainLoginMenu from "./mainLoginMenu";
import MainSignUpMenu from "./mainSignUpMenu";
import licenciaImg from '../../Pics/Licencia-Logo.png';
import '../../CSS Designs/MainPage/LoginMenu.css';
import '../../CSS Designs/MainPage/loginSignupInput.css'
import {greenColor, loginMenu, signUpMenu} from "../../Js Functionals/MainPage/Login SignUp Show";


export default class LogSinMenu extends Component {
    render() {
        let open, setOpen;
        [open, setOpen] = useState(false);
        let style1 = {
            display: "none"
        }
        let style2 = {
            backgroundColor: greenColor
        }

        return (
            <Modal
                onClose={() => setOpen(false)}
                onOpen={() => {
                    setOpen(true);
                    loginMenu();
                }}
                open={open}
                // trigger={<Button className="loginButton" onClick="openLogSinMenu()">ورود / ثبت نام</Button>}
                trigger={<Button>Show Modal</Button>}
            >
                <Modal.Content>
                    <div className="header" id="Login-Menu-Header">
                        <div id="Signup-Login">
                            <div style={style2} className="Signup-login-text" id="LoginMenuButton"
                                 onClick={() => loginMenu()}>ورود
                            </div>
                            <div className="Signup-login-text" id="SignUpMenuButton" onClick={() => signUpMenu()}>ثبت
                                نام
                            </div>
                        </div>
                        <div className="image content">
                            <img src={licenciaImg} id="logoImage" alt="logoLicencia"/>
                        </div>
                        <h3 id="welcomeHeader">Welcome To Licencia</h3>
                    </div>
                    <MainLoginMenu  id="Login-Menu"/>
                    <MainSignUpMenu style = {style1}  id="SignUp-Menu"/>
                </Modal.Content>
            </Modal>
        )
    }
}
