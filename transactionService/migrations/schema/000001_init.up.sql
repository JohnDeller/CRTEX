CREATE TABLE IF NOT EXISTS users
(
    id uuid not null unique,
    balance float,
    created_at timestamp not null default CURRENT_TIMESTAMP,
    deleted_at timestamp null default null
);

CREATE TABLE IF NOT EXISTS transactions
(
    id uuid not null unique,
    user_id uuid references users(id) on delete cascade not null,
    operation_type varchar(36) not null,
    price float,
    currency varchar(24),
    created_at timestamp not null default CURRENT_TIMESTAMP
);
