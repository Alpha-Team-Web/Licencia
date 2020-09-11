import React from 'react';
import MainFormComponent from "./mainFormComponent";

class MainTextArea extends MainFormComponent {

    createMainFormElement() {
        return(
            <textarea maxLength={this.props.maxLength} type={this.props.type ? this.props.type : "text"}
                   placeholder={this.props.placeHolder}
                   id={this.props.id}>
                {this.props.children}
            </textarea>
        )
    }

}

export default MainTextArea;
