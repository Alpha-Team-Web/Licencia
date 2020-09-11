import React, {Fragment, Component} from 'react';
import MainTextField from "./mainTextField";
import {signUp} from "../../Js Functionals/MainPage/IOMethods/signUpMethods";
import '../../CSS Designs/extra-css.css'
import {emailMaxLength, passwordMaxLength} from "../../Js Functionals/MainPage/Login SignUp Show";
import {
    emailMaxLengthInput, firstnameMaxLengthInput,
    lastnameMaxLengthInput,
    passwordMaxLengthInput, usernameMaxLengthInput
} from "../../Js Functionals/MainPage/ioInputLengths";

class MainSignUpMenu extends Component {
    constructor(props, context) {
        super(props, context);
    }

    render() {
        return (
            <div id={this.props.id} style={this.props.style} className="content Login-SignUp-Menu">
                <div className="ui form formPadding ui-rtl">

                    <MainTextField id='SignUp-UserName' maxLength={usernameMaxLengthInput} textName='نام کاربری' placeHolder='Username'/>

                    <MainTextField id="SignUp-FirstName" maxLength={firstnameMaxLengthInput} textName='نام' placeHolder='FirstName' />

                    <MainTextField id="SignUp-LastName" maxLength={lastnameMaxLengthInput} textName='نام خانوادگی' placeHolder='LastName' />

                    <MainTextField id="SignUp-Email" maxLength={emailMaxLengthInput} textName='ایمیل' placeHolder='Email address'/>

                    <MainTextField id="SignUp-Password" maxLength={passwordMaxLengthInput} textName='رمز عبور' placeHolder='Password'
                                   isPassword={true}/>

                    <MainTextField id="SignUp-RepeatPassword" maxLength={passwordMaxLengthInput} textName='تکرار رمز عبور'
                                   placeHolder='Repeat Password' isPassword={true}/>

                    <div className="ui form formPadding">
                        <div className="ui paragraph paragraphInput">نوع</div>
                        <label>
                            <select className="ui dropdown" id="signUpKind">
                                <option value="employer">کارفرما</option>
                                <option value="freelancer">فریلنسر</option>
                            </select>
                        </label>
                    </div>

                    <div className="ui form formPadding">
                        <label>
                            <input onClick={() => {
                                signUp(this.props.onClose)
                            }} type="submit" id="signUpButton" className="ui green button"
                                   value="ثبت نام"/>
                        </label>
                    </div>
                </div>
            </div>
        );
    }
}

export default MainSignUpMenu;
