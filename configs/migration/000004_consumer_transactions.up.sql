CREATE TABLE `consumer_transactions` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `consumer_id` bigint NOT NULL,
  `contract_number` varchar(26) NOT NULL COMMENT 'Nomor Kontrak untuk setiap transaksi konsumen',
  `admin_fee_amount` decimal(15,2) COMMENT 'Angka fee admin',
  `installment_amount` decimal(15,2) COMMENT 'Angka jumlah cicilan',
  `otr_amount` decimal(15,2) COMMENT 'Angka On The Road transaksi barang',
  `interest_rate` decimal(15,2) COMMENT 'angka bunga yang ditangihkan setiap transaksi',
  `transaction_date` varchar(255) NOT NULL COMMENT 'format should be like 0001-01-01 00:00:00Z',
  `affiliated_dealer_id` bigint NOT NULL,
  `created_at` varchar(255) NOT NULL COMMENT 'format should be like 0001-01-01 00:00:00Z',
  `updated_at` varchar(255) NOT NULL DEFAULT '0001-01-01 00:00:00Z' COMMENT 'format should be like 0001-01-01 00:00:00Z',
  `updated_by` varchar(255)
);

ALTER TABLE `consumer_transactions`
  ADD CONSTRAINT `fk_consumer_transactions_consumer_id`
  FOREIGN KEY (`consumer_id`) REFERENCES `consumers` (`id`);

ALTER TABLE `consumer_transactions`
  ADD CONSTRAINT `fk_consumer_transactions_affiliated_dealer_id`
  FOREIGN KEY (`affiliated_dealer_id`) REFERENCES `affiliated_dealers` (`id`);
