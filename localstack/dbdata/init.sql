CREATE DATABASE serverless_example;
--  Replace 'test-password' with a secure password of your choice
CREATE USER srv_serverless WITH PASSWORD 'test-password';
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO srv_serverless;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO srv_serverless;

create table authors
(
	id serial not null primary key ,
	first_name text not null,
	last_name text not null,
	created_at timestamp default now() not null,
	updated_at timestamp default now() not null
);

create table articles
(
    id serial not null primary key ,
    name text not null,
    author_id int not null references authors,
    created_at timestamp default now() not null,
    updated_at timestamp default now() not null
);