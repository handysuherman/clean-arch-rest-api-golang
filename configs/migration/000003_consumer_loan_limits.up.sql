CREATE TABLE `consumer_loan_limits` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `consumer_id` bigint NOT NULL,
  `tenor` smallint NOT NULL,
  `amount` decimal(15,2) NOT NULL,
  `created_at` varchar(255) NOT NULL COMMENT 'format should be like 0001-01-01 00:00:00Z',
  `updated_at` varchar(255) NOT NULL DEFAULT '0001-01-01 00:00:00Z' COMMENT 'format should be like 0001-01-01 00:00:00Z',
  `updated_by` varchar(255)
);

CREATE INDEX `consumer_loan_limits_index_3` ON `consumer_loan_limits` (`consumer_id`);
CREATE INDEX `consumer_loan_limits_index_4` ON `consumer_loan_limits` (`tenor`);

ALTER TABLE `consumer_loan_limits` 
  ADD CONSTRAINT `fk_consumer_loan_limits_consumer_id` 
  FOREIGN KEY (`consumer_id`) REFERENCES `consumers` (`id`);
