-- name: ProjectsGetAll :many
SELECT * FROM projects ORDER BY name ASC;

-- name: ProjectsGetNoUI :many
SELECT * FROM projects WHERE has_ui = 0 ORDER BY name ASC;

-- name: ProjectsGetUI :many
SELECT * FROM projects WHERE has_ui = 1 AND ? <= points AND points <= ? ORDER BY name ASC;

-- name: ProjectInsert :exec
INSERT INTO projects (name, url, has_ui, points) VALUES(?, ?, ?, ?);

-- name: ProjectCount :one
SELECT COUNT(*) FROM projects;
