import React, { useState } from 'react';
import Button from '../button/Button';
import Input from '../input/Input';
import Image from '../image/Image';
import { addCookie, api } from '../../utils';

import './modal.scss';

import { TYPE } from './constants';
import { object } from 'prop-types';
import Axios from 'axios';

const Modal = ({ label, onClose, type, object, isEdit }) => {
	const mainClassName = 'modal';

	const [image, setImage] = useState('');
	const [fullName, setFullName] = useState('');
	const [email, setEmail] = useState('');
	const [username, setUsername] = useState('');
	const [password, setPassword] = useState('');
	const [errorMsg, setErrorMsg] = useState('');
	const [title, setTitle] = useState('');
	const [fullText, setFullText] = useState('');
	const [summary, setSummary] = useState('');
	const [isPublic, setIsPublic] = useState(false);
	const [isLimited, setIsLimited] = useState(false);
	const [quantity, setQuantity] = useState(0);
	const [price, setPrice] = useState(0);
	const [isRegister, setIsRegister] = useState(false);

	const logIn = () => {
		addCookie('username', username);
		addCookie('passHash', password);

		api('http://localhost:8080/login', 'POST')
			.then((response) => {
				addCookie('jwtToken', response.data);
				window.location.reload();
			})
			.catch((error) => {
				setErrorMsg('Bad username or password');
			});
	};

	const createArticle = () => {
		api('http://localhost:8080/article', 'POST',
			{
				"Title": title,
				"Content": summary,
				"Full_text": fullText,
				"IsPublic": false
			})
			.then((response) => {
				console.log(response);
				window.location.reload();
			});
	};

	const createProject = () => {
		api('http://localhost:8080/project', 'POST',
			{
				"Name": title,
				"Description": summary,
				"Is_public": isPublic,
				"Price": Number(price),
				"Multiplicity": isLimited ? Number(quantity) : 0
			})
			.then((response) => {
				console.log(response);
				window.location.reload();
			});
	};

	const createUser = () => {
		api('http://localhost:8080/user', 'POST',
			{
				"full_name": fullName,
				"username": username,
				"email": email,
				"password_hash": password,
				"photo_path": image
			})
			.then((response) => {
				console.log(response);
				window.location.reload();
			});
	}

	switch (type) {
		case TYPE.LOGIN:
			if (!isRegister) {
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
							<Button text="Create account" onClick={() => { setIsRegister(!isRegister) }} />
						</div>
					</div>
				);
			} else {
				return (
					<div className={mainClassName}>
						<div className={`${mainClassName}__container`}>
							<Button className={`${mainClassName}__close`} text="x" onClick={onClose} />
							<h2 className={`${mainClassName}__label`}>Create account</h2>
							<Input
								label="Full name"
								name="full_name"
								placeholder="Enter your full name"
								type="text"
								value={fullName}
								onChange={(e) => {
									setFullName(e.target.value);
								}}
							/>
							<Input
								label="Username"
								name="username"
								placeholder="Enter your username"
								type="text"
								value={username}
								onChange={(e) => {
									setUsername(e.target.value);
								}}
							/>
							<Input
								label="Email"
								name="email"
								placeholder="Enter your email"
								type="email"
								value={email}
								onChange={(e) => {
									setEmail(e.target.value);
								}}
							/>
							<Input
								label="Password"
								name="password"
								placeholder="Enter your password"
								type="password"
								value={password}
								onChange={(e) => {
									setPassword(e.target.value);
								}}
							/>
							<Input
								label="Avatar photo path"
								name="photo"
								placeholder="Enter your photo path"
								type="text"
								value={image}
								onChange={(e) => {
									setImage(e.target.value);
								}}
							/>
							<p className={`${mainClassName}__error`}>{errorMsg}</p>
							<Button text="Create article" onClick={createUser} />
						</div>
					</div>
				);
			}
		case TYPE.ARTICLE_INFO:
			if (!isEdit) {
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
			} else {
				return (
					<div className={mainClassName}>
						<div className={`${mainClassName}__container`}>
							<Button className={`${mainClassName}__close`} text="x" onClick={onClose} />
							<h2 className={`${mainClassName}__label`}>{label}</h2>
							<Input
								label="Title"
								name="title"
								placeholder="Enter title"
								type="text"
								value={title}
								onChange={(e) => {
									setTitle(e.target.value);
								}}
							/>
							<Input
								label="Summary"
								name="content"
								placeholder="Enter article summary"
								type="textarea"
								value={summary}
								onChange={(e) => {
									setSummary(e.target.value);
								}}
							/>
							<Input
								label="Full text"
								name="full_text"
								placeholder="Enter full text"
								type="textarea"
								value={fullText}
								onChange={(e) => {
									setFullText(e.target.value);
								}}
							/>
							<p className={`${mainClassName}__error`}>{errorMsg}</p>
							<Button text="Create article" onClick={createArticle} />
						</div>
					</div>
				);
			};
		case TYPE.PROJECT_INFO:
			if (!isEdit) {
				return (
					<div className={mainClassName}>
						<div className={`${mainClassName}__container`}>
							<Button className={`${mainClassName}__close`} text="x" onClick={onClose} />
							<h2 className={`${mainClassName}__label`}>{object.name}</h2>
							<table>
								<tbody>
									<tr><td><h3>description: </h3></td><td><p>{object.description}</p></td></tr>
									<tr><td><h3>Price: </h3></td><td><p>{object.price === 0 ? 'free' : object.price}</p></td></tr>
									<tr><td><h3>Multiplicity: </h3></td><td><p>{object.multiplicity === 0 ? 'unlimited' : object.multiplicity}</p></td></tr>
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
			} else {
				return (
					<div className={mainClassName}>
						<div className={`${mainClassName}__container`}>
							<Button className={`${mainClassName}__close`} text="x" onClick={onClose} />
							<h2 className={`${mainClassName}__label`}>{label}</h2>
							<Input
								label="Name"
								name="name"
								placeholder="Enter name"
								type="text"
								value={title}
								onChange={(e) => {
									setTitle(e.target.value);
								}}
							/>
							<Input
								label="Description"
								name="description"
								placeholder="Enter description"
								type="textarea"
								value={summary}
								onChange={(e) => {
									setSummary(e.target.value);
								}}
							/>
							<Input
								label="Is public?"
								name="isPublic"
								placeholder=""
								type="checkbox"
								value={isPublic}
								onChange={(e) => {
									setIsPublic(!isPublic);
								}}
							/>
							<Input
								label="Is limited?"
								name="isLimited"
								placeholder=""
								type="checkbox"
								value={isLimited}
								onChange={(e) => {
									setIsLimited(!isLimited);
								}}
							/>
							{isLimited &&
								<Input
									label="Quantity"
									name="quantity"
									placeholder="quantity"
									type="number"
									value={quantity}
									onChange={(e) => {
										setQuantity(e.target.value);
									}}
								/>
							}
							<Input
								label="Price"
								name="price"
								placeholder="price"
								type="text"
								value={price}
								onChange={(e) => {
									setPrice(e.target.value);
								}}
							/>
							<p className={`${mainClassName}__error`}>{errorMsg}</p>
							<Button text="Create project" onClick={createProject} />
						</div>
					</div>);
			}
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
