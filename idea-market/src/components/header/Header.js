import React, {useState} from 'react';
import Button from '../button/Button';
import Logo from '../logo/Logo';
import Menu from '../menu/Menu';
import Modal from '../modal/Modal';
import {getCookie, deleteCookie, api} from '../../utils';
import './header.scss';

const jwt_decode = require('jwt-decode');

const Header = ({onRouteChange}) => {
    const jwtToken = getCookie('jwtToken');
    const userData = jwtToken ? jwt_decode(atob(jwtToken)) : null;
    const mainClass = 'header';
    console.log(userData);

    const [isLoggedIn, setIsLoggedIn] = useState(jwtToken);
    const [isLogginModalOpen, setIsLogginModalOpen] = useState(false);

    jwtToken && api(`http://localhost:8080/user/${userData.id}`, 'GET')
        .then(function (response) {
            setIsLoggedIn(true);
        });

    const openLogIn = () => {
        setIsLogginModalOpen(true);
    };

    const logOut = () => {
        deleteCookie('jwtToken');
        setIsLoggedIn(false);
        window.location.reload();
    };

    return (
        <div className={mainClass}>
            <div className={`${mainClass}__wrapper`}>
                <Logo onRouteChange={onRouteChange}/>
                <Menu role={userData? userData.role : -1} onRouteChange={onRouteChange}/>
                <Button 
                    text={isLoggedIn ? 'Log out' : 'Log in'}
                    onClick={isLoggedIn ? logOut : openLogIn} />
            </div>
            {isLogginModalOpen && <Modal onClose={() => {setIsLogginModalOpen(false)}} label="Log in"/>}
        </div>
    );
};

export default Header;