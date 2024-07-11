begin;
create table if not exists Client (
    id bigint generated always as identity primary key,
    reference text not null,
    full_name text not null,
    birth date not null,
    cpf text not null,
    rg text not null,
    phone text not null,
    cep text not null,
    address text not null,
    address_number int not null,
    district text not null,
    city text not null,
    state text not null,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp,
    deleted_at timestamptz
);
commit;
