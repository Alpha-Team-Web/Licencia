export function setFieldError(field, isError) {
    if (field) {
        if ((isError === undefined || isError) && !containsError(field)) {
            field.value = "";
            field.parentElement.classList.add("error");
        } else if (isError === false && containsError(field)) {
            field.parentElement.classList.remove("error")
            // this.showErrorLabel(field, false)
        }
    }
}

export function showErrorLabel(field, errorLabel) {
    if (field) {
        if (/*isError === undefined || isError*/errorLabel || errorLabel === '') {
            if (errorLabel) {
                getLabel(field).innerHTML = errorLabel;
            }
            getLabel(field).style.display = 'block';
        } else {
            getLabel(field).style.display = 'none';
        }
    }
}

export let getLabel = (field) => field ? field.parentElement.children.item(2) : null;
export let containsError = (field) => field && field.parentElement.classList.contains('error');
