CREATE TABLE IF NOT EXISTS books (
    `id` INT NOT NULL AUTO_INCREMENT,
    `isbn` VARCHAR(45) NULL,
    `title` VARCHAR(45) NULL,
    `stock` INT NULL,
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `publisher_id` INT NOT NULL,
    PRIMARY KEY (`id`, `publisher_id`),
    UNIQUE INDEX `title_UNIQUE` (`title` ASC),
    INDEX `fk_book_publisher1_idx` (`publisher_id` ASC),
    CONSTRAINT `fk_book_publisher1` FOREIGN KEY (`publisher_id`) REFERENCES publishers (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE = InnoDB;