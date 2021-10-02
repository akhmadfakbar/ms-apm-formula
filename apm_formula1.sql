-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Oct 02, 2021 at 12:37 PM
-- Server version: 10.4.17-MariaDB
-- PHP Version: 8.0.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `apm`
--

-- --------------------------------------------------------

--
-- Table structure for table `apm_formula1`
--

CREATE TABLE `apm_formula1` (
  `id` int(11) NOT NULL,
  `metric_id` text NOT NULL,
  `metric_name` text NOT NULL,
  `value_1_bb` float NOT NULL,
  `value_1_ba` float NOT NULL,
  `value_2_bb` float NOT NULL,
  `value_2_ba` float NOT NULL,
  `value_3_bb` float NOT NULL,
  `value_3_ba` float NOT NULL,
  `value_4_bb` float NOT NULL,
  `value_4_ba` float NOT NULL,
  `value_5_bb` float NOT NULL,
  `value_5_ba` float NOT NULL,
  `is_range` tinyint(1) NOT NULL,
  `created_at` datetime NOT NULL,
  `created_by` text NOT NULL,
  `modified_at` datetime NOT NULL,
  `modified_by` text NOT NULL,
  `record_status` int(2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `apm_formula1`
--

INSERT INTO `apm_formula1` (`id`, `metric_id`, `metric_name`, `value_1_bb`, `value_1_ba`, `value_2_bb`, `value_2_ba`, `value_3_bb`, `value_3_ba`, `value_4_bb`, `value_4_ba`, `value_5_bb`, `value_5_ba`, `is_range`, `created_at`, `created_by`, `modified_at`, `modified_by`, `record_status`) VALUES
(1, 'IC01', 'Investment Value', 0, 10000000, 10000000, 100000000, 100000000, 1000000000, 1000000000, 10000000000, 10000000000, 100000000000, 1, '0000-00-00 00:00:00', 'SYSTEM', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 1),
(2, 'AC01', 'ITEA Alignment', 1, 20, 21, 40, 41, 60, 61, 80, 81, 100, 1, '0000-00-00 00:00:00', '', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0),
(3, 'TQ021', 'Response Time', 0, 3, 4, 10, 11, 60, 61, 120, 121, 1000000000, 1, '0000-00-00 00:00:00', '', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0),
(4, 'TQ022', 'Resource Utilization', 0, 20, 21, 40, 41, 60, 61, 80, 81, 100, 1, '0000-00-00 00:00:00', '', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0),
(5, 'TQ023', 'Capacity Measure', 0, 20, 21, 40, 41, 60, 61, 80, 81, 100, 1, '0000-00-00 00:00:00', '', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0),
(6, 'TQ041', 'MTBF', 0, 3, 4, 7, 8, 30, 31, 60, 61, 1000000000, 1, '0000-00-00 00:00:00', '', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0),
(7, 'TQ043', 'Availability', 0, 79, 80, 84, 85, 89, 90, 99.4, 99.5, 100, 1, '0000-00-00 00:00:00', '', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0),
(8, 'TC08', 'Expected Retirement Date', 8, 7, 6, 5, 4, 3, 2, 1, 0.5, 2, 1, '0000-00-00 00:00:00', '', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0),
(9, 'TC09', 'Application Age', 8, 100000000, 6, 7, 4, 5, 2, 3, 0, 1, 1, '0000-00-00 00:00:00', '', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0),
(10, 'TC12', 'Application Downtime', 0, 79, 80, 84, 85, 89, 90, 99.4, 99.5, 100, 1, '0000-00-00 00:00:00', '', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `apm_formula1`
--
ALTER TABLE `apm_formula1`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `apm_formula1`
--
ALTER TABLE `apm_formula1`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
