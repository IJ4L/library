CREATE TABLE IF NOT EXISTS members (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(45) NULL,
    `address` VARCHAR(45) NULL,
    `identity_number` INT NULL,
    `create_at` TIMESTAMP NULL,
    `update_at` TIMESTAMP NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `identity_number_UNIQUE` (`identity_number` ASC)
) ENGINE = InnoDB