import React, {Component, Fragment} from 'react';
import ProfileHeader from "../Components/ProfilePageComponents/profileHeaderComponent";
import ProfileContent from "../Components/ProfilePageComponents/profileContent";

class ProfilePage extends Component {
    render() {
        return (
            <Fragment>
                <ProfileHeader/>
                <ProfileContent/>
            </Fragment>
        );
    }
}

export default ProfilePage;
