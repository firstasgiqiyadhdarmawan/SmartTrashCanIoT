-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 06 Jan 2023 pada 15.43
-- Versi server: 10.4.21-MariaDB
-- Versi PHP: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `monitoring_platform_iot`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `capacity_history`
--

CREATE TABLE `capacity_history` (
  `capacity_history_id` int(11) NOT NULL,
  `device_id` varchar(150) NOT NULL,
  `capacity` int(11) NOT NULL,
  `status_device` varchar(255) NOT NULL,
  `date_updated` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `capacity_history`
--

INSERT INTO `capacity_history` (`capacity_history_id`, `device_id`, `capacity`, `status_device`, `date_updated`) VALUES
(14, 'n5EpridtQcOhhE5i', 11, 'Penuh', '2023-01-06 19:23:46'),
(15, 'n5EpridtQcOhhE5i', 5, 'Penuh', '2023-01-06 19:24:37');

-- --------------------------------------------------------

--
-- Struktur dari tabel `device`
--

CREATE TABLE `device` (
  `device_id` varchar(150) NOT NULL,
  `device_name` varchar(255) NOT NULL,
  `date_created` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `device`
--

INSERT INTO `device` (`device_id`, `device_name`, `date_created`) VALUES
('n5EpridtQcOhhE5i', 'Sensor Dalam', '2023-01-05');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `capacity_history`
--
ALTER TABLE `capacity_history`
  ADD PRIMARY KEY (`capacity_history_id`),
  ADD KEY `device_id` (`device_id`);

--
-- Indeks untuk tabel `device`
--
ALTER TABLE `device`
  ADD PRIMARY KEY (`device_id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `capacity_history`
--
ALTER TABLE `capacity_history`
  MODIFY `capacity_history_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `capacity_history`
--
ALTER TABLE `capacity_history`
  ADD CONSTRAINT `capacity_history_ibfk_1` FOREIGN KEY (`device_id`) REFERENCES `device` (`device_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
