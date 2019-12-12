import React from 'react';
import Image from '../image/Image';
import './card.scss';

const Card = ({title, avatar, image, description}) => {
    return (
        <div className="card">
            <Image src={image} />
            <h2>{title}</h2>

        </div>
    );
};

export default Card;