import React, {Component} from 'react'
import {Button, Divider, Image, Transition} from 'semantic-ui-react'

export default class TransitionComponent extends Component {
    constructor(props, context) {
        super(props, context);
        this.state = {
            visible: props.visibility
        };
        this.changeState = this.changeState.bind(this);
    }

    state;

    changeState() {
        alert("visible before : " + this.state.visible)
        this.setState((prevState) => {
            {
                alert('visible alan:' + prevState.visible)
                return {visible: !prevState.visible};
                
            }
        })
        alert("visible now : " + this.state.visible)
    }

    render() {
        alert('visibility: ' + this.state.visible)

        return (
            <Transition  visible={this.state.visible} animation={this.props.animation}
                        duration={500}>
                {this.props.content}
            </Transition>
        )
    }
}
