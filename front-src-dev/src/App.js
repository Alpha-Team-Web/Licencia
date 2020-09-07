import React, {Component} from 'react';
import {
    BrowserRouter as Router, Route
} from 'react-router-dom';
import MainPage from "./Pages/mainPage";
import ProfilePage from "./Pages/profilePage";

class App extends Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <Router>
                <Route path="/dgdgdg">
                    <MainPage/>
                </Route>
                <Route path="/">
                    <ProfilePage/>
                </Route>
            </Router>
        );
    }
}

App.propTypes = {};

export default App;
