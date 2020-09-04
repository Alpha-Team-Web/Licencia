import React, {Component, Fragment} from 'react';
import '../../CSS Designs/MainPage/content1.css';
import codingPersonImg from '../../Pics/codingPerson.jpg';
import MainLoginMenu from "./mainLoginMenu";

class MainContent extends Component {
    render() {
        return (
            <Fragment>
                <div className="content div-content" id="div-content1">
                    <div className="image content codingPersonImageDiv">
                        <img src={codingPersonImg} alt="a person is coding"/>
                    </div>
                    <div className="ui paragraph content-paragraphHeader">
                        <h2 id="textH">لیسنسیا بزرگ است و بزرگ میندیشد</h2>
                        <p id="textP">بزرگترین سایت ایران برای فریلنسر ها</p>
                        <div type="submit" className="ui green button" id="subscribeButton">عضوی از ما باشید</div>
                    </div>
                </div>

                <div className="content div-content">
                    <div className="image content codingPersonImageDiv">
                        <img src={codingPersonImg} alt="a person is coding"/>
                    </div>
                    <div className="ui paragraph content-paragraphHeader">
                        <h2 className="textH">لیسنسیا بزرگ است و بزرگ میندیشد</h2>
                        <p className="textP">بزرگترین سایت ایران برای فریلنسر ها</p>
                    </div>
                </div>

                <MainLoginMenu/>
            </Fragment>
        );
    }
}

export default MainContent;
