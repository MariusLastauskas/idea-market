import React from 'react';
import Link from '../link/Link';
import {ROUTE} from '../content/constants';
import './menu.scss';

const Menu = ({role, onRouteChange}) => {
    return (
        <ul className="menu">
            {role === 1 && 
                <li className="menu__link">
                    <Link text="admin panel" route={ROUTE.ADMIN_PANEL} onRouteChange={() => {onRouteChange(ROUTE.ADMIN)}} />
                </li>
            }
            {role !== -1 &&
                <li className="menu__link">
                    <Link text="my profile" route={ROUTE.MY_PROFILE} onRouteChange={() => {onRouteChange(ROUTE.PROFILE)}} />
                </li>
            }
            <li className="menu__link">
                <Link text="projects" route={ROUTE.PROJECTS} onRouteChange={() => {onRouteChange(ROUTE.PROJECT)}} />
            </li>
            <li className="menu__link">
                <Link text="articles" onRouteChange={() => {onRouteChange(ROUTE.ARTICLE)}} />
            </li>
        </ul>
    );
};

export default Menu;