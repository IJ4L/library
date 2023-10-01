CREATE TABLE IF NOT EXISTS borrowings (
    `id` INT NOT NULL AUTO_INCREMENT,
    `borrowing_date` DATE NULL,
    `return_date` DATE NULL,
    `create_at` TIMESTAMP NULL,
    `update_at` TIMESTAMP NULL,
    `book_id` INT NOT NULL,
    `member_id` INT NOT NULL,
    PRIMARY KEY (`id`, `book_id`, `member_id`),
    INDEX `fk_loans_book1_idx` (`book_id` ASC),
    INDEX `fk_loans_member1_idx` (`member_id` ASC),
    CONSTRAINT `fk_loans_book1` FOREIGN KEY (`book_id`) REFERENCES books (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
    CONSTRAINT `fk_loans_member1` FOREIGN KEY (`member_id`) REFERENCES members (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE = InnoDB