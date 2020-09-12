import React, {Component} from 'react';
import '../../../CSS Designs/basic.css';
import '../../../CSS Designs/extra-css.css'
import {Grid} from 'semantic-ui-react'
import profilePic from "../../../Pics/codingPerson.jpg"
import {saveProfile} from "../../../Js Functionals/ProfilePage/profileContent";
import ProfilePictureComponent from "../Utils/profilePictureComponent";
import {
    addressMaxLengthInput, descriptionMaxLengthInput,
    firstnameMaxLengthInput,
    lastnameMaxLengthInput, phoneNumberMaxLengthInput,
    shownNameMaxLengthInput
} from "../../../Js Functionals/MainPage/ioInputLengths";

class ProfileForm extends Component {
    state = {}

    render() {
        return (
            <Grid centered className="flexRow ui-rtl" relaxed='very' style={this.props.style} id={this.props.id}>
                <Grid.Row stretched >
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
                                <input maxLength={shownNameMaxLengthInput} className="" type="text" placeholder="Showing Name"
                                       id="showingNameField"></input>
                            </div>
                            <div className="two fields ui-rtl">
                                <div className="six wide field ui-rtl">
                                    <label>نام</label>
                                    <input maxLength={firstnameMaxLengthInput} type="text" placeholder="First Name" id="firstNameField"></input>
                                </div>
                                <div className="six wide field ui-rtl">
                                    <label>نام خانوادگی</label>
                                    <input maxLength={lastnameMaxLengthInput} type="text" placeholder="Last Name" id="lastNameField"></input>
                                </div>
                            </div>
                            <div className="two fields ui-rtl">
                                <div className="six wide field">
                                    <label>آدرس</label>
                                    <input maxLength={addressMaxLengthInput} className="input-size" type="text" placeholder="Address"
                                           id="addressField"></input>
                                </div>
                                <div className="six wide field">
                                    <label className='rightAligned no-space-break label-size'>شماره تلفن</label>
                                    <input maxLength={phoneNumberMaxLengthInput} className="input-size" type="text" placeholder="Phone Number"
                                           id="telephoneNumberField"></input>
                                </div>
                            </div>
                            <div className=" ui-rtl">
                                <label>توضیحات</label>
                                <textarea maxLength={descriptionMaxLengthInput} rows="3" placeholder="Description" id="descriptionField"></textarea>
                            </div>
                            <button className="ui positive button" onClick={() => saveProfile()}>
                                ثبت پروفایل
                            </button>
                        </div>
                    </Grid.Column>
                    <Grid.Column centered width={6} verticalAlign={"middle"}>
                        <div className="picture-align" id="leftDiv">
                            {/*<img className="ui circular bordered image" src={profilePic}
                                 alt="*User*'s Picture"/>*/}
                            <ProfilePictureComponent className="ui circular bordered image" contentStyle="picture-align" src={profilePic}
                                                     alt="*User*'s Picture" imageId='profilePicture'
                                                     addImageInputId='addPictureInput' id="ImageTippoContent"/>
                        </div>
                    </Grid.Column>
                </Grid.Row>
            </Grid>
        );
    }
}

export default ProfileForm;
