export default {
    // Location of the backend API
    API_URL: "http://localhost:8080/api",

    // All of the supported algorithms with their corresponding IDs
    ALGORITHMS: Object.freeze({
        1: "SIGABA",
        2: "DES",
        3: "Triple DES",
        4: "AES",
        5: "Diffie-Hellman + DES",
        6: "Diffie-Hellman + Triple DES",
        7: "Diffie-Hellman + AES",
        8: "RSA"
    })
}
