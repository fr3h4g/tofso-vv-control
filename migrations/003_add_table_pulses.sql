-- +goose Up
-- +goose StatementBegin
CREATE TABLE `pulses` (
  `id` INT NOT NULL AUTO_INCREMENT , 
  `timestamp` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP , 
  `meter1` INT(1) NULL , 
  `meter2` INT(1) NULL , 
  `meter3` INT(1) NULL , 
  `meter4` INT(1) NULL , 
  PRIMARY KEY (`id`)
) ENGINE = InnoDB; 
-- +goose StatementEnd
