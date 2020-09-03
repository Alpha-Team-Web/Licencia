import React, {Component} from 'react';
import {
    BrowserRouter as Router, Route
} from 'react-router-dom';

class App extends Component {
    constructor(props) {
        super(props);

    }

    render() {
        return (
            <Router>
                <Route path="/MainPage">

                </Route>
                <Route path="/ProfilePage">
                    <p>klsjdfklsdjfkljsd</p>
                    <Second/>
                </Route>
            </Router>
        );
    }
}

App.propTypes = {};

export default App;
