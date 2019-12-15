import React, { useState } from 'react';
import Button from '../button/Button';
import Input from '../input/Input';
import Image from '../image/Image';
import { addCookie, api } from '../../utils';

import './modal.scss';

import { TYPE } from './constants';
import { object } from 'prop-types';

const Modal = ({ label, onClose, type, object }) => {
	const mainClassName = 'modal';

	const [username, setUsername] = useState('');
	const [password, setPassword] = useState('');
	const [errorMsg, setErrorMsg] = useState('');

	const logIn = () => {
		addCookie('username', username);
		addCookie('passHash', password);

		api('http://localhost:8080/login', 'POST')
			.then((response) => {
				addCookie('jwtToken', response.data);
				window.location.reload();
			})
			.catch((errpr) => {
				setErrorMsg('Bad username or password');
			});
	};

	switch (type) {
		case TYPE.LOGIN:
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
							onChange={(e) => {
								setUsername(e.target.value);
							}}
						/>
						<Input
							label="Password"
							name="password"
							placeholder="Enter password"
							type="password"
							value={password}
							onChange={(e) => {
								setPassword(e.target.value);
							}}
						/>
						<p className={`${mainClassName}__error`}>{errorMsg}</p>
						<Button text="Log In" onClick={logIn} />
					</div>
				</div>
			);
		case TYPE.ARTICLE_INFO:
			return (
				<div className={mainClassName}>
					<div className={`${mainClassName}__container`}>
						<Button className={`${mainClassName}__close`} text="x" onClick={onClose} />
						<h2 className={`${mainClassName}__label`}>{object.title}</h2>
						<table>
							<tbody>
								<tr><td><h3>Content: </h3></td><td><p>{object.content}</p></td></tr>
								<tr><td><h3>Full text: </h3></td><td><p>{object.full_text}</p></td></tr>
								<tr><td><h3>Visibility	: </h3></td><td><p>{object.is_public ? 'public' : 'private'}</p></td></tr>
							</tbody>
						</table>
						<Button text="edit" onClick={() => { }} />
					</div>
				</div>
			);
		case TYPE.PROJECT_INFO:
			return (
				<div className={mainClassName}>
					<div className={`${mainClassName}__container`}>
						<Button className={`${mainClassName}__close`} text="x" onClick={onClose} />
						<h2 className={`${mainClassName}__label`}>{object.name}</h2>
						<table>
							<tbody>
								<tr><td><h3>description: </h3></td><td><p>{object.description}</p></td></tr>
								<tr><td><h3>Price: </h3></td><td><p>{object.price}</p></td></tr>
								<tr><td><h3>Multiplicity: </h3></td><td><p>{object.multiplicity === 0 ? 'free' : object.multiplicity}</p></td></tr>
								<tr><td><h3>Visibility: </h3></td><td><p>{object.is_public ? 'public' : 'private'}</p></td></tr>
								<tr><td><h3>Owner: </h3></td><td><p>{object.owner.username}</p></td></tr>
								<tr><td><h3>Buyers: </h3></td><td><ul>{object.buyers.map((buyer, key) => {
									return <li>{buyer.username}</li>
								})}</ul></td></tr>
							</tbody>
						</table>
						<Button text="edit" onClick={() => { }} />
						<Button text="buy" onClick={() => { }} />
						<Button text="follow" onClick={() => { }} />
					</div>
				</div>
			);
		case TYPE.USER_INFO:
			console.log(object);
			return (
				<div className={mainClassName}>
					<div className={`${mainClassName}__container`}>
						<Button className={`${mainClassName}__close`} text="x" onClick={onClose} />
						<h2 className={`${mainClassName}__label`}>{object.username}</h2>
						<table>
							<tbody>
								<tr><td><h3>Photo: </h3></td><td><Image className={`${mainClassName}__image`} src={object.photo_path} /></td></tr>
								<tr><td><h3>Full name: </h3></td><td><p>{object.full_name}</p></td></tr>
								<tr><td><h3>Email: </h3></td><td><p>{object.email}</p></td></tr>
								<tr><td><h3>Role: </h3></td><td><p>{object.role === 1 ? 'admin' : 'user'}</p></td></tr>
								<tr><td><h3>Activity: </h3></td><td><p>{object.is_active ? 'active' : 'blocked'}</p></td></tr>
								<tr><td><h3>Articles: </h3></td><td><ul>{object.articles.map((article, key) => {
									return <li>{article.title}</li>
								})}</ul></td></tr>
								<tr><td><h3>Owned projects: </h3></td><td><ul>{object.owned_projects.map((project, key) => {
									return <li>{project.name}</li>
								})}</ul></td></tr>
								<tr><td><h3>Bought projects: </h3></td><td><ul>{object.bought_projects.map((project, key) => {
									return <li>{project.name}</li>
								})}</ul></td></tr>
								<tr><td><h3>Followed projects: </h3></td><td><ul>{object.followed_projects.map((project, key) => {
									return <li>{project.name}</li>
								})}</ul></td></tr>
							</tbody>
						</table>
						<Button text="edit" onClick={() => { }} />
					</div>
				</div>
			);
		default:
			return <></>;
	}
};

export default Modal;
