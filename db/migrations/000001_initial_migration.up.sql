create table users (
  id text primary key,
  name text not null,
  email text unique not null,
  created_at timestamp default current_timestamp
);

create table sessions (
  id text primary key,
  user_id text not null,
  created_at timestamp default current_timestamp,
  foreign key (user_id) references users (id) on delete cascade
);
