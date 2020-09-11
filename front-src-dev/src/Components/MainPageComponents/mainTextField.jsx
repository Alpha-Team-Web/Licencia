import React, {Component} from 'react';
import '../../CSS Designs/MainPage/loginSignupInput.css';
import '../../CSS Designs/MainPage/LoginMenu.css';
// import setFieldError from "./setFieldError";

class MainTextField extends Component {
    constructor(props, context) {
        super(props, context);
    }

    render() {
        return (
            <div className="ui form formPadding">
                <div className="ui field">
                    <p className="paragraphInput">{this.props.textName}</p>
                    <input maxLength={this.props.maxLength} type={this.props.isPassword ? 'password' : "text"} placeholder={this.props.placeHolder}
                           id={this.props.id} onFocus={MainTextField.setFieldError(this.props.id, false)} />
                    <div className="ui pointing label red" id={this.props.errorId} style={{display:'none'}}>
                        {this.props.errorText}
                    </div>
                </div>
            </div>
        );
    }

    static setFieldError(id, isError) {
        let field = document.getElementById(id)

        if (field != null) {
            if ((isError === undefined || isError) && !field.parentElement.classList.contains("error")) {
                // field.style.border = "1px solid red";
                field.parentElement.classList.add("error");
            } else if (!isError && field.parentElement.classList.contains("error")) {
                field.parentElement.classList.remove("error")
            }
        }
    }

}

export default MainTextField;
