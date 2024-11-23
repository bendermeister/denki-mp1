CREATE TABLE projects (
       id     INTEGER NOT NULL UNIQUE,
       name   TEXT NOT NULL,
       url    TEXT NOT NULL,
       has_ui INTEGER NOT NULL,
       points INTEGER NOT NULL,

       PRIMARY KEY(id)
);
