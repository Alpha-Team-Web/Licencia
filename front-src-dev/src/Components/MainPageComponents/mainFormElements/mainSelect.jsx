import React, from 'react';
import MainFormComponent from "./mainFormComponent";

class MainSelect extends MainFormComponent {

    createMainFormElement() {
        return(
            <select maxLength={this.props.maxLength} type={this.props.type ? this.props.type : "text"}
                      placeholder={this.props.placeHolder}
                      id={this.props.id}>
                {this.createElements(this.props.options)}
                {this.props.children}
            </select>
        )
    }

    createElements = () => this.props.options.map(this.createElement)
    createElement = (option) => <option value={option.value}>{option.child}</option>

}

export default MainSelect;
