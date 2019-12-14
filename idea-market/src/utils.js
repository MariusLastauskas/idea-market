const axios = require('axios');

export const getCookie = (name) => {
    var value = "; " + document.cookie;
    var parts = value.split("; " + name + "=");
    if (parts.length == 2) return parts.pop().split(";").shift();
}

export const deleteCookie = ( name ) => {
    document.cookie = name + '=; expires=Thu, 01 Jan 1970 00:00:01 GMT;';
}

export const addCookie = (name, value) => {
    document.cookie = `${name}=${value}`;
}

export const api = (path, method) => {
    const options = {
        url: path,
        method: method,
        withCredentials: true
    };
      
    return axios(options);
}