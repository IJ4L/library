CREATE TABLE IF NOT EXISTS member_rating_mappings (
    `member_id` INT NOT NULL,
    `rating_id` INT NOT NULL,
    PRIMARY KEY (`member_id`, `rating_id`),
    INDEX `fk_member_has_rating_rating1_idx` (`rating_id` ASC),
    INDEX `fk_member_has_rating_member1_idx` (`member_id` ASC),
    CONSTRAINT `fk_member_has_rating_member1` FOREIGN KEY (`member_id`) REFERENCES members (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
    CONSTRAINT `fk_member_has_rating_rating1` FOREIGN KEY (`rating_id`) REFERENCES ratings (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE = InnoDB