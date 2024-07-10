CREATE table users (
  user_id varchar(10) not null PRIMARY key,
  avatar text,
  full_name text,
  birthday time,
  email text,
  phone text,
  address text,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP on UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE boards (
  board_id varchar(10) not null PRIMARY key,
  title text,
  description text,
  type text,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP on UPDATE CURRENT_TIMESTAMP
);

CREATE table board_users (
  board_id varchar(10) not null,
  user_id varchar(10) not null,
  role  varchar(20) DEFAULT "member",
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP on UPDATE CURRENT_TIMESTAMP,
  PRIMARY key(board_id, user_id),
  FOREIGN key (board_id) REFERENCES boards(board_id),
  FOREIGN key (user_id) REFERENCES users(user_id)
);
CREATE TABLE columns (
  board_id varchar(10) not null,
  column_id varchar(10) not null,
  column_order int,
  title text,
  PRIMARY KEY (board_id, column_id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP on UPDATE CURRENT_TIMESTAMP,
  FOREIGN key (board_id) REFERENCES boards(board_id)
);
CREATE TABLE cards (
  board_id varchar(10) not null,
  column_id varchar(10) not null,
  card_id varchar(10) not null,
  card_order int,
  title text,
  description text,
  thumbnail text,
  PRIMARY KEY (board_id, column_id, card_id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP on UPDATE CURRENT_TIMESTAMP,
  FOREIGN key (board_id, column_id) REFERENCES columns(board_id, column_id)

);

create table card_members (
  board_id varchar(10) not null,
  column_id varchar(10) not null,
  card_id varchar(10) not null,
  user_id varchar(10) not null,
  PRIMARY KEY (board_id, column_id, card_id, user_id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP on UPDATE CURRENT_TIMESTAMP,
  FOREIGN key (board_id, column_id, card_id) REFERENCES cards(board_id, column_id, card_id),
  FOREIGN key (user_id) REFERENCES users(user_id)
);

create table card_comments (
  board_id varchar(10) not null,
  column_id varchar(10) not null,
  card_id varchar(10) not null,
  user_id varchar(10) not null,
  content text,
  image text,
  PRIMARY KEY (board_id, column_id, card_id, user_id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP on UPDATE CURRENT_TIMESTAMP,
  FOREIGN key (board_id, column_id, card_id) REFERENCES cards(board_id, column_id, card_id),
  FOREIGN key (user_id) REFERENCES users(user_id)
);

create table card_attachments (
  board_id varchar(10) not null,
  column_id varchar(10) not null,
  card_id varchar(10) not null,
  user_id varchar(10) not null,
  url_file text,
  PRIMARY KEY (board_id, column_id, card_id, user_id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP on UPDATE CURRENT_TIMESTAMP,
  FOREIGN key (board_id, column_id, card_id) REFERENCES cards(board_id, column_id, card_id),
  FOREIGN key (user_id) REFERENCES users(user_id)
);


DROP TABLE `trello-db`.`card_attachments`;

DROP TABLE `trello-db`.`card_comments`;

DROP TABLE `trello-db`.`card_members`;

DROP TABLE `trello-db`.`board_users`;

DROP TABLE `trello-db`.`cards`;

DROP TABLE `trello-db`.`columns`;

DROP TABLE `trello-db`.`boards`;

DROP TABLE `trello-db`.`users`;