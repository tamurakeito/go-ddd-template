USE go-ddd-template_app;

CREATE TABLE hello_world(
	id INT(11) AUTO_INCREMENT NOT NULL, 
  name VARCHAR(30) NOT NULL,
  tag BOOLEAN NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO hello_world(name, tag) VALUES
  ('hello, world!', true),
  ('こんにちは！', false),
  ('안녕하세요!', false);

CREATE TABLE accounts(
	id INT(11) AUTO_INCREMENT NOT NULL,
  user_id VARCHAR(30) NOT NULL,
  password VARCHAR(60) NOT NULL,
  name VARCHAR(30) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (user_id) 
);


INSERT INTO accounts(user_id, password, name) VALUES
  ('user', '$2a$10$7gGCP2gWlwUXF/gqUzSK1uUk1hwze/peeRnT.enW.TLvtHUbaFVQm', 'テストユーザー');