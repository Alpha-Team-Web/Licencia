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
                        <input className='rightAligned' type="text" placeholder="Username"></input>
                        <label className='rightAligned no-space-break label-size'> نام کاربری</label>
                    </div>
                    <div className="six wide field">
                        <input type="text" placeholder="Showing Name"></input>
                        <label className='rightAligned no-space-break label-size'>نام نمایشی</label>
                    </div>
                </div>
                <div className="inline fields">
                    <div className="five wide field">
                        <input type="text" placeholder="Last Name"></input>
                    </div>
                    <div className="eight wide field">
                        <input type="text" placeholder="First Name"></input>
                        <label>نام</label>
                    </div>
                </div>
                <div class="field inline">
                    <input type="email" placeholder="Email"></input>
                    <label>E-mail</label>
                </div>
                <div className="inline fields">
                    <div className="six wide field">
                        <input type="text" placeholder="Address"></input>
                        <label>آدرس</label>
                    </div>
                    <div className="six wide field">
                        <input type="text" placeholder="Phone Number"></input>
                        <label className='rightAligned no-space-break label-size'>شماره تلفن</label>
                    </div>
                </div>
                <div class="fields inline">
                    <textarea rows="3" placeholder="Description"></textarea>
                    <label>توضیحات</label>
                </div>
            </div>

        );
    }
}

export default ProfileForm;