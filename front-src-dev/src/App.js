import React, {Component} from 'react';
import {
    BrowserRouter as Router, Link, Route
} from 'react-router-dom';
import MainPage from "./Pages/mainPage";
import ProfilePage from "./Pages/profilePage";
import {mainPagePath, profilePagePath, skillsPagePath} from "./Js Functionals/PagePaths";
import SkillsPage from "./Pages/skillsPage";

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
                <Route path={skillsPagePath}>
                    <SkillsPage/>
                </Route>
            </Router>
        );
    }
}

App.propTypes = {};

export default App;
