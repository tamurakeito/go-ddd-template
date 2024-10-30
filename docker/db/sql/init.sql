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
