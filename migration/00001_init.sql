-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
	id INT NOT NULL AUTO_INCREMENT,
	username VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	flags BIGINT NOT NULL,
	last_logged_in_at DATETIME NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	deleted_at DATETIME,
	PRIMARY KEY (id),
	UNIQUE KEY ukey_users_email (email)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
