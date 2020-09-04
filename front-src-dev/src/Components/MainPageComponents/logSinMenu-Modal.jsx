import React from 'react'
import { Button, Header, Icon, Modal } from 'semantic-ui-react'
import MainLoginMenu from "./mainLoginMenu";
import MainSignUpMenu from "./mainSignUpMenu";

function getModalLogSinMenu() {
    const [open, setOpen] = React.useState(false)

    return (
        <Modal
            basic
            onClose={() => setOpen(false)}
            onOpen={() => setOpen(true)}
            open={open}
            size='small'
            trigger=/*{<Button>Basic Modal</Button>}*/
                {<Button className="loginButton" onClick="openLogSinMenu()">ورود / ثبت نام</Button>}
        >
            {/*<Header icon>
                <Icon name='archive' />
                Archive Old Messages
            </Header>*/}
            <Modal.Content>
                {/*<p>
                    Your inbox is getting full, would you like us to enable automatic
                    archiving of old messages?
                </p>*/}
                <div className="header" id="Login-Menu-Header">
                    <div id="Signup-Login">
                        <div className="Signup-login-text" id="LoginMenuButton" onClick="loginMenu()">ورود</div>
                        <div className="Signup-login-text" id="SignUpMenuButton" onClick="signUpMenu()">ثبت نام</div>
                    </div>
                    <div className="image content">
                        <img src="../Pics/Licencia-Logo.png" id="logoImage" alt="logoLicencia"/>
                    </div>
                    <h3 id="welcomeHeader">Welcome To Licencia</h3>
                </div>
                <MainLoginMenu/>
                <MainSignUpMenu/>
            </Modal.Content>
            {/*<Modal.Actions>
                <Button basic color='red' inverted onClick={() => setOpen(false)}>
                    <Icon name='remove' /> No
                </Button>
                <Button color='green' inverted onClick={() => setOpen(false)}>
                    <Icon name='checkmark' /> Yes
                </Button>
            </Modal.Actions>*/}
        </Modal>
    )
}

export default getModalLogSinMenu
