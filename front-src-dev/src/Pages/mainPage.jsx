import React, {Component} from 'react';
import MainHeader from "../Components/MainPageComponents/mainHeader";
import MainContent from "../Components/MainPageComponents/mainContent";
import ModalLogSin from "../Components/MainPageComponents/LogSinMenuModal";
import MyModal from "../Components/ProfilePageComponents/ProfileSectionComponents/passwordComponent";

class MainPage extends Component {
    render() {
        return (
            <React.Fragment>
                <MainHeader/>
                <MainContent/>
            </React.Fragment>
        );
    }
}

export default MainPage;
