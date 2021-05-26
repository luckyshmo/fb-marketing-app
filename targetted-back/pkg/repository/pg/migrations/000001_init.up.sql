CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE user_tb
(
    id            uuid DEFAULT uuid_generate_v4 () not null unique,
    name          varchar(255) not null,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null
);