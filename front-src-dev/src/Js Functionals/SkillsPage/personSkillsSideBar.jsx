import * as React from "react";
import {Grid, Checkbox, Sidebar, Segment, Menu, Icon, Header, Image} from "semantic-ui-react";

import {Component} from 'react';

class PersonSkillsSideBar extends Component {

    render() {

        return (
            <Sidebar.Pushable as={Segment}>
                <Sidebar
                    as={Menu}
                    animation='overlay'
                    icon='labeled'
                    inverted
                    onHide={() => this.props.setVisible(false)}
                    vertical
                    visible={this.props.visible}
                    width='thin'
                >
                    <Menu.Item as='a'>
                        <Icon name='home'/>
                        Home
                    </Menu.Item>
                    <Menu.Item as='a'>
                        <Icon name='gamepad'/>
                        Games
                    </Menu.Item>
                    <Menu.Item as='a'>
                        <Icon name='camera'/>
                        Channels
                    </Menu.Item>
                </Sidebar>

                <Sidebar.Pusher>
                    {this.props.children}
                </Sidebar.Pusher>
            </Sidebar.Pushable>
        )
    }

}

export default PersonSkillsSideBar;
