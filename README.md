# A virtual classroom application in Go using the Chi router and GORM + MySQL

This application as built is intended to be used together with the complementary frontend application at [https://github.com/ikietzman/classroom-frontend](https://github.com/ikietzman/classroom-frontend), though of course you are welcome to wrangle it however you see fit.

This is very much a work in project. See below for a roadmap of next steps.

## Application install and setup
This assumes you have Go installed on your machine, and that you have a running MySQL instance to connect to. A docker compose file is included utilizing the official MySQL and phpMyAdmin images, but you are more than welcome to spin your own if you aren't comfortable with that.

Clone the repo:
```
git clone https://github.com/ikietzman/classroom-backend.git
```
Create a new database in your MySQL instance. Change the name of db/credentials_sample.go to credentials.go, un-comment the code and add in your MySQL connection credentials.

To optionally add some sample data:  
```
go run init/db_init.go
```

```
cd classroom-backend
go run .
```

The application should now be running on port 3000, though of course you'll need to install the front end (or grok the API and use curl or some other REST testing application) to use any endpoints that aren't of http request type GET.

## Roadmap for next steps
* Thoroughly comment  
* Utilize more router groups and contextualize request data where appropriate  
* Improve error handling capabilities, especially malformed requests or non-allowed DB inserts  
* Add logging to files
* Finish GORM model JSON conversion setup  
* Leverage Go's built in testing capabilities  
* Add resource not found catchall route(s)  
* Update initiation script to automatically create database
* Add ability to use static files in course lessons
* Build ML model to automate lesson selection based on previous test results
* Build, containerize, and deploy to EC2 instance
