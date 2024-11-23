-- name: ProjectsGetAll :many
SELECT * FROM projects;

-- name: ProjectsGetNoUI :many
SELECT * FROM projects WHERE has_ui = 0;

-- name: ProjectsGetUI :many
SELECT * FROM projects WHERE has_ui = 1 AND ? <= points AND points <= ?;
