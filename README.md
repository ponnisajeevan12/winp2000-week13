## Objective: Building a Go API for returning current Toronto Time with MySQL Database Logging

The aim of this project is to build a API in Go programming language that returns the current time (America/ Toronto) in JSON format. The Go application is also integrated with mysql database in the backend and each API call inserts the current time to the table 'time_log'.

### Tasks performed: Set Up MySQL Database | API Development | Time Zone Conversion | Database Connection | Return Time in JSON | Error Handling

### Workflow:

```
1. Create docker-compose.yml & Dockerfile
2. Create init.sql & Configure .env
3. Initialize Go Project & Modules
4. Write Go Application
5. Build and Start Containers
6. Test the Application
7. Verify Database Logging
8. Stop containers
```

### Implementation Steps:

1) I have created 2 files named docker-compose.yaml & Dockerfile. The Docker Compose file ensures that the Go application can connect to MySQL and communicate with it. 
2) The docker-compose.yaml pulls the base image mysql:8.0 and runs the mysql service on port 3306. The mysql volume is mounted to the default database directory /var/lib/mysql. We have used init.sql to create our database 'time_api_db' and table 'time_log'. This file makes sure that the data is persistent everytime the mysql container is launched.
3) The docker-compose.yaml also spins up our GO container as it has reference to the Dockerfile (that contains instructions to build a go image). We need to initialize go and the below 2 go modules have also been initialized as we need to use mysql & env in our code.

```
go mod init github.com/ponnisajeevan12/winp2000-week13
go get github.com/joho/godotenv
go get -u github.com/go-sql-driver/mysql
```

4) Then, used the below command to build the images and run the containers as mentioned in our compose file.

```
$ docker-compose up --build
```
5) The images are built successfully and the containers are up and running.

![Container-launched](https://github.com/ponnisajeevan12/winp2000-week13/blob/master/images/Container-launched.png)
![container-status](https://github.com/ponnisajeevan12/winp2000-week13/blob/master/images/container-status.png)

6) While accessing the API http://localhost:8080/current-time , it returns the code 200 OK and also returns the current America/ Toronto time in the JSON format as we have mentioned in our code.

![API-timein-json](https://github.com/ponnisajeevan12/winp2000-week13/blob/master/images/API-timein-json.png)

7) Now, connect to the mysql container using the below command.

```
$ docker exec -it mysql-container mysql -u root -p<ENTER ROOT PASSWORD>
mysql> USE time_api_db;
mysql> SELECT * FROM time_log;
```

![DB-status](https://github.com/ponnisajeevan12/winp2000-week13/blob/master/images/DB-status.png)
![Table-status](https://github.com/ponnisajeevan12/winp2000-week13/blob/master/images/Table-status.png)

8) We can see that the application saves the current time (America/ Toronto) to mysql table each time we call the API. The logic for the time conversion/ format is fed into our main.go file and used Go's time package for performing this.

![time-format](https://github.com/ponnisajeevan12/winp2000-week13/blob/master/images/time-format.png)

9) As per the best practices, we have not hardcoded our DB credentials to the source code directly. Instead, we have used environment variables added this to .gitignore before pushing the code to the github.

&#10060; IMPORTANT:   Always make sure NOT to push your valid credentials to GitHub. You can avoid this by adding the .env file to .gitignore file. 

10) We also handled error in our code. For example, when the database password is mentioned wrongly in the env file, it returned the below error.

![Error-Handling](https://github.com/ponnisajeevan12/winp2000-week13/blob/master/images/Error-Handling.png)

11) We have already pushed our code to the github repository and we can use the below command to bring the containers down.

```
$ docker-compose down
```
![docker-down](https://github.com/ponnisajeevan12/winp2000-week13/blob/master/images/docker-down.png)

## Conclusion:

We have learnt how to dockerize our Go application and creating a GO API that feeds the data to the mysql database.

## References:

https://go.dev/doc/tutorial/database-access
https://medium.com/novai-devops-101/building-a-go-api-with-mysql-on-docker-using-docker-compose-3953e457bf25


















