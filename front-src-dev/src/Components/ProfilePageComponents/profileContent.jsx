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
import {initTransitionsStart, loadProfileMenu} from "../../Js Functionals/ProfilePage/JS1";

class ProfileContent extends Component {
    constructor(props, context) {
        super(props, context);

    }

    state = {
        profileDisplay: true,
        linksDisplay: false,
    }

    showProfile = () => {
        if (!this.state.profileDisplay) {
            this.setState({profileDisplay: true, linksDisplay: false})
        }
    }

    showLinks = () => {
        if (!this.state.linksDisplay) {
            this.setState({profileDisplay: false, linksDisplay: true})
        }
    }


    initProfileContent = () => {
        initTransitionsStart(this.profileComponent, this.linksComponent);
        loadProfileMenu();
    }

    profileComponent = <TransitionComponent animation='scale' content={<p>jfofsie</p>}
                                            visibility={this.state.profileDisplay}/>

    linksComponent = <TransitionComponent
        animation='scale'
        content={<LinksComponent id='gitHubReposContent'/>} visibility={this.state.linksDisplay}/>

    render() {
        return (
            <div className='content' id={this.props.id}>
                <div className='content transition visible flexRow' id='ProfileContents'>
                    {this.profileComponent}

                    {this.linksComponent}
                </div>

                <Divider id='divider1'/>

                <div className='content' id='profileCards'>

                    <ProfileCard onClick={this.showProfile} hId='viewProfile' number={35}
                                 cardContent='مشاهده پروفایل'/>

                    <ProfileCard id='gitHubAccountPart' onClick={this.showLinks}
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
