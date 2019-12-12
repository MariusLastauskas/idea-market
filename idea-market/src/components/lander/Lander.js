import React from 'react';
import Footer from '../footer/Footer';
import Header from '../header/Header';
import Image from '../image/Image';
import LandingImage from '../../imgs/landingImage.jpg';

import './lander.scss';

const Lander = () => {
    return (
        <div className="lander">
            <Header />
            <Image className="lander__image" src={LandingImage} />
            <div className="lander__container">
                <h1 className="lander__slogon">Find the idea you are looking for TODAY!!!</h1>
            </div>
            <Footer />
        </div>
    );
};

export default Lander;