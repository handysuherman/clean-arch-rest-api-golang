CREATE TABLE `affiliated_dealers` (
  `id` bigint PRIMARY KEY NOT NULL,
  `affiliated_dealer_name` varchar(255) NOT NULL,
  `created_at` varchar(255) NOT NULL COMMENT 'format should be like 0001-01-01 00:00:00Z',
  `updated_at` varchar(255) NOT NULL DEFAULT '0001-01-01 00:00:00Z' COMMENT 'format should be like 0001-01-01 00:00:00Z',
  `updated_by` varchar(255),
  `is_activated` boolean NOT NULL DEFAULT true,
  `is_activated_at` varchar(255) NOT NULL COMMENT 'format should be like 0001-01-01 00:00:00Z',
  `is_activated_updated_at` varchar(255) NOT NULL DEFAULT '0001-01-01 00:00:00Z' COMMENT 'format should be like 0001-01-01 00:00:00Z'
);

CREATE INDEX `affiliated_dealers_index_9` ON `affiliated_dealers` (`affiliated_dealer_name`);
