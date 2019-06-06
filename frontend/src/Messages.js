import React from 'react';
import {Messages as ApiMessages} from './api';
import {
    Button,
    ButtonGroup,
    Card,
    Classes,
    Dialog,
    Divider,
    Elevation,
    FormGroup,
    H1,
    H3,
    InputGroup, Radio, RadioGroup, TextArea
} from "@blueprintjs/core";
import toastr from 'toastr';

import MessageItem from './MessageItem';

class Messages extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            messages: [],
            composeIsOpen: false,
            composeIsProcessing: false,
            composeData: {
                to: "",
                algorithm: 1,
                subject: "",
                message: ""
            }
        };

        this.refresh();
        this.refreshTimeout = setTimeout(() => this.refresh(), 60000);
    }

    refresh() {
        if (!this.isAuthenticated()) return;

        ApiMessages.List("in", localStorage.getItem("token")).then(res => {
            this.setState({messages: res.data});
            for (let msg in res.data) {
                this.getMessage(res.data[msg].id, msg);
            }
        }).catch(err => {
            toastr.error(err.response.data.reason + ". Please logout and log back in", "Unable to retrieve messages");
        });
    }

    getMessage(id, index) {
        ApiMessages.Read(id, localStorage.getItem("token")).then(res => {
            let newState = this.state.messages;
            newState[index] = res.data;
            this.setState(newState);
        }).catch(err => {
            toastr.error(err.response.data.reason + ". Please logout and log back in", `Unable to retrieve message ID ${id}`)
        })
    }

    componentWillUnmount() {
        clearTimeout(this.refreshTimeout);
    }

    isAuthenticated = () => localStorage.getItem("token") !== null;

    toggleComposeDialog = () => this.setState({composeIsOpen: !this.state.composeIsOpen});
    updateComposeData = (e) => this.setState({composeData: {to: (e.target.id === "compose-to") ? e.target.value : this.state.composeData.to, subject: (e.target.id === "compose-subject") ? e.target.value : this.state.composeData.subject, message: (e.target.id === "compose-message") ? e.target.value : this.state.composeData.message, algorithm: (e.target.id.substring(0, 12) === "compose-algo") ? parseInt(e.target.value) : this.state.composeData.algorithm}});

    compose() {
        this.setState({composeIsProcessing: true});
        ApiMessages.Create(this.state.composeData.to, this.state.composeData.algorithm, this.state.composeData.subject, this.state.composeData.message, localStorage.getItem("token")).then(() => {
            toastr.success(`Message sent to ${this.state.composeData.to}`);
            this.setState({composeIsProcessing: false, composeData: {to: "", algorithm: 0, subject: "", message: ""}});
            this.toggleComposeDialog();
        }).catch(err => {
            toastr.error(err.response.data.reason, "Unable to send message");
            this.setState({composeIsProcessing: false});
        });
    }

    render() {
        return (
            <>
                <Card elevation={Elevation.TWO}>
                    <H1>Messages</H1>

                    <br/><br/>

                    <p style={{fontSize: "16px"}}>
                        Below are all of the messages that have been sent to you.
                        Each one has who it is from and the subject of the message.
                        To see the full message, click on the subject line and it will open revealing the original, unencrypted message.
                    </p>
                    <p style={{fontSize: "16px"}}>Your messages will refresh every minute, but if you would like to force a refresh, click the button labeled <u>Refresh</u>.</p>
                    <p style={{fontSize: "16px"}}>To compose a new message, click the button labeled <u>Compose</u>.</p>

                    <br/><br/>

                    <ButtonGroup minimal={true}>
                        <Button icon="document-share" text="Compose" intent="success" onClick={this.toggleComposeDialog.bind(this)}/>
                        <Divider/>
                        <Button icon="refresh" text="Refresh" onClick={this.refresh.bind(this)}/>
                    </ButtonGroup>
                </Card>

                { !this.isAuthenticated() && (
                    <Card elevation={Elevation.ONE}>
                        <H3>You must be logged in to view messages</H3>
                    </Card>
                )}

                { this.isAuthenticated() && this.state.messages.map((msg, key) => <MessageItem key={key} message={msg}/>)}

                { this.isAuthenticated() && this.state.messages.length === 0 && (
                    <Card elevation={Elevation.ONE}>
                        <H3>Inbox Empty: You have no messages</H3>
                    </Card>
                )}

                <Dialog icon="document-share" title="Compose" isOpen={this.state.composeIsOpen} onClose={this.toggleComposeDialog.bind(this)}>
                    <div className={Classes.DIALOG_BODY}>
                        <FormGroup label="To" labelFor="compose-to" helperText="Username to send to">
                            <InputGroup id="compose-to" placeholder="jsmith1988" value={this.state.composeData.to} onChange={this.updateComposeData.bind(this)}/>
                        </FormGroup>
                        <FormGroup label="Subject" labelFor="compose-subject">
                            <InputGroup id="compose-subject" placeholder="Hi" value={this.state.composeData.subject} onChange={this.updateComposeData.bind(this)}/>
                        </FormGroup>
                        <FormGroup label="To" labelFor="compose-message">
                            <TextArea id="compose-message" fill={true} placeholder="Hello World" value={this.state.composeData.message} onChange={this.updateComposeData.bind(this)}/>
                        </FormGroup>
                        <FormGroup label="Encryption Type:" labelFor="compose-algo">
                            <RadioGroup id="compose-algo" onChange={this.updateComposeData.bind(this)} selectedValue={this.state.composeData.algorithm}>
                                <Radio id="compose-algo-s" label="SIGABA" value={1}/>
                                <Radio id="compose-algo-d" label="DES" value={2}/>
                                <Radio id="compose-algo-td" label="Triple DES" value={3}/>
                                <Radio id="compose-algo-a" label="AES" value={4}/>
                                <Radio id="compose-algo-dhd" label="Diffie-Hellman + DES" value={5}/>
                                <Radio id="compose-algo-dhtd" label="Diffie-Hellman + Triple DES" value={6}/>
                                <Radio id="compose-algo-dha" label="Diffie-Hellman + AES" value={7}/>
                                <Radio id="compose-algo-r" label="RSA" value={8}/>
                            </RadioGroup>
                        </FormGroup>
                    </div>
                    <div className={Classes.DIALOG_FOOTER}>
                        <div className={Classes.DIALOG_FOOTER_ACTIONS}>
                            <Button icon="cross" text="Cancel" onClick={this.toggleComposeDialog.bind(this)} disabled={this.state.composeIsProcessing}/>
                            <Button icon="saved" text="Save" intent="success" onClick={this.compose.bind(this)} disabled={this.state.composeIsProcessing}/>
                        </div>
                    </div>
                </Dialog>
            </>
        )
    }
}

export default Messages;
