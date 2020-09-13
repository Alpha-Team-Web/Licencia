import React, {Component} from 'react';
import ProfileIndicesComponent from "../Components/SkillPageComponents/profileIndicesComponent";
import ProfileForm from "../Components/ProfilePageComponents/ProfileSectionComponents/ProfileForm";

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
