create database if not exists data_link;
use DataAPiLink;
create table if not exists Dados_URL (id int AUTO_INCREMENT, linkUrl varchar(1000) not null,
linkmenor varchar(500),PRIMARY KEY (id));

