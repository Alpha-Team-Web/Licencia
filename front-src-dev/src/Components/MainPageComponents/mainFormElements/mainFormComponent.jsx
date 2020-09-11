import React, {Component} from 'react';
import {setFieldError} from "../../../Js Functionals/Utils/handleErrors";

class MainFormComponent extends Component {

    errorLabelStyle = {
        display: 'none',
        maxWidth: '150px',
        textAlign: 'center'
    }

    render() {
        return (
            <div className={"ui field formPadding ui-rtl " + this.props.className}>
                {/*<label className="rightAligned marginBottom10">{this.props.textName}</label>*/}
                <p className="paragraphInput">{this.props.textName}</p>
                {this.createMainFormElement()}
                <div className="ui pointing label red" id={this.props.errorId} style={this.errorLabelStyle}>
                    {this.props.errorText}
                </div>
            </div>
        );
    }

    createMainFormElement() {
        return (
            <div id={this.props.id}>
                {this.props.children}
            </div>
        )
    }

    componentDidMount() {
        let field = document.getElementById(this.props.id)
        field.addEventListener('focus', () => {
            setFieldError(field, false)
        })
    }

}

export default MainFormComponent;
