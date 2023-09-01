CREATE TABLE users
(
    id serial not null unique,
    username varchar(32) not null,
    vk_id int default null,
    vk_sync text default null,
    sync_utc_time varchar(5) default '00:00'
);