import {httpExcGetFile} from "../AlphaAPI";
import {isFreeLancer} from "./profilePageContent";
import {uploadProfilePicUrlEmployer, uploadProfilePicUrlFreelancer} from "../urlNames";
import Cookies from 'js-cookie';
import {reload} from "../PageRouter";

export const acceptedImageExtensions = '.png, .jpg, .jpeg, .bmp';
export const maximumImageSize = 8;    //Mb
export const imageSizeUnit = 1024 * 1024;

let imageInput
let profileImage

export function fillPictureFields() {
    imageInput = document.getElementById('addPictureInput')
    profileImage = document.getElementById('profilePicture')
}

let addedImageValue;

function emptyAddedValues() {
    addedImageValue = '';
    imageInput.value = '';
}

let originImageValue;

export function fillProfileImage(data) {
    fillPictureFields()
    originImageValue = data;
    profileImage.src = getImageSrc(data);
    //todo
}

export function choosePicture() {
    fillPictureFields()
    imageInput.click();
}

export function addPictureInputChanged() {
    if (imageInput.value !== '') {
        fillPictureFields();
        let imageFile = imageInput.files.item(0);
        if (!acceptedImageExtensions.includes(getExtension(imageFile.name).toLowerCase())) {
            alert('Ridi. Dorost Entekhab Kon')
            emptyAddedValues()
        } else if (imageFile.size / imageSizeUnit > maximumImageSize) {
            alert('Sizesh Ziade')
            emptyAddedValues()
        } else {
            let reader = new FileReader();
            reader.onload = function (e) {
                profileImage.src = e.target.result;
            }
            reader.readAsDataURL(imageFile);
            addedImageValue = imageFile;
        }
    }
}

function getExtension(formData) {
    let splitFile = formData.split('.')
    return splitFile[splitFile.length - 1];
}

let getImageSrc = (data) => 'data:image/' + getExtension(data.name) + ';base64,' + data.data;

export function saveProfilePicture() {
    if (addedImageValue && addedImageValue !== originImageValue) {
        let imageData = new FormData();
        imageData.append('profileImage', addedImageValue)
        let url = isFreeLancer ? uploadProfilePicUrlFreelancer : uploadProfilePicUrlEmployer;
        httpExcGetFile('POST', url, imageData, successPictureUpload, denyPictureUpload, {
            'Token': Cookies.get('auth'),
        })
    }
}

function successPictureUpload(value) {
    alert('successSendPic')
    alert('message: ' + value.message)
    reload()
}

function denyPictureUpload(value) {
    alert('failSendPic')
    alert('message: ' + value.message)
}
