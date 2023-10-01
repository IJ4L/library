CREATE TABLE IF NOT EXISTS publishers (
    `id` INT NOT NULL AUTO_INCREMENT,
    `publisher_name` VARCHAR(45) NULL,
    `publisher_address` VARCHAR(45) NULL,
    `create_at` TIMESTAMP NULL,
    `update_at` TIMESTAMP NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB