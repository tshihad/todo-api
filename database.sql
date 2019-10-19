CREATE TABLE todo_list(
    id INTEGER PRIMARY KEY,
    todo_name VARCHAR(50) NOT NULL,
    users_id VARCHAR(50) NOT NULL,
    UNIQUE(todo_name,users_id)     
);
CREATE TABLE todo(
    id INTEGER PRIMARY KEY,
    content VARCHAR(200),
    status INTEGER,
    todo_list_id INTEGER REFERENCES todo_list(id)
);
CREATE TABLE users(
    id INTEGER PRIMARY KEY,
    users_name VARCHAR(50),
    email VARCHAR(50)
);