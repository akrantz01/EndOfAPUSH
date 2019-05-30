import React  from 'react';
import {Card, Elevation, H1, H3, H5} from '@blueprintjs/core';

class Home extends React.Component {
    render() {
        return (
            <>
                <Card elevation={Elevation.TWO}>
                    <H1><u>Encrypted Messenger</u></H1>
                    <H5>By Alex Krantz</H5>

                    <br/><br/>

                    <p style={{fontSize: "16px"}}>This is a showcase of 8 different encryption mechanisms that were invented between 1945 and 2019 through a simple messaging service.
                        With this, you are able to send and receive encrypted messages that no one else can read, not even the server.
                        Every message that you send is scrambled based on the algorithm that you choose and is then sent off to the server and copied to the recipient's inbox for them to read.
                        All of the encryption and decryption happens auto-magically in the background so you don't have to deal with it yourself.
                    </p>

                    <p style={{fontSize: "16px"}}>Below are all of the encryption mechanisms that are implemented:</p>
                    <ul style={{fontSize: "16px", marginTop: "-6px"}}>
                        <li>SIGABA</li>
                        <li>Data Encryption Standard (DES)</li>
                        <li>Advanced Encryption Standard (AES)</li>
                        <li>Diffie-Hellman Key Exchange (Diffie-Hellman)</li>
                        <li>Rivest-Shamir-Adleman Cryptosystem (RSA)</li>
                        <li>Transport Layer Security (TLS)</li>
                    </ul>
                    <p style={{fontSize: "16px"}}>To learn more about any of those mentioned above, click the <u>Learn</u> button in the navigation bar.</p>

                    <p style={{fontSize: "16px"}}>The encryption algorithms that you can use on each method are:</p>
                    <ul style={{fontSize: "16px", marginTop: "-6px"}}>
                        <li>SIGABA</li>
                        <li>DES</li>
                        <li>Triple DES</li>
                        <li>AES</li>
                        <li>Diffie-Hellman + DES</li>
                        <li>Diffie-Hellman + Triple DES</li>
                        <li>Diffie-Hellman + AES</li>
                        <li>RSA</li>
                    </ul>
                    <p style={{fontSize: "16px"}}>If you are wondering why TLS is not on that list, or why Diffie-Hellman is combined with some another encryption algorithm, or why there is a Triple DES, click the <u>Learn</u> button in the navigation bar.</p>
                </Card>
                <Card elevation={Elevation.ONE}>
                    <H3><u>The Code</u></H3>

                    <br/><br/>

                    <p style={{fontSize: "16px"}}>If for whatever reason you are curious as to how this was made, click <a href="https://github.com/akrantz01/EndOfAPUSH" target="_blank">here</a> to view the GitHub repository.</p>
                </Card>
            </>
        )
    }
}

export default Home;
