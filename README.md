Song API is a REST API that allows you to manage songs in a playlist corresponding to the user. It is built in Golang with Gin and Gorm library that uses a MySQL database to store the data. 

## Installation
1. Create new database in MySQL ```create database <database-name>```
2. Create a ```.env``` file with the following contents and keep it in the same directory as main.go
    ```
    HOST="<Your hostname>"
    PORT=<Port Number>
    USER="<Database Username>"
    PASSWORD="<Database Password>"
    DB_NAME="<Name of the Database>"
    AUTH_KEY="Authorization Secret Key for generating bearer token"
    ```
3. Run the program using ```go run main.go```
4. To test the api use ``` go test ./test```

## Usage
1. POST request on ```/v1/accounts/new``` creates a new account.
  - Request Body
    ```
    {
        "user" : "MrStark",
        "password" : "l0v3_y0u_3000"
    }
    ```
  - Response Body
    ```
    {
        "account": {
            "id": 3,
            "user": "MrStark",
            "password": "l0v3_y0u_3000"
        },
        "response": "Account created successfully",
        "status": 200
    }
    ```
    
2. POST request on ```/v1/accounts/``` returns account details and auth token.
  - Request Body
    ```
    {
        "user" : "MrStark",
        "password" : "l0v3_y0u_3000"
    }
    ```
  - Response Body
    ```
    {
        "account": {
            "id": 3,
            "user": "MrStark",
            "password": "l0v3_y0u_3000"
        },
        "status": 200,
        "token": "<A very long Authentication token>"
    }
    ```
3. POST request on ```/v1/songs/``` adds a new song to the database. Auth token is required to insert song in the playlist.
  - Request Body
    ```
    {
        "song": "Sunflower",
        "artist" : "Post Malone",
        "plays" : 100,
        "release_date" : "2018-10-18"
    }
    ```
  - Response Body
    ```
    {
        "data": {
            "id": 4,
            "song": "Sunflower",
            "artist": "Post Malone",
            "plays": 100,
            "release_date": "2018-10-18",
            "user": "MrStark"
        },
        "response": "Song added successfully",
        "status": 200
    }
    ```
4. GET request on ```/v1/songs/``` returns all the songs in the playlist that belongs to the user whose token is passed in the Authorization header.
  - Response
    ```
    {
        "data": [
            {
                "id": 4,
                "song": "Sunflower",
                "artist": "Post Malone",
                "plays": 100,
                "release_date": "2018-10-18",
                "user": "MrStark"
            }
        ],
        "status": 200
    }
    ```
5. PUT request on ```/v1/songs/:id``` updates the song that has the given id. Auth token is required. Example url is /v1/songs/4
  - Request Body
    ```
    {
        "plays": 120
    }
    ```
  - Response Body
    ```
    {
        "data": {
            "id": 4,
            "song": "Sunflower",
            "artist": "Post Malone",
            "plays": 120,
            "release_date": "2018-10-18",
            "user": "MrStark"
        },
        "response": "Song updated successfully",
        "status": 200
    }
    ```
6. GET request on ```/v1/songs/:id``` returns the song that has the given id. Auth token is required. Example url is /v1/songs/4
  - Response Body
    ```
    {
        "data": {
            "id": 4,
            "song": "Sunflower",
            "artist": "Post Malone",
            "plays": 120,
            "release_date": "2018-10-18",
            "user": "MrStark"
        },
        "status": 200
    }
    ```
7. DELETE request on ```/v1/songs/:id``` deletes the song that has the given id. Auth token is required. Example url is /v1/songs/4
  - Response Body
    ```
    {
        "response": "id 4 deleted by MrStark",
        "status": 200
    }
    ```