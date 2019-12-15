import React, { useState } from 'react';
import Image from '../image/Image';
import './card.scss';

const Card = ({ title, image, description, price, multiplicity, buyers, owner, picture, role, email, isBlocked }) => {
	const mainClassName = 'card';

	return (
		<div
			className={`${mainClassName} ${role !== undefined
				? role === 1 ? mainClassName + '--admin' : isBlocked ? mainClassName + '--blocked' : ''
				: ''}`}
		>
			<div className={`${mainClassName}__primary`}>
				<h2 className={`${mainClassName}__title`}>{title}</h2>
				<p className={`${mainClassName}__description`}>{description}</p>
				<Image className={`${mainClassName}__image`} src={image} />
			</div>
			{price !== undefined && (
				<div className={`${mainClassName}__secondary`}>
					{multiplicity !== undefined && (
						<span className={`${mainClassName}__multiplicity`}>
							{multiplicity > 0 ? `${buyers ? multiplicity - buyers.length : multiplicity} units left` : 'unlimited'}
						</span>
					)}
					<span className={`${mainClassName}__price`}>{price > 0 ? `${price}€` : 'free'}</span>
					{owner && <span className={`${mainClassName}__owner`}>Created by {owner.username}</span>}
				</div>
			)}
			{email && (
				<div className={`${mainClassName}__secondary`}>
					<Image className={`${mainClassName}__avatar`} src={picture} />
					<span className={`${mainClassName}__role`}>{role === 1 ? 'admin' : 'user'}</span>
					<span className={`${mainClassName}__email`}>{email}</span>
				</div>
			)}
		</div>
	);
};

export default Card;
