import React, {useState} from 'react';
import Button from '../button/Button';
import Logo from '../logo/Logo';
import Menu from '../menu/Menu';
import Modal from '../modal/Modal';
import {getCookie, deleteCookie} from '../../utils';
import './header.scss';

const axios = require('axios');


const Header = () => {
    const mainClass = 'header';
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [isLogginModalOpen, setIsLogginModalOpen] = useState(false);

    // axios.get('http://localhost:8080/article/3')
    // .then(function (response) {
    //     setIsLoggedIn(true);
    //     console.log(response);
    // });

    const openLogIn = () => {
        setIsLogginModalOpen(true);
    };

    const logOut = () => {
        deleteCookie('jwtToken');
        setIsLoggedIn(false);
    };

    return (
        <div className={mainClass}>
            <div className={`${mainClass}__wrapper`}>
                <Logo />
                <Menu />
                <Button 
                    text={isLoggedIn ? 'Log out' : 'Log in'}
                    onClick={isLoggedIn ? logOut : openLogIn} />
            </div>
            {isLogginModalOpen && <Modal onClose={() => {setIsLogginModalOpen(false)}} label="Log in"/>}
        </div>
    );
};

export default Header;