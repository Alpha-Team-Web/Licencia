import React, {Component} from 'react';
import '../../CSS Designs/MainPage/LoginMenu.css';
import MainInput from "./mainFormElements/mainInput";
import {signUpMenu} from "../../Js Functionals/MainPage/Login SignUp Show";
import {Label} from "semantic-ui-react";
import {login} from "../../Js Functionals/MainPage/IOMethods/loginMethods";
import '../../CSS Designs/extra-css.css'
import {emailMaxLengthInput, passwordMaxLengthInput} from "../../Js Functionals/MainPage/ioInputLengths";
import MainSelect from "./mainFormElements/mainSelect";

class MainLoginMenu extends Component {
    render() {
        return (
            <div id={this.props.id} style={this.props.style} className="content Login-SignUp-Menu" >
                <div className="ui form formPadding ui-rtl">

                    <MainInput id='login-KeyPoint' maxLength={emailMaxLengthInput} textName='نام کاربری یا ایمیل' placeHolder='Username Or Email' errorId="loginKeyPointError"/>

                    <MainInput id='login-Password' maxLength={passwordMaxLengthInput} textName='رمز عبور' placeHolder='Password' isPassword={true} errorID="loginPasswordError"/>

                    <MainSelect id='loginKind' value1 = 'employer' textValue1='کارفرما' value2 = 'freelancer' textValue2 = 'فریلنسر' textName='نوع'/>

                    <div className="ui form formPadding" id="Login-Footer-Form">
                        <label>
                            <input onClick={() => {
                                login(this.props.onClose)
                            }} type="submit" id="loginButton" className="ui green button"
                                   value="login" />
                        </label>
                        <footer className="loginFooter">
                            <p className="div-div-footer-p" id="LoginMenu-SignUp-Link">
                                <Label as='a' onClick={() => signUpMenu()} className="div-div-footer-signupLink" >ثبت نام</Label>
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
