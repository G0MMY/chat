
CREATE TABLE IF NOT EXISTS users(
   id serial PRIMARY KEY,
   username VARCHAR (50) UNIQUE NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL,
   password CHAR (64) NOT NULL
);

CREATE TABLE IF NOT EXISTS rooms(
   id serial PRIMARY KEY,
   name VARCHAR(250) NOT NULL
);

CREATE TABLE IF NOT EXISTS rooms_users(
   room_id int NOT NULL,
   user_id int NOT NULL,
   CONSTRAINT fk_room FOREIGN KEY(room_id) REFERENCES rooms(id),
   CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
   UNIQUE (room_id, user_id)
);

CREATE TABLE IF NOT EXISTS invitations(
   id serial PRIMARY KEY,
   sender VARCHAR (50) NOT NULL,
   receiver VARCHAR (50) NOT NULL,
   room_id int NOT NULL,
   CONSTRAINT fk_room FOREIGN KEY(room_id) REFERENCES rooms(id),
   CONSTRAINT fk_sender FOREIGN KEY(sender) REFERENCES users(username),
   CONSTRAINT fk_receiver FOREIGN KEY(receiver) REFERENCES users(username)
);

CREATE TABLE IF NOT EXISTS messages(
   id serial PRIMARY KEY,
   room_id int NOT NULL,
   sender VARCHAR (50) NOT NULL,
   send_time timestamp NOT NULL,
   msg TEXT NOT NULL,
   CONSTRAINT fk_room FOREIGN KEY(room_id) REFERENCES rooms(id),
   CONSTRAINT fk_username FOREIGN KEY(sender) REFERENCES users(username)
);