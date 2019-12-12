import React from 'react';
import './link.scss';

const Link = ({className, href, text}) => {
    return (
        <a className={`link ${className}`} href={href}>{text}</a>
    );
};

export default Link;