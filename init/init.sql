DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id integer AUTO_INCREMENT,
  name varchar(255),
  mail varchar(255),
  password varchar(255),
  primary key(id)
);

INSERT INTO users (name, mail, password)
VALUES
  ('po3rin', 'test1@mail.com', 'pass1'),
  ('taro', 'test2@mail.com', 'pass2');