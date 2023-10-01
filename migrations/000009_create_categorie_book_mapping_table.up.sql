CREATE TABLE IF NOT EXISTS categories_book_mapping (
    `categories_id` INT NOT NULL,
    `book_id` INT NOT NULL,
    PRIMARY KEY (`categories_id`, `book_id`),
    INDEX `fk_categories_has_book_book1_idx` (`book_id` ASC),
    INDEX `fk_categories_has_book_categories1_idx` (`categories_id` ASC),
    CONSTRAINT `fk_categories_has_book_categories1` FOREIGN KEY (`categories_id`) REFERENCES categories (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
    CONSTRAINT `fk_categories_has_book_book1` FOREIGN KEY (`book_id`) REFERENCES books (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE = InnoDB