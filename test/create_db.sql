create database learn_test;

create table learn_test.category (
    id int AUTO_INCREMENT PRIMARY KEY not null,
    name varchar(200) not null
) engine = innodb;

select * from learn_test.category