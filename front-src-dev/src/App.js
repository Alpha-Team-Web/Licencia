import React, {Component} from 'react';
import {
    BrowserRouter as Router, Link, Route
} from 'react-router-dom';
import MainPage from "./Pages/mainPage";
import ProfilePage from "./Pages/profilePage";
import {mainPageName, profilePageName} from "./Js Functionals/FileNames";

class App extends Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <Router>
                <Route exact path={mainPageName}>
                    <MainPage/>
                </Route>
                <Route path={profilePageName}>
                    <ProfilePage/>
                </Route>
            </Router>
        );
    }
}

App.propTypes = {};

export default App;
