import {httpExcGetFile} from "../AlphaAPI";
import {isFreeLancer} from "./profilePageContent";
import {uploadProfilePicUrlEmployer, uploadProfilePicUrlFreelancer} from "../urlNames";
export const acceptedImageExtensions = '.png, .jpg, .jpeg, .bmp';
export const maximumImageSize = 5;    //Mb
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
        let fileSize = imageInput.files.item(0).size;
        let extension = getExtension(imageInput.value)
        if (!acceptedImageExtensions.includes(extension.toLowerCase())) {
            alert('Ridi. Dorost Entekhab Kon')
            emptyAddedValues()
        } else if (fileSize / imageSizeUnit > maximumImageSize) {
            alert('Sizesh Ziade')
            emptyAddedValues()
        } else {
            let reader = new FileReader();
            reader.onload = function (e) {
                profileImage.src = e.target.result;
                addedImageValue = e.target.result;
            }
            reader.readAsDataURL(imageFile);
        }
    }

}

function getExtension(formData) {
    let splitFile = formData.split('.')
    return splitFile[splitFile.length - 1];
}

export function saveProfilePicture(){
    if(originImageValue!==profileImage){
        let url = isFreeLancer?uploadProfilePicUrlFreelancer:uploadProfilePicUrlEmployer;
        let header={
            token:Cookies.get('auth'),
        }
        httpExcGetFile('POST', url, profileImage, ()=>{
            alert('successSendPic')
        }, ()=>{
            alert('failSendPic')
        }, )
    }
}
