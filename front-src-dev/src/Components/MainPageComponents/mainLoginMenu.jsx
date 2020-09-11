import React, {Component} from 'react';
import '../../CSS Designs/MainPage/LoginMenu.css';
import MainTextField from "./mainTextField";
import {signUpMenu} from "../../Js Functionals/MainPage/Login SignUp Show";
import {Label} from "semantic-ui-react";
import {login} from "../../Js Functionals/MainPage/IOMethods/loginMethods";
import '../../CSS Designs/extra-css.css'

class MainLoginMenu extends Component {
    render() {
        return (
            <div id={this.props.id} style={this.props.style} className="content Login-SignUp-Menu" >
                <div className="ui form formPadding ui-rtl">

                    <MainTextField id='login-KeyPoint' maxLength='50' textName='نام کاربری یا ایمیل' placeHolder='Username Or Email' errorId="loginKeyPointError"/>

                    <MainTextField id='login-Password' maxLength='30' textName='رمز عبور' placeHolder='Password' isPassword={true} errorID="loginPasswordError"/>

                    <div className="ui form formPadding">
                        <div className="ui field">
                            <p className="paragraphInput">نوع</p>
                            <select className="ui dropdown" id="loginKind">
                                <option value="employer">کارفرما</option>
                                <option value="freelancer">فریلنسر</option>
                            </select>
                        </div>
                    </div>

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
