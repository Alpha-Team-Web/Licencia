import React, {Component} from 'react';
import {fillRepoContentFields, removeRepo} from "../../Js Functionals/ProfilePage/linksContent";
import {
    fillRepoContentFields,
    removeRepo,
} from "../../Js Functionals/ProfilePage/linksContent";
import '../../CSS Designs/ProfilePage/CSS1.css'

class GithubRepoComponent extends Component {
    constructor(props, context) {
        super(props, context);
    }

    render() {
        return (
            <div className="item" id={this.props.id}>
                <i className="large github middle aligned icon"/>
                <div className="content">
                    <a className="header" id={this.props.link}>Semantic-Org/Semantic-UI</a>
                </div>
                <i className="large window minimize middle aligned link icon"
                   onClick={() => {
                       fillRepoContentFields();
                       removeRepo(this.props.repo);
                   }}/>
            </div>
        )
    }
}

export default GithubRepoComponent;