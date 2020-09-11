import React from 'react';
import MainFormComponent from "./mainFormComponent";

class MainSelect extends MainFormComponent {

    createMainFormElement() {
        return(
            <select id={this.props.id}>
                {this.createElements()}
                {this.props.children}
            </select>
        )
    }

    createElements = () => this.props.options.map(this.createElement)
    createElement = (option) => <option value={option.value}>{option.child}</option>

}

export default MainSelect;
