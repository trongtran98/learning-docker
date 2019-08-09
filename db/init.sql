create table accounts
(
    id         int unsigned auto_increment
        primary key,
    created_at timestamp    null,
    updated_at timestamp    null,
    deleted_at timestamp    null,
    email      varchar(255) null,
    password   varchar(255) null,
    token      varchar(255) null
);

create index idx_accounts_deleted_at
    on accounts (deleted_at);

create table contacts
(
    id         int unsigned auto_increment
        primary key,
    created_at timestamp    null,
    updated_at timestamp    null,
    deleted_at timestamp    null,
    name       varchar(255) null,
    phone      varchar(255) null,
    user_id    int unsigned null
);

create index idx_contacts_deleted_at
    on contacts (deleted_at);

INSERT INTO accounts (id, created_at, updated_at, deleted_at, email, password, token) VALUES (1, '2019-01-30 09:57:50', '2019-01-30 09:57:50', null, 'test@gmail.com', '$2a$10$xVK88WmV9Eudpe/Nvg6HhetC8RQPzFnictr4qkuQaLvV1lkYjon/W', '');

INSERT INTO contacts (id, created_at, updated_at, deleted_at, name, phone, user_id) VALUES (1, '2019-08-08 11:10:51', '2019-08-08 11:10:54', null, 'vi', '0345123', 1);