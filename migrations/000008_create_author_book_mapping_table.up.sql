CREATE TABLE IF NOT EXISTS author_book_mapping (
    `book_id` INT NOT NULL,
    `book_publisher_id` INT NOT NULL,
    `author_id` INT NOT NULL,
    PRIMARY KEY (`book_id`, `book_publisher_id`, `author_id`),
    INDEX `fk_book_has_author_author1_idx` (`author_id` ASC),
    INDEX `fk_book_has_author_book1_idx` (`book_id` ASC, `book_publisher_id` ASC),
    CONSTRAINT `fk_book_has_author_book1` FOREIGN KEY (`book_id`, `book_publisher_id`) REFERENCES books (`id`, `publisher_id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
    CONSTRAINT `fk_book_has_author_author1` FOREIGN KEY (`author_id`) REFERENCES authors (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE = InnoDB