CREATE TABLE IF NOT EXISTS `squadcast`.`restaurant` (
  `restaurant_id` INT NOT NULL,
  `name` VARCHAR(64) NULL,
  `url` VARCHAR(256) NULL,
  `cuisines` VARCHAR(256) NULL,
  `image` VARCHAR(256) NULL,
  `address` VARCHAR(512) NULL,
  `city` VARCHAR(64) NULL,
  `rating` FLOAT NULL,
  `veg` BOOLEAN DEFAULT FALSE NULL,
  `createdon` DATETIME NULL DEFAULT current_timestamp,
  `updatedon` DATETIME NULL DEFAULT current_timestamp on update current_timestamp,
  PRIMARY KEY (`restaurant_id`));