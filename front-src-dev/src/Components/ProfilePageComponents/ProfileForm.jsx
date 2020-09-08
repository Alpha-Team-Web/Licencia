import React, { Component } from 'react';
import '../../CSS Designs/basic.css';
import '../../CSS Designs/extra-css.css'
import { Divider, Grid, Image, Segment } from 'semantic-ui-react'
import profilePic from "../../Pics/codingPerson.jpg"

class ProfileForm extends Component {
    state = {  }
    render() {
        return (
            <Segment style={this.props.style} id={this.props.id}>
                <Grid className="flexRow" relaxed='very'>
                    <Grid.Column width={5}>
                        <div className="flexColumn" id="leftDiv">
                            <div className="content" id="ImageTippoContent">
                                <img className="ui circular bordered image" src={profilePic}
                                     alt="*User*'s Picture"/>
                            </div>
                        </div>
                    </Grid.Column>
                    <Grid.Column width={8} >
                        <div className="ui form ">
                            <div className="two inline fields ">
                                <div className="six wide field">
                                    <input className='rightAligned input-size' type="text" placeholder="Username" id="username"></input>
                                    <label className='rightAligned no-space-break label-size'> نام کاربری</label>
                                </div>
                                <div className="six wide field">
                                    <input className="input-size" type="text" placeholder="Showing Name" id="nickName"></input>
                                    <label className='rightAligned no-space-break label-size'>نام نمایشی</label>
                                </div>
                            </div>
                            <div className="two inline fields">
                                <div className="six wide field">
                                    <input className="input-size" type="text" placeholder="Last Name" id="lastName"></input>
                                </div>
                                <div className="six wide field">
                                    <input className="input-size" type="text" placeholder="First Name" id="firstName"></input>
                                    <label>نام</label>
                                </div>
                            </div>
                            <div className="twenty wide field inline">
                                <input className="input-size" type="email" placeholder="Email" id="email"></input>
                                <label>E-mail</label>
                            </div>
                            <div className="two inline fields">
                                <div className="six wide field">
                                    <input className="input-size" type="text" placeholder="Address" id="address"></input>
                                    <label>آدرس</label>
                                </div>
                                <div className="six wide field">
                                    <input  className="input-size" type="text" placeholder="Phone Number" id="phoneNumber"></input>
                                    <label className='rightAligned no-space-break label-size'>شماره تلفن</label>
                                </div>
                            </div>
                            <div className="fields inline">
                                <textarea rows="3" placeholder="Description" id="description"></textarea>
                                <label>توضیحات</label>
                            </div>
                            <button className="ui button" >
                                confirm
                            </button>
                        </div>
                    </Grid.Column>
                </Grid>
            </Segment>
        );
    }
}

export default ProfileForm;
