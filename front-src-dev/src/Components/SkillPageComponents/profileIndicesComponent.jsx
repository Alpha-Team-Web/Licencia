import React, {Component} from 'react';
import ProfileSideBar from "../../Js Functionals/SkillsPage/profileSideBar";
import ProfileHeader from "../ProfilePageComponents/profileHeaderComponent";
import {Button} from "semantic-ui-react";
import SkillComponent from "./MinorComponents/skillComponent";

class ProfileIndicesComponent extends Component {

    state = {
        visible: false,
    }

    setVisible = (visible) => this.setState({visible: visible})
    getVisible = () => this.state.visible;

    render() {
        return (
            <div>
                <div className='flexRow'>
                    <ProfileHeader/>
                    <Button icon='list alternate' onClick={() => this.setVisible(!this.getVisible())}/>
                </div>
                <div>
                    <ProfileSideBar visible={this.getVisible()} setVisible={this.setVisible}>
                        {this.props.children}
                    </ProfileSideBar>
                </div>
                <div>
                    <SkillComponent contentText="motarjemi">

                    </SkillComponent>
                </div>
            </div>
        );
    }
}

export default ProfileIndicesComponent;
