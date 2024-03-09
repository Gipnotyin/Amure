create table if not exists "user"
(
    user_id       uuid DEFAULT gen_random_uuid() primary key,
    login         varchar(255) unique not null,
    name          varchar(255)        not null,
    last_name     varchar(255)        not null,
    email         varchar(255) unique not null,
    phone         varchar(255) unique not null,
    hash_password varchar(255)        not null
)