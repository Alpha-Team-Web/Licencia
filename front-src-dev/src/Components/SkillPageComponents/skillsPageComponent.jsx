import React, {Component} from 'react';
import PersonSkillsComponent from "./personSkillsComponent";
import ProfileForm from "../ProfilePageComponents/ProfileSectionComponents/ProfileForm";
import {Input} from "semantic-ui-react";

class SkillsPageComponent extends Component {
    mainDivStyle = {

    };

    firstRealStyle = {

    };

    secondRealStyle = {

    };


    render() {
        return (
            <div style={this.mainDivStyle}>
                <div style={this.firstRealStyle}>
                    <ProfileForm/>
                </div>

                <div style={this.secondRealStyle}>
                    <Input icon='search' className='ui-rtl'/>
                    <ProfileForm/>
                </div>
            </div>
        );
    }
}

export default SkillsPageComponent;
