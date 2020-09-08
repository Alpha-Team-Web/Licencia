import React, {Component} from 'react';
import {Divider} from "semantic-ui-react";
import ProfileComponent from "./profileComponent";
import LinksComponent from "./LinksComponent";
import ProfileCard from "./profileCard";
// import {changeMainProfileContent, gitHubRepoContent, profile} from "../../Js Functionals/ProfilePage/JS1";
import '../../CSS Designs/basic.css';
import '../../CSS Designs/ProfilePage/CSS1.css';
import '../../CSS Designs/Header.css';
import ModalPassword from "./passwordComponent";
import TransitionComponent from "./transitionComponent";
import {
    initTransitionsStart,
    loadProfileMenu,
    switchLinksToProfile,
    switchProfileToLinks
} from "../../Js Functionals/ProfilePage/JS1";
import ProfileForm from "./ProfileForm";

class ProfileContent extends Component {
    constructor(props, context) {
        super(props, context);
    }

    state = {
        profileDisplay: true,
        linksDisplay: false,
    }

    render() {
        return (
            <div className='content' id={this.props.id}>
                <div className='content transition visible flexRow' id='ProfileContents'>
                    <ProfileForm style={{display: 'block'}} id='profileComponent'/>

                    <LinksComponent style={{display: 'none'}} id='linksComponent'/>
                </div>

                <Divider id='divider1'/>

                <div className='content' id='profileCards'>

                    <ProfileCard onClick={switchLinksToProfile} hId='viewProfile' number={35}
                                 cardContent='مشاهده پروفایل'/>

                    <ProfileCard id='gitHubAccountPart' onClick={switchProfileToLinks}
                                 hId='viewLinks' number={35} cardContent='مشاهده پیوند ها'/>

                    <ModalPassword/>

                    <ProfileCard onClick={{/*() => changeMainProfileContent(gitHubRepoContent)*/}} hId='doneProjects'
                                 number={35} cardContent='پروژه های انجام شده در لیسنسیا'/>

                    <ProfileCard onClick={{/*() => changeMainProfileContent(gitHubRepoContent)*/}} hId='skills'
                                 number={12} cardContent='مهارت های ثبت شده در لیسنسیا'/>

                    <ProfileCard onClick={{/*() => changeMainProfileContent(gitHubRepoContent)*/}} hId='elTippo'
                                 number={95} cardContent='ای تیپو کسب شده'/>

                    <ProfileCard onClick={{/*() => changeMainProfileContent(gitHubRepoContent)*/}} hId='workInstance'
                                 number={95} cardContent='نمونه کارها'/>

                    <ProfileCard id='gitHubAccountPart'
                                 onClick={{/*() => changeMainProfileContent(gitHubRepoContent)*/}} hId='upgradeAccount'
                                 number={95} cardContent='اکانت خود را ارتقا دهید'/>

                </div>
            </div>
        );
    }
}

export default ProfileContent;
