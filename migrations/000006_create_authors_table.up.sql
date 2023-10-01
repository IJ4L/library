CREATE TABLE IF NOT EXISTS authors (
    `id` INT NOT NULL AUTO_INCREMENT,
    `author_name` VARCHAR(45) NULL,
    `biography` TEXT NULL,
    `create_at` TIMESTAMP NULL,
    `update_at` TIMESTAMP NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB