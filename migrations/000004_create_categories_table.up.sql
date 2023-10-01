CREATE TABLE IF NOT EXISTS categories (
    `id` INT NOT NULL,
    `categories` VARCHAR(45) NULL,
    `description` VARCHAR(45) NULL,
    `create_at` TIMESTAMP NULL,
    `update_at` TIMESTAMP NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB