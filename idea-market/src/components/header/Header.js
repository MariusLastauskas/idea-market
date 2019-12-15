import React, { useState } from 'react';
import Button from '../button/Button';
import Logo from '../logo/Logo';
import Menu from '../menu/Menu';
import Modal from '../modal/Modal';
import Image from '../image/Image';
import { TYPE } from '../modal/constants';
import { getCookie, deleteCookie, api, jwtDecode } from '../../utils';
import hamburger from '../../imgs/hamburger.svg';
import './header.scss';

const Header = ({ onRouteChange }) => {
	const jwtToken = getCookie('jwtToken');
	const userData = jwtToken ? jwtDecode(atob(jwtToken)) : null;
	const mainClass = 'header';

	const [isLoggedIn, setIsLoggedIn] = useState(jwtToken);
	const [isLogginModalOpen, setIsLogginModalOpen] = useState(false);
	const [menuOpen, setMenuOpen] = useState(false);

	jwtToken &&
		api(`http://localhost:8080/user/${userData.id}`, 'GET').then(function (response) {
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
		<>
			<header className={mainClass}>
				<div className={`${mainClass}__wrapper`}>
					<Logo onRouteChange={onRouteChange} />
					<div className={`${mainClass}__container`}>
						<Image className={`${mainClass}__hamburger`} src={hamburger} onClick={() => setMenuOpen(!menuOpen)} />
						<Menu className={menuOpen ? '' : `${mainClass}--closed`} role={userData ? userData.role : -1} onRouteChange={onRouteChange} />
						<Button className={menuOpen ? '' : `${mainClass}--closed`} text={isLoggedIn ? 'Log out' : 'Log in'} onClick={isLoggedIn ? logOut : openLogIn} />
					</div>
				</div>
			</header>
			{isLogginModalOpen && (
				<Modal
					onClose={() => {
						setIsLogginModalOpen(false);
					}}
					label="Log in"
					type={TYPE.LOGIN}
				/>
			)}
		</>
	);
};

export default Header;
