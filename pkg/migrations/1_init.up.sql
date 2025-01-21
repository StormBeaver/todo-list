CREATE TABLE IF NOT EXISTS users
(
    users_id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    username VARCHAR(20) NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS todo_lists
(
    todolists_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT
);


CREATE TABLE IF NOT EXISTS todo_items
(
    todoitems_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    done BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS users_lists
(
    userslists_id SERIAL PRIMARY KEY,
    user_id int,
    list_id int,
    FOREIGN KEY (user_id) REFERENCES users(users_id) ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES todo_lists(todolists_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS items_lists
(
    itemslists_id SERIAL PRIMARY KEY,
    item_id INT,
    list_id INT,
    FOREIGN KEY (item_id) REFERENCES todo_items(todoitems_id) ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES todo_lists(todolists_id) ON DELETE CASCADE
);