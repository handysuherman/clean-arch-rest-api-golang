ALTER TABLE `consumer_loan_limits` DROP FOREIGN KEY `fk_consumer_loan_limits_consumer_id`;

DROP TABLE IF EXISTS `consumer_loan_limits` CASCADE;