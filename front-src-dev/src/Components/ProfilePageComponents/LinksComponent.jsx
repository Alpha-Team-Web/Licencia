import React, {Component} from 'react';
import {Button} from "semantic-ui-react";
import {
    accountGithubChanged, closeAddRepoDiv, fill,
    firstRepoDiv, openAddRepoDiv,
    removeRepo,
    secondRepoDiv, submitGitPart,
    thirdRepoDiv
} from "../../Js Functionals/ProfilePage/JS1";

class LinksComponent extends Component {
    render() {
        return (
            <div className="ui form flexColumn formPadding" id={this.props.id} style={this.props.style}>
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
                               onfocusout={() => {
                                   accountGithubChanged();
                               }}/>
                    </div>
                </div>

                <div className="ui relaxed divided list" id="gitHubRepos">
                    <div className="item" id="firstRepo">
                        <i className="large github middle aligned icon"/>
                        <div className="content">
                            <a className="header" id="linkRepo1">Semantic-Org/Semantic-UI</a>
                        </div>
                        <i className="large window minimize middle aligned link icon"
                           onClick={() => {
                               fill();
                               removeRepo(firstRepoDiv);
                           }}/>
                    </div>
                    <div className="item" id="secondRepo">
                        <i className="large github middle aligned icon"/>
                        <div className="content">
                            <a className="header" id="linkRepo2">Semantic-Org/Semantic-UI-Docs</a>
                        </div>
                        <i className="large window minimize middle aligned link icon"
                           onClick={() => {
                               fill()
                               removeRepo(secondRepoDiv);
                           }}/>
                    </div>
                    <div className="item" id="thirdRepo">
                        <i className="large github middle aligned icon"/>
                        <div className="content">
                            <a className="header" id="linkRepo3">Semantic-Org/Semantic-UI-Meteor</a>
                        </div>
                        <i className="large window minimize middle aligned link icon"
                           onClick={() => {
                               fill()
                               removeRepo(thirdRepoDiv);
                           }}/>
                    </div>
                    <div className="item" id="addRepoDiv">
                        <div className="ui action input" id="addGitHubRepoInput">
                            <input type="text" placeholder="Search..." id="addRepoInput"
                                   onfocusout={() => {
                                       closeAddRepoDiv()
                                   }}/>
                            }
                            <Button className="ui icon button" onClick={() => {
                                closeAddRepoDiv();
                            }}>
                                <i className="plus circle icon"/>
                            </Button>
                        </div>
                    </div>
                    <div className="item" id="plusRepoIconDiv" onClick={() => {
                        openAddRepoDiv();
                    }}>
                        <i className="large plus circle middle aligned link icon central"/>
                    </div>
                </div>

                <button className="positive ui button rightAligned" id="saveLinksButton" onClick={() => {
                    submitGitPart();
                }}>ثبت
                    پیوند
                    ها
                </button>
            </div>
        );
    }
}

export default LinksComponent;
