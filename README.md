# Simple TODO
This is a simple todo application using postgres as database

## Architecture
Here I am using clean Architecture consists of 3 layers.
Listing layers with implementation
- entities ==> app/models
- repository ==> app/data
- usecases/api ==> app/api

project structure references: https://github.com/golang-standards/project-layout

## Setting up
Run following command will runs both database and app
```
docker-compose up
```
Postman collection can be found in docs directory
### Local setup
#### Unit test
```
go test ./...
```
#### Run App
```
export DB_HOST=localhost
go run main.go
```

## Documentation
### DB
Tables
* users (id(pk), name, email)
* todolist (id(pk), name, user_id(fk))
* todo (id(pk), content, status, todo_list_id(fk))

### API
* GET - /app/user/{id} ==> to get-user
* POST - /app/user ==> new user
* GET - /app/todolist/{user_id} ==> todo lists under a user
* POST - /app/todolist/{user_id} ==> new todo list for a user
* PUT - /app/todolist/{id} ==> update todolist
* DELETE - /app/todolist/{id} ==> delete todo list
* GET - /app/todo/{todo_list_id} ==> get todos under a todo list with filer option
* POST - /app/todo/{todo_list_id} ==> new todo on a todo list
* PUT - /app/todo/{id} ==> update todo with/without status
* DELETE -/app/todo/{id} ==> delete todo
