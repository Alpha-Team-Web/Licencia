import React, {Component} from 'react';
import ProfileIndicesComponent from "./profileIndicesComponent";
import ProfileForm from "../ProfilePageComponents/ProfileSectionComponents/ProfileForm";

class SkillsPage extends Component {
    render() {
        return (
            <ProfileIndicesComponent>
                <ProfileForm/>
            </ProfileIndicesComponent>
        );
    }
}

export default SkillsPage;
