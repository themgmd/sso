-- +migrate Up

create table if not exists users (
  id uuid primary key not null,
  family_name text not null,
  given_Name text not null,
  patronymic text default null,
  email text not null,
  password text not null,
  created_at timestamp without time zone default now(),
  updated_at timestamp without time zone default now(),
  deleted_at timestamp without time zone default null
);

insert into users
    (id, family_name, given_Name, patronymic, email, password)
values
    (gen_random_uuid(), 'Иванов', 'Иван', 'Иванович', 'ivany@bb.ru', '$2y$10$nfswjWcVAGmq2DvvIHGkj.5ltNG8BlEac0zZxdlai4/gcwZmKBr1W');

insert into users
    (id, family_name, given_Name, patronymic, email, password)
values
    (gen_random_uuid(), 'Петров', 'Алексей', 'Владимирович', 'pav@mm.ru', '$2y$10$EhKcSTUf3gPQ9D31S3iH5eqB6DXeI34fB3NuPTPHswqKQDWlavzmC');

-- +migrate Down
