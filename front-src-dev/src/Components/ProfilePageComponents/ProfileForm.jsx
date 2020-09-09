import React, {Component} from 'react';
import '../../CSS Designs/basic.css';
import '../../CSS Designs/extra-css.css'
import {Divider, Grid, Image, Segment} from 'semantic-ui-react'
import profilePic from "../../Pics/codingPerson.jpg"
import {loadProfileMenu, saveProfile} from "../../Js Functionals/ProfilePage/JS1";

class ProfileForm extends Component {
    state = {}

    render() {
        return (
            <Segment style={this.props.style} id={this.props.id}>
                <Grid className="flexRow ui-rtl" relaxed='very'>
                    <Grid.Column width={8}>
                        <div className="ui form ui-rtl">
                            <div className="six wide field ui-rtl">
                                <label className='rightAligned no-space-break label-size'> نام کاربری</label>
                                <input readOnly className='rightAligned' type="text" placeholder="Username"
                                       id="usernameField"></input>
                            </div>
                            <div className="ten wide field ui-rtl">
                                <label>E-mail</label>
                                <input readOnly className="input-size" type="email" placeholder="Email"
                                       id="emailField"></input>
                            </div>
                            <div className="six wide field ui-rtl">
                                <label className='rightAligned no-space-break label-size'>نام نمایشی</label>
                                <input className="" type="text" placeholder="Showing Name"
                                       id="showingNameField"></input>
                            </div>
                            <div className="two fields ui-rtl">
                                <div className="six wide field ui-rtl">
                                    <label>نام</label>
                                    <input type="text" placeholder="First Name" id="firstNameField"></input>
                                </div>
                                <div className="six wide field ui-rtl">
                                    <label>نام خانوادگی</label>
                                    <input type="text" placeholder="Last Name" id="lastNameField"></input>
                                </div>
                            </div>
                            <div className="two fields ui-rtl">
                                <div className="six wide field">
                                    <label>آدرس</label>
                                    <input className="input-size" type="text" placeholder="Address"
                                           id="addressField"></input>
                                </div>
                                <div className="six wide field">
                                    <label className='rightAligned no-space-break label-size'>شماره تلفن</label>
                                    <input className="input-size" type="text" placeholder="Phone Number"
                                           id="telephoneNumberField"></input>
                                </div>
                            </div>
                            <div className=" ui-rtl">
                                <label>توضیحات</label>
                                <textarea rows="3" placeholder="Description" id="descriptionField"></textarea>
                            </div>
                            <button className="ui positive button" onClick={() => saveProfile()}>
                                ثبت پروفایل
                            </button>
                        </div>
                    </Grid.Column>
                    <Grid.Column width={5}>
                        <div className="flexColumn" id="leftDiv">
                            <div className="content" id="ImageTippoContent">
                                <img className="ui circular bordered image" src={profilePic}
                                     alt="*User*'s Picture"/>
                            </div>
                        </div>
                    </Grid.Column>
                </Grid>
            </Segment>
        );
    }
}

export default ProfileForm;
