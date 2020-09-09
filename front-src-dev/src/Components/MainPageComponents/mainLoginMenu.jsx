import React, {Component} from 'react';
import '../../CSS Designs/MainPage/LoginMenu.css';
import MainTextField from "./mainTextField";
import {login} from "../../Js Functionals/MainPage/Login SignUp Method";
import {signUpMenu} from "../../Js Functionals/MainPage/Login SignUp Show";
import {Label} from "semantic-ui-react";

class MainLoginMenu extends Component {
    render() {
        return (
            <div id={this.props.id} style={this.props.style} className="content Login-SignUp-Menu" >
                <div className="ui form formPadding">
                    {/*<div className="ui form formPadding">
                        <div className="ui field">
                            <p className="paragraphInput">نام کاربری یا ایمیل</p>
                            <input maxLength="50" type="text" placeholder="Username Or Email"
                                   id="login-KeyPoint" onFocus="setFieldError(this, false)" />
                        </div>
                    </div>*/}

                    {/*<div className="ui form formPadding">
                        <div className="ui field">
                            <p className="paragraphInput">رمز عبور</p>
                            <input maxLength="30" type="password" className="div-div-div-input" id="login-Password"
                                   placeholder="Password" onFocus="setFieldError(this, false)" />
                        </div>
                    </div>*/}

                    <MainTextField id='login-KeyPoint' maxLength='50' textName='نام کاربری یا ایمیل' placeHolder='Username Or Email' />

                    <MainTextField id='login-Password' maxLength='30' textName='رمز عبور' placeHolder='Password' isPassword={true} />

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
                            <input onClick={() => login()} type="submit" id="loginButton" className="ui green button"
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
