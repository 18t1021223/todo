-- name: CreateTodo :execresult
insert into todo(id, title, description)
values (?, ?, ?);

-- name: GetListTodo :many
select *
from todo
limit ? offset ?;

-- name: CountTodo :one
select count(*)
from todo;

-- name: UpdateTodoByID :exec
update todo
set title       = ?,
    description = ?
where id = ?;


-- name: DeleteTodoByID :exec
delete
from todo
where id = ?;
