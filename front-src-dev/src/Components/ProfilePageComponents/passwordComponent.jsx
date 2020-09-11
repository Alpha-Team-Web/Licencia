import React from 'react'
import {Button, Header, Image, Modal} from 'semantic-ui-react'
import '../../CSS Designs/ProfilePage/CSS1.css'
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
                <div className="three fields" id="passwordFields">

                    {/*<div className="field">
                        <label className="rightAligned">رمز عبور قدیمی</label>
                        <input type='password' id="oldPasswordField" placeholder="Old Password"/>
                    </div>*/}
                    <MainInput type='password' id="oldPasswordField" placeHolder="Old Password"
                               textName='رمز عبور قدیمی' />

                    {/*<div className="field">
                        <label className="rightAligned">رمز عبور</label>
                        <input type='password' id="passwordField" placeholder="Password"/>
                    </div>*/}
                    <MainInput type='password' id="passwordField" placeHolder="Password"
                               textName='رمز عبور' />

                    {/*<div className="field">
                        <label className="rightAligned">تکرار رمز عبور</label>
                        <input type='password' id="repeatPasswordField" placeholder="Repeat Password"/>
                    </div>*/}
                    <MainInput type='password' id="repeatPasswordField" placeHolder="Repeat Password"
                               textName='تکرار رمز عبور' />

                </div>

                <button className="positive ui button rightAligned" id="changePasswordButton"
                        onClick={() => {
                            setOpen(false);
                            changePassword();
                        }}>تغییر رمز عبور
                </button>
            </Modal.Content>
        </Modal>
    );
}

export default ModalPassword;
