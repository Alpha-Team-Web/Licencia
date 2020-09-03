import React, {Component} from 'react';
import {cyan} from "color-name";

const headerTextColor = cyan;

class HeaderListItem extends Component {
    constructor(props, context) {
        super(props, context);
    }

    render() {
        return (
            <li className='Header-UnOrderedList-Item' value={this.props.value}>
                <a href={this.props.href} className='Header-UnOrderedList-Item'>{this.props.value}</a>
            </li>
        );
    }
}

export default HeaderListItem;
