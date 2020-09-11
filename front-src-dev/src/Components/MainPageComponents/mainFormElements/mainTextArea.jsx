import React from 'react';
import MainFormComponent from "./mainFormComponent";

class MainTextArea extends MainFormComponent {

    createMainFormElement() {
        return(
            <textarea maxLength={this.props.maxLength} placeholder={this.props.placeHolder}
                      rows={this.props.rows ? this.props.rows : '3'} cols={this.props.cols ? this.props.cols : '20'}
                      id={this.props.id}>
                {this.props.children}
            </textarea>
        )
    }

}

export default MainTextArea;
