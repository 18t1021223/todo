-- +goose Up
-- +goose StatementBegin
create table todo
(
    id          varchar(100) primary key,
    title       varchar(100) not null,
    description varchar(500)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table todo;
-- +goose StatementEnd
