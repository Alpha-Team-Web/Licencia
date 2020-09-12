import React, {Component} from 'react';
import PersonSkillsSideBar from "../../Js Functionals/SkillsPage/personSkillsSideBar";
import ProfileHeader from "../ProfilePageComponents/profileHeaderComponent";
import {Button} from "semantic-ui-react";

class ProfileIndicesComponent extends Component {

    state = {
        visible: false,
    }

    setVisible = (visible) => {
        alert('fuck')
        alert('previous Visible: ' + this.getVisible())
        alert('visible: ' + visible)
        this.setState({visible: visible}, () => alert('visible Again: ' + this.getVisible()))
    }

    getVisible = () => this.state.visible;

    render() {
        return (
            <div>
                <div className='flexRow'>
                    <ProfileHeader/>
                    <Button onClick={() => this.setVisible(!this.getVisible())}/>
                </div>
                <div>
                    <PersonSkillsSideBar visible={this.getVisible()} setVisible={this.setVisible}>
                        {this.props.children}
                    </PersonSkillsSideBar>
                </div>
            </div>
        );
    }
}

export default ProfileIndicesComponent;
