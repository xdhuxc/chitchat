create database chitchat if not exists;

use chitchat;

create table if not exits `users`(
  id int(11) auto_increment,
  uuid varchar(64) not null unique,
  name varchar(100) not null,
)