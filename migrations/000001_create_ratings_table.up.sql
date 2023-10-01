CREATE TABLE IF NOT EXISTS ratings (
  `id` INT NOT NULL,
  `rating_value` INT NULL,
  `create_at` TIMESTAMP NULL,
  `update_at` TIMESTAMP NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB