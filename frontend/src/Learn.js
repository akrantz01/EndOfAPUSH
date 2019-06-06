import React from 'react';
import {Card, Elevation, H1, H3} from "@blueprintjs/core";

class Learn extends React.Component {
    render() {
        return (
            <>
                <Card elevation={Elevation.TWO}>
                    <H1>Learn</H1>

                    <br/><br/>

                    <p style={{fontSize: "16px"}}>
                        Below are descriptions are all of the primary cryptography algorithms invented between 1945 and 2019.
                        Some even look into the future by accounting for the impact of quantum computers.
                    </p>
                </Card>
                <Card elevation={Elevation.ONE}>
                    <H3>SIGABA</H3>

                    <br/><br/>

                    <p style={{fontSize: "16px"}}>
                        SIGABA was the United States' response to the German Enigma machine.
                        It was invented by William A. Friedman and Frank Rowlett prior to World War II.
                        This was the United States' first journey into cryptography and it was very successful.
                        Throughout the entirety of WWII, it was never broken by Axis powers, and was even called "unbreakable" by the German cryptanalysis lab.
                        This was not off base as it had an effective search space of 2<sup>902</sup>, which is larger than many cryptographic algorithms today.
                        The machine was entirely electromechanical and developed entirely by the US army and navy.
                        It consisted of 5 cipher rotors, 5 control rotors and 5 index rotors which all worked together to make the messages unreadable.
                    </p>
                </Card>
                <Card elevation={Elevation.ONE}>
                    <H3>Data Encryption Standard</H3>

                    <br/><br/>

                    <p style={{fontSize: "16px"}}>
                        The United States Data Encryption Standard (DES) was a cryptographic algorithm that was used once computers became more common place.
                        It was developed entirely within the National Security Agency leading to suspicion of its validity as a cipher.
                        The original version of it was name Lucifer and invented by researcher Horst Feistel of IBM.
                        It was mandatory for all US government financial transactions, despite not being very secure.
                        In order to increase security, the US government recommended running the algorithm 3 times over a single payload, giving way to the name Triple DES.
                        As of 1999, DES was deemed obsolete due to its small search space of only 2<sup>56</sup> bits and the increasing power of computers which made cracking it considerably easier.
                    </p>
                </Card>
                <Card elevation={Elevation.ONE}>
                    <H3>Advanced Encryption Standard</H3>

                    <br/><br/>

                    <p style={{fontSize: "16px"}}>
                        The Advanced Encryption Standard was introduced as a replacement to DES in 2001, despite work beginning in 1997.
                        Rather than being done entirely in house, the US government consulted industry leaders and evaluated cryptographic algorithms in 3 rounds.
                        After the three rounds, the algorithm called Rijndael was chosen as the definition of AES.
                        It supports 128-, 192-, and 256-bit encryption, considerably more secure than the 56-bit encryption of DES.
                        AES is still used today, and as a requirement, remains entirely open source and free to use.
                    </p>
                </Card>
                <Card elevation={Elevation.ONE}>
                    <H3>Diffie-Hellman Key Exchange</H3>

                    <br/><br/>

                    <p style={{fontSize: "16px"}}>
                        The Diffie-Hellman Key Exchange was the first introduction into public key cryptography which was previously thought impossible.
                        Initially published in November of 1976, Whitfield Diffie and Martin Hellman thought they had proven that public key cryptography was possible.
                        Public key cryptography is the use of two keys, one kept private that is used for decryption, and the other made public that is used for encryption.
                        It wasn't until the two began working with Ralph Merkle of UC Berkeley when the Diffie-Hellman key exchange began to become a reality.
                        The Diffie-Hellman Key Exchange works by using modular arithmetic to establish a shared secret between two parties.
                        The exchange itself is unable to encrypt or decrypt which is why it is required to be paired with a second, symmetric algorithm.
                        The Diffie-Hellman key exchange is used everywhere from the web to VPNs to provide secure encrpytion.
                    </p>
                </Card>
                <Card elevation={Elevation.ONE}>
                    <H3>Rivest–Shamir–Adleman Cryptosystem</H3>

                    <br/><br/>

                    <p style={{fontSize: "16px"}}>
                        The Rivest-Shamir-Adleman Cryptosystem (RSA) is the first public key cryptosystem that can encrypt and decrypt without help from another algorithm like the Diffie-Hellman Key Exchange requires.
                        It was inspired by the Diffie-Hellman key exchange and improved on the deficiencies of it, namely the fact there was no one-way function.
                        RSA improved on the Diffie-Hellman key exchange by applying number theory and Fermat's Little Theorem.
                        The foundation of RSA is the difficulty to solve the discrete logarithm problem.
                        During the time it was invented, the government was the only section that used cryptography, but RSA made it far more accessible.
                        RSA is one of the most widely used cryptosystem as it is one of the foundations of the internet.
                        Despite being invented in 1977, RSA is still secure and has been made more secure by the application of elliptic curve cryptography.
                    </p>
                </Card>
                <Card elevation={Elevation.ONE}>
                    <H3>Transport Layer Security</H3>

                    <br/><br/>

                    <p style={{fontSize: "16px"}}>
                        Transport Layer Security (TLS) is modern foundation of the secure internet.
                        TLS replaces the widely outdated Secure Socket Layer invented by the Netscape corporation.
                        It began as SSL 3.1, but due to the dramatic differences, it got a new name.
                        TLS works on top of the base networking protocols (TCP and UDP) and adds a new layer of encryption to prevent interception.
                        There are currently four versions - 1.0, 1.1, 1.2, and 1.3 - with the latest being published by the Internet Engineering Task Force in 2018.
                        TLS is the reason that you see the lock icon by a web address.
                        It works based on encryption of data in transit through the Diffie-Hellman key exchange, authentication that the server and client are who they say they are through certificates, and verification of the integrity of data through message authentication codes.
                        TLS cannot be used for general data encryption and is why TLS is not an option to encrypt data with this messenger.
                        It is, however, how this website is served to you (notice the green lock icon if you are using Chrome).
                    </p>
                </Card>
                <Card elevation={Elevation.ONE}>
                    <H3>Post-Quantum Cryptography</H3>

                    <br/><br/>

                    <p style={{fontSize: "16px"}}>
                        With the rise of quantum computers, our entire basis of cryptography is on the verge of being broken.
                        Quantum computers use the principles of quantum physics in order to speed up computation times.
                        Rather than classical computers which have two states, 1 and 0, quantum computers can have states between the two.
                        This is called quantum superposition.
                        When running a program on a quantum computer, the program will generally be sped up in square root time.
                        This allows for quantum computers to solve currently computationally intensive problems, such as the discrete logarithm problem.
                        Through the use of Shor's algorithm, quantum computers are able to break the discrete logarithm problem, and by extension, RSA and the Diffie-Hellman key exchange.
                        Symmetric algorithms are not safe either as with sufficient qubits, a quantum computer can break any cryptographic algorithm.
                        As such, work is currently being done into post-quantum cryptography.
                        These algorithms are intended to be unbreakable by quantum computers.
                        The currently proposed algorithms that are being tested all use different styles, such as lattice-based cryptography, learning with errors, and many others.
                        The proposed algorithms are <a href="https://www.microsoft.com/en-us/research/project/frodokem/" target="_blank" rel="noopener noreferrer">FrodoKEM</a>, <a href="https://www.microsoft.com/en-us/research/project/sike/" target="_blank" rel="noopener noreferrer">SIKE</a>, <a href="https://www.microsoft.com/en-us/research/project/picnic/" target="_blank" rel="noopener noreferrer">Picnic</a>, and <a href="https://www.microsoft.com/en-us/research/project/qtesla/" target="_blank" rel="noopener noreferrer">qTESLA</a>.
                    </p>
                </Card>
            </>
        )
    }
}

export default Learn;
