import MainTextField from "../../../Components/MainPageComponents/mainTextField";

export function hasEmpty(...args) {
    for (let doc of args) {
        if (doc.value === "") {
            return doc;
        }
    }
    return null;
}

export function emptyFieldsFromErrors(...fields) {
    fields.forEach((value => {
        if (value) {
            MainTextField.setFieldError(value, false)
            MainTextField.showErrorLabel(value, false)
        }
    }))
}
