import React, { useState } from 'react';
import './input.scss';

const Input = ({ label, name, placeholder, type, onChange, value }) => {
    const mainClassName = 'input';
    let Tag = type === 'textarea' ? 'textarea' : 'input';
    const [toggler, setToggler] = useState(value);

    console.log('render', value)

    return (
        <div className={mainClassName}>
            <label className={`${mainClassName}__label ${type === 'checkbox' && mainClassName + '__label--checkbox'} ${value && mainClassName + '__label--checked'}`} htmlFor={name}>{label}</label>
            <Tag
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