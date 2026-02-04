-- name: CreateUser :exec
insert into user(id, name, email, password)
values (?, ?, ?, ?);

-- name: GetUserByID :one
select *
from user
where id = ?;

-- name: GetUserByEmail :one
select *
from user
where email = ?;
