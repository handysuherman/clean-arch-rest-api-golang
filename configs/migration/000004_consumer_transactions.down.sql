ALTER TABLE `consumer_transactions` DROP FOREIGN KEY `fk_consumer_transactions_consumer_id`;

ALTER TABLE `consumer_transactions` DROP FOREIGN KEY `fk_consumer_transactions_affiliated_dealer_id`;

ALTER TABLE `consumer_transactions` DROP FOREIGN KEY `uq_consumer_contract_dealer`;

DROP TABLE IF EXISTS `consumer_transactions` CASCADE;