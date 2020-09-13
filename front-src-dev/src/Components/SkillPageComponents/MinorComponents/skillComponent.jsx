import React, {Component} from 'react';
import '../../../CSS Designs/extra-css.css'


class SkillComponent extends Component {

    constructor(props, context) {
        super(props, context);
    }

    render() {
        return (
            <div className="skillComponent" onMouseOver={
                function (){
                    let deleteBtn = document.getElementById("deleteSkillIcon")
                     let fakeIcon = document.getElementById("fakeIcon")
                     fakeIcon.style.display = "none"
                    deleteBtn.style.display = "block"
                }
            }
            onMouseOut={
                function () {
                    let deleteBtn = document.getElementById("deleteSkillIcon")
                    let fakeIcon = document.getElementById("fakeIcon")
                    fakeIcon.style.display = "block"
                    deleteBtn.style.display = "none"
                }
            }>
                <div>
                    <i className="close link icon iconButton" >
                    </i>
                   {this.props.contentText}
                </div>
                <i className="close link icon iconButton" aria-hidden="true" id="deleteSkillIcon" >
                </i>
                <i id="fakeIcon"></i>
            </div>
        );
    }

}

export default SkillComponent;
