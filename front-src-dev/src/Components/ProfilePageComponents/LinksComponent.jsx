import React, {Component} from 'react';
import {Button} from "semantic-ui-react";
import {
    addedRepoInputFocusOut, clickedPlusIcon, gitHubAccount, gitHubAccountChanged, submitGitPart,
} from "../../Js Functionals/ProfilePage/linksContent";
import '../../CSS Designs/ProfilePage/CSS1.css'
import Background from '../../Pics/githubImage2.png'
import GithubRepoComponent from "./GithubRepoComponent";
import {gitHubUrl} from "../../Js Functionals/urlNames";

class LinksComponent extends Component {
    constructor(props, context) {
        super(props, context);
    }

    render() {
        return (
            <div className="ui form flexColumn formPadding" style={{
                backgroundImage: "url(" + Background + ")",
                backgroundRepeat: 'no-repeat',
                marginLeft: '100px',
                display: 'none'
            }} id={this.props.id}>
                <div className="two fields">
                    <div className="field">
                        <label className="rightAligned marginBottom10">آدرس سایت</label>
                        <input maxLength="50" id="siteAddressField" placeholder="Site Address" type="text"/>
                    </div>
                </div>

                <div className="two fields">
                    <div className="field">
                        <label className="rightAligned marginBottom10">اکانت گیت هاب</label>
                        <input maxLength="40" id="githubAccountField" placeholder="Github-Account" type="text"
                               onBlur={() => gitHubAccountChanged()/*accountGithubChanged()*/}/>
                    </div>
                </div>

                <div className="ui relaxed" id="gitHubRepos">

                    <div className='ui relaxed divided list' id='gitHubRepositories' />

                    <div className="item" id="addRepoDiv">
                        <div className="ui action input" id="addGitHubRepoInput">
                            <input type="text" placeholder="Search..." id="addRepoInput"
                                   onBlur={() => addedRepoInputFocusOut()}/>
                            <Button className="ui icon button" onClick={() => addedRepoInputFocusOut()}>
                                <i className="plus circle icon"/>
                            </Button>
                        </div>
                    </div>
                    <div className="item" id="plusRepoIconDiv" onClick={() => clickedPlusIcon()}>
                        <i className="large plus circle middle aligned link icon central"/>
                    </div>
                </div>

                <button className="positive ui button rightAligned" id="saveLinksButton" onClick={() => submitGitPart()}>ثبت
                    پیوند
                    ها
                </button>
            </div>
        );
    }

}


export default LinksComponent;