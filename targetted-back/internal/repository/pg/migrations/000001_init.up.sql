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
    fb_page_id                  varchar(255) DEFAULT '',
    instagram_id                varchar(255) DEFAULT '',
    business_address            varchar(255) DEFAULT '',
    field                       varchar(255) DEFAULT '',
    name                        varchar(255) not null,
    objective                   varchar(255) DEFAULT '',
    creative_status             varchar(255) DEFAULT '',
    post_description            text DEFAULT '',
    duration                    integer DEFAULT 0,
	daily_budget                double precision DEFAULT 0.0,
    is_started                  boolean default false not null,
    time_started                timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    time_created                timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    images                      JSONB,
    UNIQUE (name, user_id)
);

INSERT INTO user_tb (id, name, email, password_hash) 
VALUES ('c0f4accb-5757-422e-9575-82251ce4ec42', 'admin@admin.com', 'admin@admin.com', '686a7172686a7177313234363137616a6668616a73d3942dce589a8baf879be01b717184712b119a72');

INSERT INTO ad_campaign (user_id, name, Images) 
VALUES ('c0f4accb-5757-422e-9575-82251ce4ec42', 'test_campaign', '{
    "images": [
        {
            "hash": "8241e4576957307c8f24e11b71456845",
            "url": "https://scontent-arn2-1.xx.fbcdn.net/v/t45.1600-4/252356672_23849292369590457_479229103746556719_n.png?_nc_cat=107&ccb=1-5&_nc_sid=2aac32&_nc_ohc=21_rXlRUblwAX9vv7L5&_nc_ht=scontent-arn2-1.xx&edm=AJNyvH4EAAAA&oh=00_AT_8V2WgQyZEY1puEf13oQOq_nQw26fy9H6tNq5aMq7z0A&oe=61BEA638"
        },
        {
            "hash": "8241e4576957307c8f24e11b71456845",
            "url": "https://scontent-arn2-1.xx.fbcdn.net/v/t45.1600-4/252356672_23849292369590457_479229103746556719_n.png?_nc_cat=107&ccb=1-5&_nc_sid=2aac32&_nc_ohc=21_rXlRUblwAX9vv7L5&_nc_ht=scontent-arn2-1.xx&edm=AJNyvH4EAAAA&oh=00_AT_8V2WgQyZEY1puEf13oQOq_nQw26fy9H6tNq5aMq7z0A&oe=61BEA638"
        }
    ]
}');