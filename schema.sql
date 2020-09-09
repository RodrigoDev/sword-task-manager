create schema if not exists `sword-task` collate utf8mb4_unicode_ci;

create table if not exists notifications
(
    id int auto_increment
        primary key,
    user_id int not null,
    message varchar(255) not null
);

create table if not exists tasks
(
    id int auto_increment
        primary key,
    user_id int not null,
    assigned_to int null,
    title varchar(100) not null,
    summary varchar(255) null,
    created_at datetime default CURRENT_TIMESTAMP not null,
    done_at datetime null
);

create table if not exists users
(
    id int auto_increment
        primary key,
    name varchar(100) not null,
    email varchar(100) not null,
    password varchar(100) not null,
    type smallint not null
);

