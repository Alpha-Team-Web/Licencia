import React, {Component} from 'react';
import {Reveal, Image} from "semantic-ui-react";
import addPictureJPG from '../../../Pics/addPictureJPG.jpg';
import addPicturePng from '../../../Pics/addPicturePng.png';
import '../../../CSS Designs/extra-css.css'

import {
    acceptedImageExtensions,
    addPictureInputChanged,
    choosePicture,
} from "../../../Js Functionals/ProfilePage/profilePictureContent";

class ProfilePictureComponent extends Component {
    render() {
        return (
            <div className='link' id={this.props.id} onClick={() => choosePicture()} style={{cursor: 'pointer'}}>
                <Reveal animated='fade' className={this.props.className} as='a'>
                    <Reveal.Content visible>
                        <Image id={this.props.imageId} src={this.props.src} size='huge' alt={this.props.alt} fluid/>
                    </Reveal.Content>
                    <Reveal.Content hidden >
                            <Image src={addPicturePng} size='huge' fluid/>
                            <input type='file' style={{display: 'none'}} id={this.props.addImageInputId}
                                   accept={acceptedImageExtensions} onChange={() => addPictureInputChanged()}/>
                    </Reveal.Content>
                </Reveal>
            </div>
        );
    }
}

export default ProfilePictureComponent;
