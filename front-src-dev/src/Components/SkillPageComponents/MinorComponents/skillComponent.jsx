import React, {Component} from 'react';
import '../../../CSS Designs/extra-css.css'


class SkillComponent extends Component {

    constructor(props, context) {
        super(props, context);
        this.props.skillIncludes ? document.getElementById("functionIcon").className += " close" : document.getElementById("functionIcon").className += " add"
    }

    render() {
        return (
            <div className="skillComponent" onMouseOver={
                function (){
                    let funcBtn = document.getElementById("functionIcon")
                     let fakeIcon = document.getElementById("fakeIcon")
                     fakeIcon.style.display = "none"
                    funcBtn.style.display = "block"
                }
            }
            onMouseOut={
                function () {
                    let funcBtn = document.getElementById("functionIcon")
                    let fakeIcon = document.getElementById("fakeIcon")
                    fakeIcon.style.display = "block"
                    funcBtn.style.display = "none"
                }
            }>
                <div>
                    <i className="close link icon iconButton" >
                    </i>
                   {this.props.contentText}
                </div>
                <i className="link icon iconButton" aria-hidden="true" id="functionIcon" onClick={this.props.skillIncludes ? this.props.skillDeleter : this.props.skillAdder} >
                </i>
                <i id="fakeIcon"></i>
            </div>
        );
    }

}

export default SkillComponent;
