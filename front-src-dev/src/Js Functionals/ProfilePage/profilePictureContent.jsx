export const acceptedImageExtensions = '.png, .jpg, .jpeg, .bmp';
export const maximumImageSize = 5;    //Mb
export const imageSizeUnit = 1024 * 1024;

let imageInput
let originImage
export function fillPictureFields() {
    imageInput = document.getElementById('addPictureInput')
    originImage = document.getElementById('profilePicture')
}

let addedImageValue;
function emptyAddedValues() {
    addedImageValue = '';
    imageInput.value = '';
}

let originImageValue;
export function fillProfileImage(data) {
    //todo
}

export function choosePicture() {
    fillPictureFields()
    imageInput.click();
}

export function addPictureInputChanged() {
    alert('before Image: ' + addedImageValue)
    alert('added Image: ' + imageInput.files.length)
    if (imageInput.value !== '') {
        fillPictureFields();
        let fileSize = imageInput.files.item(0).size;
        let extension = getExtension(imageInput.value)

        /*if (acceptedImageExtensions.includes(extension)) {
            addedImageValue = imageInput.value;
        } else {
            alert('Ridi. Dorost Entekhab Kon')
            addedImageValue = '';
            imageInput.value = '';
        }*/

        alert('extension: ' + extension)
        alert('fileSize: ' + fileSize)

        if (!acceptedImageExtensions.includes(extension.toLowerCase())) {
            alert('Ridi. Dorost Entekhab Kon')
            emptyAddedValues()
        } else if(fileSize / imageSizeUnit > maximumImageSize) {
            alert('Sizesh Ziade')
            emptyAddedValues()
        } else {
            addedImageValue = imageInput.value;
        }
    }

}

function getExtension(formData) {
    let splitFile = formData.split('.')
    return splitFile[splitFile.length - 1];
}
