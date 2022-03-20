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
	`deleted_at` DATETIME,
	PRIMARY KEY (id),
	UNIQUE KEY ukey_users_email (email)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE `coffee_shops` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(255) NOT NULL,
	`shop_url` VARCHAR(255),
	`open_time` DATETIME,
	`close_time` DATETIME,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`closed_at` DATETIME,
	PRIMARY KEY (`id`)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE `coffee_beans` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(255) NOT NULL,
	`farm_name` VARCHAR(255),
	`country` VARCHAR(255),
	`shop_id` VARCHAR(255) NOT NULL,
	`roasted_degree` VARCHAR(255) NOT NULL,
	`roasted_at` DATETIME,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE `drip_recipes` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`user_id` INT NOT NULL,
	`coffee_bean_id` INT NOT NULL,
	`coffee_bean_weight` DOUBLE NOT NULL,
	`liquid_weight` DOUBLE NOT NULL,
	`temperature` DOUBLE NOT NULL,
	`steam_time` INT NOT NULL,
	`drip_time` INT NOT NULL,
	`memo` VARCHAR(255) NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `users`;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE `coffee_shops`;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE `coffee_beans`;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE `drip_recipes`;
-- +goose StatementEnd
