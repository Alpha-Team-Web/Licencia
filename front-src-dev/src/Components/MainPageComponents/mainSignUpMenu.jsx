import React, {Fragment, Component} from 'react';
import MainInput from "./mainFormElements/mainInput";
import {signUp} from "../../Js Functionals/MainPage/IOMethods/signUpMethods";
import '../../CSS Designs/extra-css.css'
import MainInput2 from "./mainFormElements/mainInput2";
import MainSelect from "./mainFormElements/mainSelect";
import {
    signUpButtonId,
    signUpEmailId, signUpFirstNameId,
    signUpKindId, signUpLastNameId,
    signUpPasswordId,
    signUpRepeatPasswordId, signUpUserNameId
} from "../../Js Functionals/MainPage/loginSignupIds";

class MainSignUpMenu extends Component {
    constructor(props, context) {
        super(props, context);
    }

    render() {
        return (
            <div id={this.props.id} style={this.props.style} className="content Login-SignUp-Menu">
                <div className="ui form formPadding ui-rtl">

                    <MainInput2 id={signUpUserNameId} maxLength='30' textName='نام کاربری' placeHolder='Username'/>

                    <MainInput2 id={signUpFirstNameId} maxLength='30' textName='نام' placeHolder='FirstName'/>

                    <MainInput2 id={signUpLastNameId} maxLength='30' textName='نام خانوادگی' placeHolder='LastName'/>

                    <MainInput2 id={signUpEmailId} maxLength='50' textName='ایمیل' placeHolder='Email address'/>

                    <MainInput2 id={signUpPasswordId} maxLength='30' textName='رمز عبور' placeHolder='Password'
                                isPassword={true}/>

                    <MainInput2 id={signUpRepeatPasswordId} maxLength='30' textName='تکرار رمز عبور'
                                placeHolder='Repeat Password' isPassword={true}/>

                    <MainSelect id={signUpKindId} options={options} textName='نوع'/>


                    <div className="ui form formPadding">
                        <label>
                            <input onClick={() => {
                                signUp(this.props.onClose)
                            }} type="submit" id={signUpButtonId} className="ui green button"
                                   value="ثبت نام"/>
                        </label>
                    </div>
                </div>
            </div>
        );
    }
}

export default MainSignUpMenu;
