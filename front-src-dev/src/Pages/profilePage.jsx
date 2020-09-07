import React, {Component, Fragment} from 'react';
import ProfileHeader from "../Components/ProfilePageComponents/profileHeader";
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
