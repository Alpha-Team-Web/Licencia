import React, {Component} from 'react';
import {Reveal, Image} from "semantic-ui-react";
import addPictureJPG from '../../Pics/addPictureJPG.jpg';
import addPicturePng from '../../Pics/addPicturePng.png';
import {choosePicture} from "../../Js Functionals/ProfilePage/profilePictureContent";


class ProfilePictureComponent extends Component {
    render() {
        return (
            <Reveal animated='fade' className={this.props.className}>
                <Reveal.Content visible>
                    <Image id={this.props.imageId} src={this.props.src} size='small' alt={this.props.alt}/>
                </Reveal.Content>
                <Reveal.Content hidden>
                    <Image src={addPicturePng} size='small' onClick={() => choosePicture()}/>
                </Reveal.Content>
            </Reveal>
        );
    }
}

export default ProfilePictureComponent;
