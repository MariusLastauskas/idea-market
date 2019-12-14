import React from 'react';
import Link from '../link/Link';
import {ROUTE} from '../content/constants';
import './logo.scss';

const Logo = ({onRouteChange}) => {
    return (
        <Link className="logo" onRouteChange={() => {onRouteChange(ROUTE.LANDING)}} text="Idea Market" />
    );
};

export default Logo;