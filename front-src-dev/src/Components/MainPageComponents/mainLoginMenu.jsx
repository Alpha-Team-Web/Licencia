import React, {Component} from 'react';
import '../../CSS Designs/MainPage/LoginMenu.css';
import MainInput from "./mainFormElements/mainInput";
import {signUpMenu} from "../../Js Functionals/MainPage/Login SignUp Show";
import {Label} from "semantic-ui-react";
import {login} from "../../Js Functionals/MainPage/IOMethods/loginMethods";
import '../../CSS Designs/extra-css.css'
import {emailMaxLengthInput, passwordMaxLengthInput} from "../../Js Functionals/MainPage/ioInputLengths";
import MainSelect from "./mainFormElements/mainSelect";
import MainInput2 from "./mainFormElements/mainInput2";
import {
    loginButtonId,
    loginKeyPointId,
    loginKindId,
    loginPasswordId,
    options
} from "../../Js Functionals/MainPage/loginSignupIds";

class MainLoginMenu extends Component {
    render() {

        return (
            <div id={this.props.id} style={this.props.style} className="content Login-SignUp-Menu">
                <div className="ui form formPadding ui-rtl">

                    <MainInput2 id={loginKeyPointId} maxLength={emailMaxLengthInput} textName='نام کاربری یا ایمیل'
                               placeHolder='Username Or Email' errorId="loginKeyPointError"/>

                    <MainInput2 id={loginPasswordId} maxLength={passwordMaxLengthInput} textName='رمز عبور'
                               placeHolder='Password' isPassword={true} errorID="loginPasswordError"/>

                    <MainSelect id={loginKindId} options={options} textName='نوع'/>

                    <div className="ui form formPadding" id="Login-Footer-Form">
                        <label>
                            <input onClick={() => {
                                login(this.props.onClose)
                            }} type="submit" id={loginButtonId} className="ui green button"
                                   value="login"/>
                        </label>
                        <footer className="loginFooter">
                            <p className="div-div-footer-p" id="LoginMenu-SignUp-Link">
                                <Label as='a' onClick={() => signUpMenu()} className="div-div-footer-signupLink">ثبت
                                    نام</Label>
                                آیا حساب ندارید؟
                            </p>
                        </footer>
                    </div>
                </div>
            </div>
        );
    }
}

export default MainLoginMenu;
