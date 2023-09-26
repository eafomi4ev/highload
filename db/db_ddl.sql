-- auto-generated definition
create database social_db
    with owner postgres;

create table cities
(
    id   integer      not null
        constraint city_pk
            primary key,
    name varchar(150) not null
);

alter table cities
    owner to postgres;

create table users
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

comment on table users is 'Таблица зарегестрированных пользователей в социальной сети';

alter table users
    owner to postgres;


