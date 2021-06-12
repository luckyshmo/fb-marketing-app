CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE user_tb
(
    id            uuid DEFAULT uuid_generate_v4 () not null unique,
    name          varchar(255) not null,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE ad_company
(
    id                          uuid DEFAULT uuid_generate_v4 () not null unique,
    user_id                     uuid references user_tb (id) on delete cascade,
    fb_page_id                  varchar(255), --not null unique,
    business_address            varchar(255),
    field                       varchar(255),
    name                        varchar(255) not null,
    purpose                     varchar(255),
    creative_status             varchar(255),
    images_description          varchar(255),
    images_small_description    varchar(255),
    post_description            varchar(255),
    current_amount              integer,
	daily_amount                integer,
	days                        integer,
    UNIQUE (name, user_id)
);