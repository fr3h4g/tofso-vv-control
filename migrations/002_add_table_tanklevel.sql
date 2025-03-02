-- +goose Up
-- +goose StatementBegin
CREATE TABLE `vvcontrol`.`tanklevel` (
  `id` INT NOT NULL AUTO_INCREMENT , 
  `timestamp` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP , 
  `sensor1` INT(1) NULL , 
  `sensor2` INT(1) NULL , 
  `sensor3` INT(1) NULL , 
  `sensor4` INT(1) NULL , 
  PRIMARY KEY (`id`)
) ENGINE = InnoDB; 
-- +goose StatementEnd
