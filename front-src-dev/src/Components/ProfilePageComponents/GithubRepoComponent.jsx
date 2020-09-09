import React, {Component} from 'react';
import {removeRepository} from "../../Js Functionals/ProfilePage/linksContent";
import '../../CSS Designs/ProfilePage/CSS1.css'
import '../../CSS Designs/Header.css';
import '../../CSS Designs/basic.css';

class GithubRepoComponent extends Component {
    constructor(props, context) {
        super(props, context);
    }

    render() {
        return (
            <div className="item" id={this.props.id} key={this.props.key}>
                <i className="large github middle aligned icon"/>
                <div className="content">
                    <a className="header" id={this.props.link} href={this.props.href} >{this.props.repoName}</a>
                </div>
                <i className="large window minimize middle aligned link icon"
                   onClick={() => removeRepository(this.props.repoIndex)}/>
            </div>
        )
    }
}

export default GithubRepoComponent;
