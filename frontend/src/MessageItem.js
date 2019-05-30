import React from 'react';
import {Alignment, Card, Collapse, Elevation, Navbar} from "@blueprintjs/core";
import config from './config';

class MessageItem extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            expanded: false
        }
    }

    toggleCollapse = () => this.setState({expanded: !this.state.expanded});

    render() {
        return (
            <>
                <Card interactive={true} elevation={Elevation.ONE} onClick={this.toggleCollapse.bind(this)}>
                    <Navbar style={{boxShadow: "none"}}>
                        <Navbar.Group align={Alignment.LEFT}>
                            <Navbar.Heading style={{fontSize: "18px"}}><b>{this.props.message.subject}</b></Navbar.Heading>
                        </Navbar.Group>

                        <Navbar.Group align={Alignment.RIGHT}>
                            <Navbar.Heading style={{fontSize: "16px"}}>{this.props.message.from}&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</Navbar.Heading>
                        </Navbar.Group>
                    </Navbar>

                    <Collapse isOpen={this.state.expanded}>
                        <br/><br/>

                        <p style={{fontSize: "16px"}}>From: <b>{this.props.message.from}</b></p>
                        <p style={{fontSize: "14px"}}>Encrypted with: <b>{config.ALGORITHMS[this.props.message.algorithm]}</b></p>

                        <br/>

                        <p style={{fontSize: "16px"}}><b>Message:</b></p>
                        <p style={{fontSize: "14px"}}>{this.props.message.message}</p>
                    </Collapse>
                </Card>
            </>
        )
    }
}

export default MessageItem;
