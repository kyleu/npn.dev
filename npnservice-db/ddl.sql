-- user
create type system_role as enum ('guest', 'user', 'admin');

create table if not exists "system_user" (
  "id" uuid not null primary key,
  "name" varchar(2048) not null,
  "role" system_role not null,
  "theme" varchar(32) not null,
  "nav_color" varchar(32) not null,
  "link_color" varchar(32) not null,
  "picture" text not null,
  "locale" varchar(32) not null,
  "created" timestamp not null default now()
);

-- auth
create type auth_provider as enum ('team', 'sprint', 'github', 'google', 'slack', 'facebook', 'amazon', 'microsoft');

create table if not exists "auth" (
  "id" uuid not null primary key,
  "user_id" uuid not null references "system_user"("id"),
  "provider" auth_provider not null,
  "provider_id" text not null,
  "user_list_id" varchar(512) not null,
  "user_list_name" varchar(2048) not null,
  "access_token" text not null,
  "expires" timestamp,
  "name" varchar(2048) not null,
  "email" varchar(2048) not null,
  "picture" text not null,
  "created" timestamp not null default now()
);

create index if not exists idx_auth_provider_provider_id on auth(provider, provider_id);
