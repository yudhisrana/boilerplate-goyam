/**
this is a sample file migration
you can use goose to migrate your database
remove this file and add file you migration in this folder
**/

-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_user (
    id char(36) PRIMARY KEY NOT NULL,
    first_name varchar(255) NOT NULL,
    last_name varchar(255) NOT NULL,
    email varchar(255) UNIQUE NOT NULL,
    address text,
    role_id int UNSIGNED NOT NULL,
    status boolean NOT NULL DEFAULT TRUE,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL,
    deleted_at timestamp NULL,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tbl_user;
-- +goose StatementEnd