import React from 'react';
import './input.scss';

const Input = ({label, name, placeholder, type, onChange, value}) => {
    const mainClassName = 'input';

    return (
        <div className={mainClassName}>
            <label className={`${mainClassName}__label`} htmlFor={name}>{label}</label>
            <input 
                className={`${mainClassName}__input`} 
                id={name} 
                name={name} 
                type={type} 
                placeholder={placeholder}
                value={value}
                onChange={onChange}
            />
        </div>
    );
};

export default Input;