-- migrate:up
create table users (
  id int not null primary key auto_increment,
  username varchar(255),
  password_hash text
);

-- migrate:down
drop table users;