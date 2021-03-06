import {setFieldError, showErrorLabel} from "./handleErrors";

export function hasEmpty(...fields) {
    for (let doc of fields) {
        if (doc.value === "") {
            return doc;
        }
    }
    return null;
}

export function emptyFields(...fields) {
    fields.forEach(value => {
        if (value) {
            value.value = '';
        }
    })
}

export function emptyFieldsFromErrors(...fields) {
    fields.forEach((value => {
        if (value) {
            setFieldError(value, false)
            showErrorLabel(value, false)
        }
    }))
}
