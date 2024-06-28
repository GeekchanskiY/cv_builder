create table if not exists employees (
    id serial primary key,
    name varchar(255) not null unique
);