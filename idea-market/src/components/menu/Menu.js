import React from 'react';
import Link from '../link/Link';
import './menu.scss';

const Menu = () => {
    return (
        <ul className="menu">
            <li className="menu__link">
                <Link text="admin panel" href="admin" />
            </li>
            <li className="menu__link">
                <Link text="my profile" href="myProfile" />
            </li>
            <li className="menu__link">
                <Link text="projects" href="projects" />
            </li>
            <li className="menu__link">
                <Link text="add project" href="addProject" />
            </li>
            <li className="menu__link">
                <Link text="articles" href="articles" />
            </li>
        </ul>
    );
};

export default Menu;