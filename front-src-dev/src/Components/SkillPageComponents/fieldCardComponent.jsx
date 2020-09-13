import React, {Component} from 'react';

class FieldCardComponent extends Component {
    render() {
        return (
            <div>
                <Card>
                    <i>
                    </i>
                    <Card.Content>
                        <Card.Header>
                            {this.props.skillName}
                        </Card.Header>
                    </Card.Content>
                </Card>
            </div>
        );
    }
}

export default FieldCardComponent;
