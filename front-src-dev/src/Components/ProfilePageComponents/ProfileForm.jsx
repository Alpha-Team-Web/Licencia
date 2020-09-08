import React, { Component } from 'react';
import '../../CSS Designs/basic.css';
import '../../CSS Designs/extra-css.css'

class ProfileForm extends Component {
    state = {  }
    render() {
        return (
            <div className="ui form">
                <div className="inline fields">
                    <div className="six wide field">
                        <input className='rightAligned' type="text" placeholder="Username" id="username"></input>
                        <label className='rightAligned no-space-break label-size'> نام کاربری</label>
                    </div>
                    <div className="six wide field">
                        <input type="text" placeholder="Showing Name" id="nickName"></input>
                        <label className='rightAligned no-space-break label-size'>نام نمایشی</label>
                    </div>
                </div>
                <div className="inline fields">
                    <div className="five wide field">
                        <input type="text" placeholder="Last Name" id="lastName"></input>
                    </div>
                    <div className="eight wide field">
                        <input type="text" placeholder="First Name" id="firstName"></input>
                        <label>نام</label>
                    </div>
                </div>
                <div class="field inline">
                    <input type="email" placeholder="Email" id="email"></input>
                    <label>E-mail</label>
                </div>
                <div className="inline fields">
                    <div className="six wide field">
                        <input type="text" placeholder="Address" id="address"></input>
                        <label>آدرس</label>
                    </div>
                    <div className="six wide field">
                        <input type="text" placeholder="Phone Number" id="phoneNumber"></input>
                        <label className='rightAligned no-space-break label-size'>شماره تلفن</label>
                    </div>
                </div>
                <div class="fields inline">
                    <textarea rows="3" placeholder="Description" id="description"></textarea>
                    <label>توضیحات</label>
                </div>
                <button className="ui button" onClick={() => }>
                    confirm
                </button>
            </div>

        );
    }
}

export default ProfileForm;