begin;
create table if not exists Account (
    id bigint generated always as identity primary key,
    name text not null,
    email text not null,
    secret text not null,
    created_at timestamptz not null default current_timestamp,
    update_at timestamptz not null default current_timestamp,
    deleted_at timestamptz
);
commit;