import React, {Component} from 'react';
import MainHeader from "../Components/MainPageComponents/mainHeader";
import MainContent from "../Components/MainPageComponents/mainContent";
import ModalExampleModal from "../Components/file";

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
