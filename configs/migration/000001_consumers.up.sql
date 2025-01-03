CREATE TABLE `consumers` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `nik` varchar(16) NOT NULL COMMENT 'Nomor KTP Konsumen',
  `full_name` varchar(255) NOT NULL COMMENT 'Nama Lengkap Konsumen',
  `legal_name` varchar(255) COMMENT 'Nama Konsumen Di KTP',
  `birth_place` varchar(255) COMMENT 'Tempat lahir konsumen sesuai KTP',
  `birth_date` DATE COMMENT 'Tanggal lhir konsumen sesuai KTP',
  `salary` decimal(15,2) COMMENT 'Gaji Konsumen',
  `ktp_photo` TEXT COMMENT 'Foto KTP Konsumen',
  `selfie_photo` TEXT COMMENT 'Foto Selfie Konsumen',
  `created_at` varchar(255) NOT NULL COMMENT 'format should be like 0001-01-01 00:00:00Z',
  `updated_at` varchar(255) NOT NULL DEFAULT '0001-01-01 00:00:00Z' COMMENT 'format should be like 0001-01-01 00:00:00Z',
  `updated_by` varchar(255),
  `is_activated` boolean NOT NULL DEFAULT true,
  `is_activated_at` varchar(255) NOT NULL DEFAULT '0001-01-01 00:00:00Z' COMMENT 'format should be like 0001-01-01 00:00:00Z',
  `is_activated_updated_at` varchar(255) NOT NULL DEFAULT '0001-01-01 00:00:00Z' COMMENT 'format should be like 0001-01-01 00:00:00Z'
);

CREATE UNIQUE INDEX `consumers_index_0` ON `consumers` (`nik`);

CREATE INDEX `consumers_index_1` ON `consumers` (`full_name`);

CREATE INDEX `consumers_index_2` ON `consumers` (`legal_name`);