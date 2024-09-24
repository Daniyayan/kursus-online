CREATE DATABASE IF NOT EXISTS kursus;

USE kursus;


CREATE TABLE IF NOT EXISTS `user` (
    `id_user` int(3) NOT NULL,
    `username` varchar(50) NOT NULL,
    `nama_user` varchar(50) NOT NULL,
    `password` varchar(50) NOT NULL,
    `email` varchar(50) DEFAULT NULL,
    `telepon` varchar(13) DEFAULT NULL,
    `foto` varchar(100) DEFAULT NULL,
    `hak_akses` enum('Super Admin','user') NOT NULL,
    `status` enum('aktif','blokir') NOT NULL DEFAULT 'aktif',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE IF NOT EXISTS `matpel` (
    `id_mapel` INT PRIMARY KEY AUTO_INCREMENT,
    `nama_mata_pelajaran` VARCHAR (100) NOT NULL,
    `maple` VARCHAR (100) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE=InnoDB DEFAULT CHARSET=latin1;



CREATE TABLE IF NOT EXISTS `register` (
    `id_register` INT PRIMARY KEY AUTO_INCREMENT,
    `id_user` INT NOT NULL,
    `id_mapel` INT NOT NULL,
    `tanggal_daftar` DATE NOT NULL,
    `status` VARCHAR(50),
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE=InnoDB DEFAULT CHARSET=latin1;





