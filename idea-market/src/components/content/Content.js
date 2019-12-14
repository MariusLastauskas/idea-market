import React, {useState} from 'react';
import {ROUTE} from './constants';
import Card from '../card/Card';

import {api} from '../../utils';
import './content.scss';

const Content = ({route}) => {
    const [articles, setArticles] = useState(null);
    const [projects, setProjects] = useState(null);
    const [users, setUsers] = useState(null);
    let mainClassName = 'content';

    switch (route) {
    case ROUTE.LANDING:
        return (
            <h1 className="lander__slogon">Find the idea you are looking for TODAY!!!</h1>
        );
    case ROUTE.ARTICLE:
        {!articles && 
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
                                />
                            </li>)
                        })}
                </ul>
            </>
        );
    case ROUTE.PROJECT:
        {!projects && 
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
                                />
                            </li>)
                        })}
                </ul>
            </>
        );
    case ROUTE.ADMIN:
        {!users && 
            api('http://localhost:8080/users', 'GET')
            .then(response => {
                console.log(response)
                setUsers(response.data);
            });
        }

        return (
            <>
                <h2 className={`${mainClassName}__title`}>Users</h2>
                <ul className={`${mainClassName}__cards`}>
                    {users && users.map((user, key) => {
                        console.log(user)
                        return (
                            <li className={`${mainClassName}__cards-li`} key={user.id}>
                                <Card 
                                    picture={user.photo_path}
                                    role={user.role}
                                    email={user.email}
                                    title={user.username} 
                                    description={`Full name: ${user.full_name}`} 
                                />
                            </li>)
                        })}
                </ul>
            </>
        );
    default:
        return <></>;
    }
};

export default Content;