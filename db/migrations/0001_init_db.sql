-- +goose Up
create table if not exists cities
(
    id   serial primary key,
    name varchar(150) not null
);

alter table if exists cities
    owner to postgres;

INSERT INTO cities(name)
VALUES ('Paris'),
       ('New York'),
       ('Birmingham'),
       ('Tokio');

create table if not exists users
(
    id            uuid        not null
        constraint user_pk
            primary key,
    first_name    varchar(50) not null,
    surname       varchar(50) not null,
    birthdate     date        not null,
    biography     text,
    password_hash text        not null,
    city_id       integer
        constraint user_city___fk
            references cities
            on delete restrict
);

comment on table users is 'Таблица зарегистрированных пользователей в социальной сети';

alter table if exists users
    owner to postgres;


