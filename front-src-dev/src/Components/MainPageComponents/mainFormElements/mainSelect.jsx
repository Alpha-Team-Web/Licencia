import React, {Component} from 'react';
import MainFormComponent from "./mainFormComponent";

class MainSelect extends MainFormComponent {

    createMainFormElement() {
        return (
            <select className="ui dropdown" id={this.props.id}>
                <option value={this.props.value1}>{this.props.textValue1}</option>
                <option value={this.props.value2}>{this.props.textValue2}</option>
            </select>
        )
    }

}

export default MainSelect;
