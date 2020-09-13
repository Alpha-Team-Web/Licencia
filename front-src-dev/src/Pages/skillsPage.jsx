import React, {Component} from 'react';
import ProfileIndicesComponent from "../Components/SkillPageComponents/profileIndicesComponent";
import ProfileForm from "../Components/ProfilePageComponents/ProfileSectionComponents/ProfileForm";
import SkillsPageComponent from "../Components/SkillPageComponents/skillsPageComponent";

class SkillsPage extends Component {
    render() {
        return (
            <ProfileIndicesComponent>
                <SkillsPageComponent/>
            </ProfileIndicesComponent>
        );
    }
}

export default SkillsPage;
