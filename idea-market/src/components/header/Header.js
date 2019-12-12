import React, {useState} from 'react';
import Button from '../button/Button';
import Logo from '../logo/Logo';
import Menu from '../menu/Menu';
import {getCookie} from '../../utils';
import './header.scss';

const axios = require('axios');


const Header = () => {
    const mainClass = 'header';
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    axios.get('http://localhost:8080/article/3')
        .then(function (response) {
            setIsLoggedIn(true);
            console.log(response);
        })
    return (
        <div className={mainClass}>
            <div className={`${mainClass}__wrapper`}>
                <Logo />
                <Menu />
                <Button text={isLoggedIn ? 'Log out' : 'Log in'} />
            </div>
        </div>
    );
};

export default Header;