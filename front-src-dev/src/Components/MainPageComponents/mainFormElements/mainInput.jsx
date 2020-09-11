import React from 'react';
import MainFormComponent from "./mainFormComponent";

class MainInput extends MainFormComponent {

    createMainFormElement() {
        return (
            <input maxLength={this.props.maxLength} type={this.props.type ? this.props.type : "text"}
                   placeholder={this.props.placeHolder} onBlur={this.props.onBlur}
                   readOnly={this.props.readOnly} id={this.props.id}>
                {this.props.children}
            </input>
        )
    }

}

export default MainInput;
