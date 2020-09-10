import React, {Component} from 'react';
import {Reveal, Image} from "semantic-ui-react";
import addPictureJPG from '../../Pics/addPictureJPG.jpg';
import addPicturePng from '../../Pics/addPicturePng.png';
import {
    acceptedImageExtensions,
    addPictureInputChanged,
    choosePicture,
} from "../../Js Functionals/ProfilePage/profilePictureContent";

class ProfilePictureComponent extends Component {
    render() {
        return (
            <div className='link' onClick={() => choosePicture()} style={{cursor: 'pointer'}}>
                <Reveal animated='fade' className={this.props.className} as='a'>
                    <Reveal.Content visible>
                        <Image id={this.props.imageId} src={this.props.src} size='small' alt={this.props.alt}/>
                    </Reveal.Content>
                    <Reveal.Content hidden>
                            <Image src={addPictureJPG} size='small'/>
                            <input type='file' style={{display: 'none'}} id={this.props.addImageInputId}
                                   accept={acceptedImageExtensions} onChange={() => addPictureInputChanged()}/>
                    </Reveal.Content>
                </Reveal>
            </div>
        );
    }
}

export default ProfilePictureComponent;
