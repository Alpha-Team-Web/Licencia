import React from 'react';
import MainFormComponent from "./mainFormComponent";

class MainInput2 extends MainFormComponent {

    createMainFormElement() {
        return (
            <input maxLength={this.props.maxLength} type={this.props.type ? this.props.type : "text"}
                   placeholder={this.props.placeHolder}
                   id={this.props.id}>
                {this.props.children}
            </input>
        )
    }

}

export default MainInput2;
