import React, {Component} from 'react';
import {
    BrowserRouter as Router, Link, Route
} from 'react-router-dom';
import MainPage from "./Pages/mainPage";
import ProfilePage from "./Pages/profilePage";
import {mainPagePath, profilePagePath} from "./Js Functionals/PagePaths";

class App extends Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <Router>
                <Route exact path={mainPagePath}>
                    <MainPage/>
                </Route>
                <Route path={profilePagePath}>
                    <ProfilePage/>
                </Route>
            </Router>
        );
    }
}

App.propTypes = {};

export default App;
