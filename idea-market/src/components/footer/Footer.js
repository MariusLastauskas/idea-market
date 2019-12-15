import React from 'react';
import './footer.scss';
import Link from '../link/Link';

const Footer = () => {
    return (
        <footer className="footer">
            <div className="footer__section">
                <h2 className="footer__title">Technologys used:</h2>
                <ul className="footer__list">
                    <li className="footer__list-item">database: mariaDB</li>
                    <li className="footer__list-item">backend: GoLang</li>
                    <li className="footer__list-item">frontend: React hooks, sass</li>
                </ul>
            </div>
            <div className="footer__section">
                <h2 className="footer__title">Project purpose:</h2>
                <p className="footer__list-item">Web developement practice, experiments with different tools.</p>
            </div>
            <div className="footer__section">
                <h2 className="footer__title">Project is accessible at:</h2>
                <a className="footer__list-item" href="https://github.com/MariusLastauskas/idea-market">github</a>
            </div>
        </footer >
    );
};

export default Footer;