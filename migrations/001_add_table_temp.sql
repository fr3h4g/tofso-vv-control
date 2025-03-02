-- +goose Up
-- +goose StatementBegin
CREATE TABLE `temperature` (
  `id` INT NOT NULL AUTO_INCREMENT , 
  `timestamp` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP , 
  `temperature` DECIMAL(10,2) NOT NULL , 
  `humidity` DECIMAL(10,2) NOT NULL , 
  PRIMARY KEY (`id`)
) ENGINE = InnoDB; 
-- +goose StatementEnd
