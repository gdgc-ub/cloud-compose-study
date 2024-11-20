CREATE TABLE IF NOT EXISTS blogs (
    id serial primary key,
    title varchar(200) not null,
    image_link varchar(200),
    content text default '',
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);