import React, {Fragment, Component} from 'react';
import MainTextField from "./mainTextField";

class MainSignUpMenu extends Component {
    constructor(props, context) {
        super(props, context);
    }

    render() {
        return (
            <div id={this.props.id} style = {this.props.style} className="content Login-SignUp-Menu" >
                <div className="ui form formPadding">
                    {/*<div className="ui form formPadding">
                        <div className="ui field">
                            <p className="paragraphInput">نام کاربری</p>
                            <input maxLength="30" type="text" className="div-div-div-input" id="SignUp-UserName"
                                   placeholder="username" onFocus="setFieldError(this, false)"/>
                        </div>
                    </div>*/}
                    <MainTextField id='SignUp-UserName' maxLength='30' textName='نام کاربری' placeHolder='Username' />

                   {/* <div className="ui form formPadding">
                        <div className="ui field">
                            <p className="paragraphInput">نام</p>
                            <input maxLength="30" type="text" className="div-div-div-input" id="SignUp-FirstName"
                                   placeholder="FirstName" onFocus="setFieldError(this, false)"/>
                        </div>
                    </div>*/}
                    <MainTextField id="SignUp-FirstName" maxLength='30' textName='نام' placeHolder='FirstName' />

                    {/*<div className="ui form formPadding">
                        <div className="ui field">
                            <p className="paragraphInput">نام خانوادگی</p>
                            <input maxLength="30" type="text" className="div-div-div-input" id="SignUp-LastName"
                                   placeholder="LastName" onFocus="setFieldError(this, false)"/>
                        </div>
                    </div>*/}
                    <MainTextField id="SignUp-LastName" maxLength='30' textName='نام خانوادگی' placeHolder='LastName' />

                    {/*<div className="ui form formPadding">
                        <div className="ui field">
                            <p className="paragraphInput">ایمیل</p>
                            <input maxLength="50" type="text" className="div-div-div-input" id="SignUp-Email"
                                   placeholder="Email Address" onFocus="setFieldError(this, false)"/>
                        </div>
                    </div>*/}
                    <MainTextField id="SignUp-Email" maxLength='50' textName='ایمیل' placeHolder='Email address' />

                    {/*<div className="ui form formPadding">
                        <div className="ui field">
                            <p className="paragraphInput">رمز عبور</p>
                            <input maxLength="30" type="password" className="div-div-div-input" id="SignUp-Password"
                                   placeholder="Password" onFocus="setFieldError(this, false)"/>
                        </div>
                    </div>*/}
                    <MainTextField id="SignUp-Password" maxLength='30' textName='رمز عبور' placeHolder='Password' />

                    {/*<div className="ui form formPadding">
                        <div className="ui field">
                            <p className="paragraphInput">تکرار رمز عبور</p>
                            <input maxLength="30" type="password" className="div-div-div-input"
                                   id="SignUp-RepeatPassword"
                                   placeholder="Repeat Password" onFocus="setFieldError(this, false)"/>
                        </div>
                    </div>*/}
                    <MainTextField id="SignUp-RepeatPassword" maxLength='30' textName='تکرار رمز عبور' placeHolder='Repeat Password' />


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
                            <input /*onClick="signUp()"*/ type="submit" id="signUpButton" className="ui green button"
                                   value="ثبت نام"/>
                        </label>
                    </div>
                </div>
            </div>
        );
    }
}

export default MainSignUpMenu;
