import React, {useState} from 'react';
import Button from '../button/Button';
import Input from '../input/Input';
import {addCookie, api} from '../../utils';

import './modal.scss';

import {TYPE} from './constants';

const Modal = ({label, onClose}) => {
    const mainClassName = 'modal';

    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [errorMsg, setErrorMsg] = useState('');

    const logIn = () => {
        addCookie('username', username);
        addCookie('passHash', password);

        api('http://localhost:8080/login', 'POST')
            .then(response => {
                console.log(response.data);
                addCookie('jwtToken', response.data);
                window.location.reload();
            }).catch(errpr => {
                setErrorMsg('Bad username or password');
            });

        // axios.post('http://localhost:8080/login')
        //     .then(response => {
        //         console.log(response.data);
        //         addCookie('jwtToken', response.data);
        //         alert(response);
        //     }).then(response => {
        //         alert(response);
        //     });
    };

    return (
        <div className={mainClassName}>
            <div className={`${mainClassName}__container`}>
                <Button className={`${mainClassName}__close`} text="x" onClick={onClose} />
                <h2 className={`${mainClassName}__label`}>{label}</h2>
                <Input 
                    label="Username" 
                    name="username" 
                    placeholder="Enter username" 
                    type="text" 
                    value={username} 
                    onChange={
                        e => {
                            setUsername(e.target.value);
                        }
                    } 
                />
                <Input 
                    label="Password" 
                    name="password" 
                    placeholder="Enter password" 
                    type="password" 
                    value={password}
                    onChange={
                        e => {
                            setPassword(e.target.value);
                        }
                    }
                />
                <p className={`${mainClassName}__error`}>{errorMsg}</p>
                <Button text="Log In" onClick={logIn} />
            </div>
        </div>
    );
};

export default Modal;