import React from 'react';
// import './image.scss';

const Image = ({className, src}) => {
    return (
        <img className={className} src={src} />
    );
};

export default Image;