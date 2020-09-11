import React from 'react'
import {Modal} from 'semantic-ui-react'
import '../../../CSS Designs/ProfilePage/CSS1.css'
import ProfileCard from "./profileCard";
import {changePassword} from "../../Js Functionals/ProfilePage/passwordContent";
import MainInput from "../MainPageComponents/mainFormElements/mainInput";

// import {changePassword} from '../../Js Functionals/ProfilePage/JS1';

function ModalPassword() {
    const [open, setOpen] = React.useState(false)

    return (
        <Modal
            onClose={() => setOpen(false)}
            onOpen={() => setOpen(true)}
            open={open}
            trigger={<ProfileCard hId='changePassword' number={35} cardContent='تغییر رمز عبور'/>}
        >

            <Modal.Content className="ui form flexColumn modal" id="changingPasswordContent">
                <div className="three fields ui-rtl" id="passwordFields">

                    <MainInput type='password' id="oldPasswordField" placeHolder="Old Password"
                               textName='رمز عبور قدیمی' />

                    <MainInput type='password' id="passwordField" placeHolder="Password"
                               textName='رمز عبور' />

                    <MainInput type='password' id="repeatPasswordField" placeHolder="Repeat Password"
                               textName='تکرار رمز عبور' />

                </div>

                <button className="positive ui button rightAligned" id="changePasswordButton"
                        onClick={() => changePassword()}>تغییر رمز عبور
                </button>
            </Modal.Content>
        </Modal>
    );
}

export default ModalPassword;
