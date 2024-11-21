use chatroom;
CREATE TABLE users (
  id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  account VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(512) NOT NULL,
  email VARCHAR(255),
  phone VARCHAR(128),
  avatar_url VARCHAR(1024)
);

CREATE TABLE servers (
  id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(100) NOT null
);

CREATE TABLE users_servers (
  user_id VARCHAR(255) NOT NULL,
  server_id VARCHAR(255) NOT NULL,
  PRIMARY KEY (user_id, server_id),
  Foreign Key (user_id) REFERENCES users(id) ON DELETE CASCADE,
  Foreign Key (server_id) REFERENCES servers(id) ON DELETE CASCADE
);

create TABLE channels (
  id varchar(255) PRIMARY KEY,
  name varchar(100) NOT NULL,
  type varchar(100) NOT NULL,
  server_id VARCHAR(255) NOT NULL,
  Foreign Key (server_id) REFERENCES servers(id) ON DELETE CASCADE
);

CREATE TABLE messages (
  id int AUTO_INCREMENT PRIMARY KEY,
  channel_id VARCHAR(255) NOT NULL,
  content TEXT,
  delete_flag BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (channel_id) REFERENCES channels(id) ON DELETE CASCADE
);