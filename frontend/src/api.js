import axios from 'axios';
import config from './config';

export class Users {
    static Create(name, username, password) {
        return new Promise(((resolve, reject) =>{
            axios({
                method: "POST",
                url: `${config.API_URL}/users`,
                headers: {"Content-Type": "application/json"},
                data: {
                    name: name,
                    username: username,
                    password: password
                }
            }).then(res => resolve(res.data)).catch(err => reject(err));
        }));
    }

    static Delete(username, token) {
        return new Promise(((resolve, reject) =>{
            axios({
                method: "DELETE",
                url: `${config.API_URL}/users/${username}`,
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": token
                },
            }).then(res => resolve(res.data)).catch(err => reject(err));
        }));
    }

    static Update(newName, newUsername, newPassword, username, token) {
        return new Promise(((resolve, reject) =>{
            axios({
                method: "PUT",
                url: `${config.API_URL}/users/${username}`,
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": token
                },
                data: {
                    name: newName,
                    username: newUsername,
                    password: newPassword
                }
            }).then(res => resolve(res.data)).catch(err => reject(err));
        }));
    }

    static Read(username, token) {
        return new Promise(((resolve, reject) =>{
            axios({
                method: "GET",
                url: `${config.API_URL}/users/${username}`,
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": token
                },
            }).then(res => resolve(res.data)).catch(err => reject(err));
        }));
    }

    static Search(username, token) {
        return new Promise(((resolve, reject) =>{
            axios({
                method: "GET",
                url: `${config.API_URL}/users/search`,
                headers: {"Authorization": token},
                params: {username: username}
            }).then(res => resolve(res.data)).catch(err => reject(err));
        }));
    }
}

export class Authentication {
    static Login(username, password) {
        return new Promise(((resolve, reject) =>{
            axios({
                method: "POST",
                url: `${config.API_URL}/auth/login`,
                headers: {"Content-Type": "application/json"},
                data: {
                    username: username,
                    password: password
                }
            }).then(res => resolve(res.data)).catch(err => reject(err));
        }));
    }

    static Logout(token) {
        return new Promise(((resolve, reject) =>{
            axios({
                method: "GET",
                url: `${config.API_URL}/auth/logout`,
                headers: {"Authorization": token},
            }).then(res => resolve(res.data)).catch(err => reject(err));
        }));
    }
}

export class Messages {
    static Create(to, algorithm, subject, message, token) {
        return new Promise(((resolve, reject) =>{
            axios({
                method: "POST",
                url: `${config.API_URL}/messages`,
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": token
                },
                data: {
                    to: to,
                    algorithm: algorithm,
                    subject: subject,
                    message: message
                }
            }).then(res => resolve(res.data)).catch(err => reject(err));
        }));
    }

    static Read(id, token) {
        return new Promise(((resolve, reject) =>{
            axios({
                method: "GET",
                url: `${config.API_URL}/messages/${id}`,
                headers: {"Authorization": token},
            }).then(res => resolve(res.data)).catch(err => reject(err));
        }));
    }

    static List(type, token) {
        return new Promise(((resolve, reject) =>{
            axios({
                method: "GET",
                url: `${config.API_URL}/messages`,
                headers: {"Authorization": token},
                params: {type: type}
            }).then(res => resolve(res.data)).catch(err => reject(err));
        }));
    }
}
