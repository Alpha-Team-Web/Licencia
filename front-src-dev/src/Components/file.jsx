import React from 'react'
import { Button, Header, Image, Modal } from 'semantic-ui-react'
import {greenColor, loginMenu, signUpMenu} from "../Js Functionals/MainPage/Login SignUp Show";
import licenciaImg from "../Pics/Licencia-Logo.png";
import MainLoginMenu from "./MainPageComponents/mainLoginMenu";
import MainSignUpMenu from "./MainPageComponents/mainSignUpMenu";

function ModalExampleModal() {
    const [open, setOpen] = React.useState(false)

    let style1 = {
        display: "none"
    }
    let style2 = {
        backgroundColor: greenColor
    }

    return (
        <Modal
            onClose={() => setOpen(false)}
            onOpen={() => setOpen(true)}
            open={open}
            // trigger={<Button>Show Modal</Button>}
            trigger={<Button className="loginButton">ورود / ثبت نام</Button>}

        >
            {/*<Modal.Header>Select a Photo</Modal.Header>*/}
            <Modal.Content >
                {/*<Image size='medium' src='https://react.semantic-ui.com/images/avatar/large/rachel.png' wrapped />
                <Modal.Description>
                    <Header>Default Profile Image</Header>
                    <p>
                        We've found the following gravatar image associated with your e-mail
                        address.
                    </p>
                    <p>Is it okay to use this photo?</p>
                </Modal.Description>*/}
                <div className="header" id="Login-Menu-Header">
                    <div id="Signup-Login">
                        <div style={style2} className="Signup-login-text" id="LoginMenuButton" onClick={() => loginMenu()}>ورود</div>
                        <div className="Signup-login-text" id="SignUpMenuButton" onClick={() => signUpMenu()}>ثبت نام</div>
                    </div>
                    <div className="image content">
                        <img src={licenciaImg} id="logoImage" alt="logoLicencia"/>
                    </div>
                    <h3 id="welcomeHeader">Welcome To Licencia</h3>
                </div>
                <MainLoginMenu/>
                <MainSignUpMenu style = {style1}/>
            </Modal.Content>
            {/*<Modal.Actions>
                <Button color='black' onClick={() => setOpen(false)}>
                    Nope
                </Button>
                <Button
                    content="Yep, that's me"
                    labelPosition='right'
                    icon='checkmark'
                    onClick={() => setOpen(false)}
                    positive
                />
            </Modal.Actions>*/}
        </Modal>
    )
}

export default ModalExampleModal