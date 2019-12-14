import React, { useState } from 'react';
import Image from '../image/Image';
import './card.scss';

const Card = ({ title, image, description, price, multiplicity, buyers, owner, picture }) => {
	const mainClassName = 'card';

	return (
		<div className={`${mainClassName}`}>
			<div className={`${mainClassName}__primary`}>
				<h2 className={`${mainClassName}__title`}>{title}</h2>
				<p className={`${mainClassName}__description`}>{description}</p>
				<Image className={`${mainClassName}__image`} src={image} />
			</div>
			{price !== undefined && (
				<div className={`${mainClassName}__secondary`}>
					{multiplicity !== undefined && (
						<span className={`${mainClassName}__multiplicity`}>
							{multiplicity > 0 ? `${multiplicity - buyers.length} units left` : 'unlimited'}
						</span>
					)}
					<span className={`${mainClassName}__price`}>{price > 0 ? `${price}€` : 'free'}</span>
					{owner && <span className={`${mainClassName}__owner`}>Created by {owner.username}</span>}
				</div>
			)}
		</div>
	);
};

export default Card;
