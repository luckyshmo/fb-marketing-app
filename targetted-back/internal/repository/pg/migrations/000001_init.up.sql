CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE user_tb
(
    id            uuid DEFAULT uuid_generate_v4 () not null unique,
    name          varchar(255) not null,
    email         varchar(255) not null unique,
    balance       double precision default 0.0 not null,
    phone_number  varchar(255),
    password_hash varchar(255) not null
);

CREATE TABLE ad_campaign
(
    id                          uuid DEFAULT uuid_generate_v4 () not null unique,
    user_id                     uuid references user_tb (id) on delete cascade,
    fb_page_id                  varchar(255), --not null unique,
    instagram_id                varchar(255),
    business_address            varchar(255),
    field                       varchar(255),
    name                        varchar(255) not null,
    objective                   varchar(255),
    creative_status             varchar(255),
    post_description            text,
    budget                      double precision,
	daily_budget                double precision,
    is_started                  boolean default false not null,
    time_started                timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    time_created                timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	days                        integer,
    UNIQUE (name, user_id)
);