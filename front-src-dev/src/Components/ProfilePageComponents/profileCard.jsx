import React, {Component} from 'react';
import '../../CSS Designs/basic.css';
import '../../CSS Designs/ProfilePage/CSS1.css';
import '../../CSS Designs/Header.css';

class ProfileCard extends Component {
    constructor(props, context) {
        super(props, context);
    }

    render() {
        return (
            <div className="card flexColumn" id={this.props.id} onClick={this.props.onClick}>
                <header className="cardHeader flexRow">
                    <i className="check large icon myIcon" />
                    <h3 className="number" id={this.props.hId}>{this.props.number}</h3>
                </header>
                <div className="ui paragraph cardP">{this.props.cardContent}</div>
            </div>
        );
    }
}

export default ProfileCard;
