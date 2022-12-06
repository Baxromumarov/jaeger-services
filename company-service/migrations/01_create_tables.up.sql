
create table if not exists company (
    id uuid primary key not null,
    name varchar not null,
    product_type varchar(255),
    created_at timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null
);
