import React, { Component } from 'react'
import { Button, Divider, Image, Transition } from 'semantic-ui-react'

export default class TransitionComponent extends Component {
    constructor(props, context) {
        super(props, context);
    }

    state = { visible: true }

    render() {
        // this.toggleVisibility.bind(this)
        const { visible } = this.state

        return (
                <Transition id={this.props.id} visible={visible} animation={this.props.animation} duration={this.props.duration == undefined ? 500 : this.props.duration}>
                    {this.props.content}
                </Transition>
        )
    }
}
