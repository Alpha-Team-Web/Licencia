import React, { Component } from 'react'
import { Button, Divider, Image, Transition } from 'semantic-ui-react'

export default class TransitionComponent extends Component {
    state = {
        visible: this.props.visibility
    };

    render() {
        alert('visibility: ' + this.state.visible)

        return (
                <Transition id={this.props.id} visible={this.state.visible} animation={this.props.animation} duration={this.props.duration == undefined ? 500 : this.props.duration}>
                    {this.props.content}
                </Transition>
        )
    }
}
