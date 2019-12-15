import React, { useState } from 'react';
import { ROUTE } from './constants';
import Card from '../card/Card';
import Button from '../button/Button';
import { TYPE } from '../modal/constants';

import { api, getCookie, jwtDecode } from '../../utils';
import './content.scss';

const Content = ({ route }) => {
    const [articles, setArticles] = useState(null);
    const [projects, setProjects] = useState(null);
    const [users, setUsers] = useState(null);
    const [myProfile, setMyProfile] = useState(null);
    let mainClassName = 'content';

    const jwtToken = getCookie('jwtToken');
    const userData = jwtToken ? jwtDecode(atob(jwtToken)) : null;
    {
        !myProfile && userData &&
            api(`http://localhost:8080/user/${userData.id}`, 'GET')
                .then(response => {
                    console.log(response.data)
                    setMyProfile(response.data);
                });
    }

    switch (route) {
        case ROUTE.LANDING:
            return (
                <h1 className="lander__slogon">Find the idea you are looking for TODAY!!!</h1>
            );
        case ROUTE.ARTICLE:
            {
                !articles &&
                    api('http://localhost:8080/articles', 'GET')
                        .then(response => {
                            setArticles(response.data);
                        });
            }

            return (
                <>
                    <h2 className={`${mainClassName}__title`}>Articles</h2>
                    <ul className={`${mainClassName}__cards`}>
                        {articles && articles.map((article, key) => {
                            return (
                                <li className={`${mainClassName}__cards-li`} key={article.id}>
                                    <Card
                                        title={article.title}
                                        description={article.content}
                                        object={article}
                                        type={TYPE.ARTICLE_INFO}
                                    />
                                </li>)
                        })}
                    </ul>
                    {myProfile && myProfile.role === 1 && <Button className={`${mainClassName}__add-btn`} text='New article' />}
                </>
            );
        case ROUTE.PROJECT:
            {
                !projects &&
                    api('http://localhost:8080/projects', 'GET')
                        .then(response => {
                            setProjects(response.data);
                        });
            }

            return (
                <>
                    <h2 className={`${mainClassName}__title`}>Projects</h2>
                    <ul className={`${mainClassName}__cards`}>
                        {projects && projects.map((project, key) => {
                            return (
                                <li className={`${mainClassName}__cards-li`} key={project.id}>
                                    <Card
                                        title={project.name}
                                        description={project.description}
                                        price={project.price}
                                        multiplicity={project.multiplicity}
                                        buyers={project.buyers}
                                        owner={project.owner}
                                        object={project}
                                        type={TYPE.PROJECT_INFO}
                                    />
                                </li>)
                        })}
                    </ul>
                    {myProfile && <Button className={`${mainClassName}__add-btn`} text='New project' />}
                </>
            );
        case ROUTE.ADMIN:
            {
                !users &&
                    api('http://localhost:8080/users', 'GET')
                        .then(response => {
                            setUsers(response.data);
                        });
            }

            return (
                <>
                    <h2 className={`${mainClassName}__title`}>Users</h2>
                    <ul className={`${mainClassName}__cards`}>
                        {users && users.map((user, key) => {
                            return (
                                <li className={`${mainClassName}__cards-li`} key={user.id}>
                                    <Card
                                        isBlocked={!user.is_active}
                                        picture={user.photo_path}
                                        role={user.role}
                                        email={user.email}
                                        title={user.username}
                                        description={`Full name: ${user.full_name}`}
                                        object={user}
                                        type={TYPE.USER_INFO}
                                    />
                                </li>)
                        })}
                    </ul>
                </>
            );
        case ROUTE.PROFILE:
            return (
                <>
                    <h2 className={`${mainClassName}__title`}>My data</h2>
                    <ul className={`${mainClassName}__cards`}>
                        {myProfile &&
                            <li>
                                <Card

                                    isBlocked={!myProfile.is_active}
                                    picture={myProfile.photo_path}
                                    role={myProfile.role}
                                    email={myProfile.email}
                                    title={myProfile.username}
                                    description={`Full name: ${myProfile.full_name}`} />
                            </li>
                        }
                    </ul>
                    {myProfile && myProfile.role === 1 &&
                        <>
                            <h2 className={`${mainClassName}__title`}>My articles</h2>
                            <ul className={`${mainClassName}__cards`}>
                                {myProfile && myProfile.articles.map((article, key) => {
                                    return (
                                        <li className={`${mainClassName}__cards-li`} key={article.id}>
                                            <Card
                                                title={article.title}
                                                description={article.content}
                                            />
                                        </li>)
                                })}
                                {myProfile && myProfile.articles.length === 0 &&
                                    <li className={`${mainClassName}__cards-li`} key={-1}>
                                        <Card
                                            title={'none'}
                                        />
                                    </li>
                                }
                            </ul>
                        </>
                    }
                    <h2 className={`${mainClassName}__title`}>My projects</h2>
                    <ul className={`${mainClassName}__cards`}>
                        {myProfile && myProfile.owned_projects.map((project, key) => {
                            return (
                                <li className={`${mainClassName}__cards-li`} key={project.id}>
                                    <Card
                                        title={project.name}
                                        description={project.description}
                                        price={project.price}
                                        multiplicity={project.multiplicity}
                                        buyers={project.buyers}
                                        owner={project.owner}
                                    />
                                </li>)
                        })}
                        {myProfile && myProfile.owned_projects.length === 0 &&
                            <li className={`${mainClassName}__cards-li`} key={-1}>
                                <Card
                                    title={'none'}
                                />
                            </li>
                        }
                    </ul>
                    <h2 className={`${mainClassName}__title`}>My bought projects</h2>
                    <ul className={`${mainClassName}__cards`}>
                        {myProfile && myProfile.bought_projects.map((project, key) => {
                            return (
                                <li className={`${mainClassName}__cards-li`} key={project.id}>
                                    <Card
                                        title={project.name}
                                        description={project.description}
                                        price={project.price}
                                        multiplicity={project.multiplicity}
                                        buyers={project.buyers}
                                        owner={project.owner}
                                    />
                                </li>)
                        })}
                        {myProfile && myProfile.bought_projects.length === 0 &&
                            <li className={`${mainClassName}__cards-li`} key={-1}>
                                <Card
                                    title={'none'}
                                />
                            </li>
                        }
                    </ul>
                    <h2 className={`${mainClassName}__title`}>My followed projects</h2>
                    <ul className={`${mainClassName}__cards`}>
                        {myProfile && myProfile.followed_projects.map((project, key) => {
                            return (
                                <li className={`${mainClassName}__cards-li`} key={project.id}>
                                    <Card
                                        title={project.name}
                                        description={project.description}
                                        price={project.price}
                                        multiplicity={project.multiplicity}
                                        buyers={project.buyers}
                                        owner={project.owner}
                                    />
                                </li>)
                        })}
                        {myProfile && myProfile.followed_projects.length === 0 &&
                            <li className={`${mainClassName}__cards-li`} key={-1}>
                                <Card
                                    title={'none'}
                                />
                            </li>
                        }
                    </ul>
                </>
            );
        default:
            return <></>;
    }
};

export default Content;