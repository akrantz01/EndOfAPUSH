# End Of APUSH Project
This is the code for a final project from my APUSH class in school. As such, it is no longer being developed. The product is an encrypted messaging service that also has some explanation of the different types of cryptography between 1945 and present day. It has a full user account system with built in messaging. It is written in Golang and JavaScript with the frontend using React. 

## Configuration
You can use either JSON or YAML as the configuration language. Example for [YAML](./config.sample.yaml), and for [JSON](./config.sample.json). Below are the defaults for each configuration value:

| Key               | Description                              | Value     |
|-------------------|------------------------------------------|-----------|
| server.host       | What interface to have the server listen | 127.0.0.1 |
| server.port       | What port to have the server listen on   | 8080      |
| database.host     | The host where the database is running   | 127.0.0.1 |
| database.port     | The port where the database is running   | 5432      |
| database.ssl      | Use SSL for the connection               | disable   |
| database.username | Username to connect with                 | postgres  |
| database.password | Password that authenticates the user     | postgres  |
| database.database | Database to connect to                   | postgres  |
| database.wipe     | Clear the database before each run       | false     |

## API
There are 7 different distinct routes each with at least 1 REST method to accept it. Some also have query or path parameters to go with them. All routes are prefixed with the path `/api` unless otherwise specified.

#### Users
POST `/users` - create a new user<br/>
GET `/users/search` - fuzzy search for other users<br/>
GET `/users/{username}` - retrieve your user information<br/>
PUT `/users/{username}` - update your user information<br/>
DELETE `/users/{username}` - delete your own account<br/>

#### Authentication
POST `/auth/login` - generate a new authentication token for the rest of the API<br/>
GET `/auth/logout` - revoke the current authentication token, effectively logging the user out<br/>

#### Messaging
POST `/messages` - send a new message to another user<br/>
GET `/messages` - get a list of all messages sent to you<br/>
GET `/messages/{id}` - get a specific message information<br/>
