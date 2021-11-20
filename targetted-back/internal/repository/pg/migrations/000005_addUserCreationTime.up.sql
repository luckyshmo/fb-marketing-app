ALTER TABLE user_tb ADD column if not exists time_created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;

UPDATE user_tb SET  time_created = date_created;

alter table user_tb alter column date_created drop default;
alter table user_tb alter column date_created drop not null;