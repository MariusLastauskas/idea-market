import React from 'react';
// import './image.scss';

const Image = ({ className, src, onClick }) => {
    return (
        <img className={className} src={src} onClick={onClick} />
    );
};

export default Image;