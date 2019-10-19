
CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    users_name VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE todo_list(
    id SERIAL PRIMARY KEY,
    todo_name VARCHAR(50) NOT NULL,
    users_id INTEGER REFERENCES users(id) NOT NULL,
    UNIQUE(todo_name,users_id)     
);
CREATE TABLE todo(
    id SERIAL PRIMARY KEY,
    content VARCHAR(200) NOT NULL,
    status INTEGER NOT NULL,
    todo_list_id INTEGER REFERENCES todo_list(id) NOT NULL
);