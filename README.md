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
    REDIS_HOST="<Redis Host>"
    REDIS_PORT=<Redis port number>
    REDIS_PASSWORD="<Redis Password>"
    SONG_CACHE_DB=<Database number for Song Caching>    
    ACCOUNT_CACHE_DB=<Database number for Account Caching>
    SONG_CACHE_EXPIRE=<Expiration time in seconds>
    ACCOUNT_CACHE_EXPIRE=<Expiration time in seconds>
    ```
3. Run the program using ```go run main.go```
4. To test the api use ``` go test ./test```

## Usage
API Documentation: `https://documenter.getpostman.com/view/27497116/2s93sf1W5R`