import React from 'react';
import './link.scss';

const Link = ({className, onRouteChange, text}) => {
    return (
        <span className={`link ${className}`} onClick={onRouteChange}>{text}</span>
    );
};

export default Link;