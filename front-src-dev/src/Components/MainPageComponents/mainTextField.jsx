import React, {Component} from 'react';
import '../../CSS Designs/MainPage/loginSignupInput.css';
import '../../CSS Designs/MainPage/LoginMenu.css';

class MainTextField extends Component {
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
                    <input maxLength={this.props.maxLength} type={this.props.isPassword ? 'password' : "text"}
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
            MainTextField.setFieldError(field, false)
        })
    }

    static setFieldError(field, isError) {
        if (field) {
            if ((isError === undefined || isError) && !this.containsError(field)) {
                field.value = "";
                field.parentElement.classList.add("error");
            } else if (isError === false && field.parentElement.classList.contains("error")) {
                field.parentElement.classList.remove("error")
                // this.showErrorLabel(field, false)
            }
        }
    }

    static showErrorLabel(field, errorLabel) {
        if (field) {
            if (/*isError === undefined || isError*/errorLabel || errorLabel === '') {
                if (errorLabel) {
                    this.getLabel(field).innerHTML = errorLabel;
                }
                this.getLabel(field).style.display = 'block';
            } else {
                this.getLabel(field).style.display = 'none';
            }
        }
    }

    static getLabel = (field) => field ? field.parentElement.children.item(2) : null;
    static containsError = (field) => field && field.parentElement.classList.contains('error');
}

export default MainTextField;
