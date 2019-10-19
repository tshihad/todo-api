CREATE TABLE todo_list(
    id SERIAL PRIMARY KEY,
    todo_name VARCHAR(50) NOT NULL,
    users_id VARCHAR(50) NOT NULL,
    UNIQUE(todo_name,users_id)     
);
CREATE TABLE todo(
    id SERIAL PRIMARY KEY,
    content VARCHAR(200),
    status INTEGER,
    todo_list_id INTEGER REFERENCES todo_list(id)
);
CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    users_name VARCHAR(50) UNIQUE,
    email VARCHAR(50) UNIQUE
);