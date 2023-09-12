CREATE TABLE
  `user` (
    `id` INT (11) NOT NULL AUTO_INCREMENT,
    `first_name` VARCHAR(150) NOT NULL,
    `last_name` VARCHAR(250) NOT NULL
    SET
      utf8 COLLATE utf8_general_ci,
      `document_number` VARCHAR(250) NOT NULL,
      `signup_date` DATETIME (250) NOT NULL,
      PRIMARY KEY (`id`)
  ) ENGINE = InnoDB;