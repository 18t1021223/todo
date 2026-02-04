-- +goose Up
-- +goose StatementBegin
create table user
(
    id       varchar(100) primary key,
    name     varchar(50)        not null,
    email    varchar(50) unique not null,
    password binary(80) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table user;
-- +goose StatementEnd
