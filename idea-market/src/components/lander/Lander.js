import React, {useState} from 'react';
import Content from '../content/Content';
import {ROUTE} from '../content/constants';
import Footer from '../footer/Footer';
import Header from '../header/Header';
import Image from '../image/Image';
import LandingImage from '../../imgs/landingImage.jpg';

import './lander.scss';

const Lander = () => {
    const [route, setRoute] = useState(ROUTE.LANDING);

    return (
        <div className="lander">
            <Header onRouteChange={setRoute} />
            <Image className="lander__image" src={LandingImage} />
            <div className={`lander__container ${route === ROUTE.LANDING ? 'lander__container--centered' : ''}`}>
                <Content route={route} />
            </div>
            <Footer />
        </div>
    );
};

export default Lander;