-- +goose Up

-- +goose StatementBegin

CREATE TABLE `users` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(255) NOT NULL,
	`email` VARCHAR(255) NOT NULL,
	`password` VARCHAR(255) NOT NULL,
	`flags` BIGINT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME DEFAULT NULL,
	PRIMARY KEY (id),
	UNIQUE KEY ukey_users_email (email)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- +goose StatementEnd

-- +goose StatementBegin

-- CREATE TABLE `coffee_shops` (

--   `id` INT NOT NULL AUTO_INCREMENT,

--   `name` VARCHAR(255) NOT NULL,

--   `shop_url` VARCHAR(255),

--   `open_time` DATETIME,

--   `close_time` DATETIME,

--   `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

--   `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

--   `closed_at` DATETIME,

--   PRIMARY KEY (`id`)

-- );

-- +goose StatementEnd

-- +goose StatementBegin

-- CREATE TABLE `coffee_beans` (

--   `id` INT NOT NULL AUTO_INCREMENT,

--   `name` VARCHAR(255) NOT NULL,

--   `farm_name` VARCHAR(255),

--   `country` VARCHAR(255),

--   -- `shop_id` INT NOT NULL,

--   `roast_degree` VARCHAR(255) NOT NULL,

--   `roasted_at` DATETIME,

--   `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

--   `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

--   PRIMARY KEY (`id`)

-- ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- +goose StatementEnd

-- +goose StatementBegin

CREATE TABLE `user_coffee_beans` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`status` INT NOT NULL,
	`user_id` INT NOT NULL,
	`name` VARCHAR(255) NOT NULL,
	`farm_name` VARCHAR(255),
	`country` VARCHAR(255),
	`roast_degree` VARCHAR(255) NOT NULL,
	`roasted_at` DATETIME,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	CONSTRAINT fkey_user_coffee_beans_user_id FOREIGN KEY (user_id) REFERENCES users (id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- +goose StatementEnd

-- +goose StatementBegin

-- CREATE TABLE `users_coffee_beans` (

--   `id` INT NOT NULL AUTO_INCREMENT,

--   `status` INT NOT NULL,

--   `user_id` INT NOT NULL,

--   `coffee_bean_id` INT NOT NULL,

--   `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

--   `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

--   `deleted_at` DATETIME DEFAULT NULL,

--   PRIMARY KEY (`id`),

--   CONSTRAINT fkey_users_coffee_beans_user_id FOREIGN KEY (user_id) REFERENCES users (id),

--   CONSTRAINT fkey_users_coffee_beans_coffee_bean_id FOREIGN KEY (coffee_bean_id) REFERENCES coffee_beans (id)

-- ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- +goose StatementEnd

-- +goose StatementBegin

CREATE TABLE `user_drip_recipes` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`user_id` INT NOT NULL,
	`coffee_bean_id` INT NOT NULL,
	`coffee_bean_weight` DOUBLE NOT NULL,
	`coffee_bean_grind` VARCHAR(255) NOT NULL,
	`liquid_weight` DOUBLE NOT NULL,
	`temperature` DOUBLE NOT NULL,
	`step_one` VARCHAR(255) NOT NULL,
	`step_two` VARCHAR(255) NOT NULL,
	`memo` VARCHAR(255) NOT NULL,
	`attrs` JSON,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME DEFAULT NULL,
	PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin

-- DROP TABLE `coffee_shops`;

-- +goose StatementEnd

-- +goose StatementBegin

DROP TABLE `user_coffee_beans`;

-- +goose StatementEnd

-- +goose StatementBegin

-- DROP TABLE `users_coffee_beans`;

-- +goose StatementEnd

-- +goose StatementBegin

-- DROP TABLE `coffee_beans`;

-- +goose StatementEnd

-- +goose StatementBegin

-- DROP TABLE `drip_recipes`;

-- +goose StatementEnd

-- +goose StatementBegin

DROP TABLE `users`;

-- +goose StatementEnd