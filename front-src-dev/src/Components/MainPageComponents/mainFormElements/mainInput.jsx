import React, {Component} from 'react';
import '../../../CSS Designs/MainPage/loginSignupInput.css';
import '../../CSS Designs/MainPage/LoginMenu.css';
import {setFieldError} from "../../Js Functionals/MainPage/IOMethods/Utils/handleErrors";

class MainInput extends Component {
    constructor(props, context) {
        super(props, context);
    }

    errorLabelStyle = {
        display: 'none',
        maxWidth: '150px',
        textAlign: 'center'
    }

    render() {
        return (
            <div className="ui form formPadding">
                <div className="ui field">
                    <p className="paragraphInput">{this.props.textName}</p>
                    <input maxLength={this.props.maxLength} type={this.props.type ? this.props.type : "text"}
                           placeholder={this.props.placeHolder}
                           id={this.props.id}/>
                    <div className="ui pointing label red" id={this.props.errorId} style={this.errorLabelStyle}>
                        {this.props.errorText}
                    </div>
                </div>
            </div>
        );
    }


    componentDidMount() {
        let field = document.getElementById(this.props.id)
        field.addEventListener('focus', () => {
            setFieldError(field, false)
        })
    }


}

export default MainInput;
