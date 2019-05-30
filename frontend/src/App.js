import React from 'react';
import {Authentication, Users} from './api';
import {Alignment, AnchorButton, Button, Classes, Dialog, FormGroup, InputGroup, Navbar} from "@blueprintjs/core";
import { Route, BrowserRouter } from "react-router-dom";
import toastr from 'toastr';

import Home from './Home';
import Messages from './Messages';

const LEARN = () => <h2>Learn</h2>;

class App extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            loginOverlay: false,
            loginData: {
                username: "",
                password: ""
            },
            signupOverlay: false,
            signupData: {
                name: "",
                username: "",
                password: ""
            },
            processing: {
                login: false,
                signup: false
            }
        }
    }

    isAuthenticated = () => localStorage.getItem("token") !== null;

    toggleLoginOverlay = () => this.setState({loginOverlay: !this.state.loginOverlay});
    loginDataChange = (e) => this.setState({loginData: {username: (e.target.id === "login-username") ? e.target.value : this.state.loginData.username, password: (e.target.id === "login-password") ? e.target.value : this.state.loginData.password}})

    login() {
        this.setState({processing: {...this.state.processing, login: true}});
        Authentication.Login(this.state.loginData.username, this.state.loginData.password).then(res => {
            this.setState({loginData: {username: "", password: ""}, processing: {...this.state.processing, login: false}});
            localStorage.setItem("token", res.data.token);
            toastr.success("Successfully logged in");
            this.toggleLoginOverlay();
        }).catch(err => {
            console.log(err);
            toastr.error(err.response.data.reason, "Unable to login");
            this.setState({processing: {...this.state.processing, login: false}});
        })
    }

    toggleSignUpOverlay = () => this.setState({signupOverlay: !this.state.signupOverlay});
    signupDataChange = (e) => this.setState({signupData: {username: (e.target.id === "signup-username") ? e.target.value : this.state.signupData.username, password: (e.target.id === "signup-password") ? e.target.value : this.state.signupData.password, name: (e.target.id === "signup-name") ? e.target.value : this.state.signupData.name}});

    signup() {
        this.setState({processing: {...this.state.processing, signup: true}});
        Users.Create(this.state.signupData.name, this.state.signupData.username, this.state.signupData.password).then(res => {
            this.setState({signupData: {username: "", password: "", name: ""}, processing: {...this.state.processing, signup: false}});
            toastr.success("You may now login...", "Successfully signed up");
            this.toggleSignUpOverlay();
        }).catch(err => {
            toastr.error(err.response.data.reason, "Unable to sign up");
            this.setState({processing: {...this.state.processing, login: false}});
        })
    }

    logout() {
        Authentication.Logout(localStorage.getItem("token")).then(res => {
            toastr.success("Successfully logged out");
            localStorage.removeItem("token");
        }).catch(err => {
            toastr.error(err.response.data.reason, "Unable to log out");
            localStorage.removeItem("token");
        });
        setTimeout(() => this.forceUpdate(), 500);
    }

    render() {
        return (
            <>
                <Navbar>
                    <Navbar.Group align={Alignment.LEFT}>
                        <Navbar.Heading>Encrypted Messenger</Navbar.Heading>
                        <Navbar.Divider/>
                        <AnchorButton icon="home" text="Home" className={Classes.MINIMAL} href="/"/>
                        <AnchorButton icon="git-repo" text="Learn" className={Classes.MINIMAL} href="/learn"/>
                        <AnchorButton icon="comment" text="Messages" className={Classes.MINIMAL} href="/messages"/>
                    </Navbar.Group>

                    <Navbar.Group align={Alignment.RIGHT}>
                        { !this.isAuthenticated() && <Button icon="plus" text="Sign Up" intent="primary" onClick={this.toggleSignUpOverlay.bind(this)}/> }
                        { !this.isAuthenticated() && <Navbar.Divider/> }
                        { !this.isAuthenticated() && <Button icon="user" text="Login" intent="success" onClick={this.toggleLoginOverlay.bind(this)}/> }
                        { this.isAuthenticated() && <Button icon="user" text="Logout" intent="danger" onClick={this.logout.bind(this)}/> }
                    </Navbar.Group>
                </Navbar>

                <Dialog onClose={this.toggleLoginOverlay.bind(this)} isOpen={this.state.loginOverlay} title="Login" icon="user">
                    <div className={Classes.DIALOG_BODY}>
                        <FormGroup label="Username:">
                            <InputGroup id="login-username" placeholder="jsmith1988" onChange={this.loginDataChange.bind(this)} value={this.state.loginData.username}/>
                        </FormGroup>
                        <FormGroup label="Password:">
                            <InputGroup id="login-password" placeholder="secure-password123" type="password" onChange={this.loginDataChange.bind(this)} value={this.state.loginData.password}/>
                        </FormGroup>
                    </div>
                    <div className={Classes.DIALOG_FOOTER}>
                        <div className={Classes.DIALOG_FOOTER_ACTIONS}>
                            <Button onClick={this.toggleLoginOverlay.bind(this)} icon="cross" text="Cancel" disabled={this.state.processing.login}/>
                            <Button onClick={this.login.bind(this)} icon="log-in" text="Login" intent="success" disabled={this.state.processing.login}/>
                        </div>
                    </div>
                </Dialog>

                <Dialog onClose={this.toggleSignUpOverlay.bind(this)} isOpen={this.state.signupOverlay} title="Sign Up" icon="plus">
                    <div className={Classes.DIALOG_BODY}>
                        <FormGroup label="Name:">
                            <InputGroup id="signup-name" placeholder="John Smith" onChange={this.signupDataChange.bind(this)} value={this.state.signupData.name}/>
                        </FormGroup>
                        <FormGroup label="Username:">
                            <InputGroup id="signup-username" placeholder="jsmith1988" onChange={this.signupDataChange.bind(this)} value={this.state.signupData.username}/>
                        </FormGroup>
                        <FormGroup label="Password:">
                            <InputGroup id="signup-password" placeholder="secure-password123" type="password" onChange={this.signupDataChange.bind(this)} value={this.state.signupData.password}/>
                        </FormGroup>
                    </div>
                    <div className={Classes.DIALOG_FOOTER}>
                        <div className={Classes.DIALOG_FOOTER_ACTIONS}>
                            <Button onClick={this.toggleSignUpOverlay.bind(this)} icon="cross" text="Cancel" disabled={this.state.processing.signup}/>
                            <Button onClick={this.signup.bind(this)} icon="plus" text="Sign Up" intent="success" disabled={this.state.processing.signup}/>
                        </div>
                    </div>
                </Dialog>

                <BrowserRouter>
                    <Route exact path="/" component={Home}/>
                    <Route path="/learn" component={LEARN}/>
                    <Route path="/messages" component={Messages}/>
                </BrowserRouter>
            </>
        );
    }
}

export default App;
