create database if not exists data_link;
use data_link;
create table if not exists dados_url (id int AUTO_INCREMENT, linkUrl varchar(1000) not null,
linkmenor varchar(500),PRIMARY KEY (id));

