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
-- Table structure for table `apm_formula2`
--

CREATE TABLE `apm_formula2` (
  `id` int(11) NOT NULL,
  `metric_id` text NOT NULL,
  `metric_name` text NOT NULL,
  `value_1` text NOT NULL,
  `value_2` text NOT NULL,
  `value_3` text NOT NULL,
  `value_4` text NOT NULL,
  `value_5` text NOT NULL,
  `is_range` tinyint(1) NOT NULL,
  `created_by` text NOT NULL,
  `modified_by` text NOT NULL,
  `created_at` datetime NOT NULL,
  `modified_at` datetime NOT NULL,
  `record_status` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `apm_formula2`
--

INSERT INTO `apm_formula2` (`id`, `metric_id`, `metric_name`, `value_1`, `value_2`, `value_3`, `value_4`, `value_5`, `is_range`, `created_by`, `modified_by`, `created_at`, `modified_at`, `record_status`) VALUES
(1, 'BV01', 'Criticality', 'Critical Severity 1', 'Critical Severity 2', 'High', 'Medium', 'Low', 0, '', '', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0),
(2, 'TC10', 'Technology Diversity', '5', '4', '3', '2', '1', 0, '', '', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `apm_formula2`
--
ALTER TABLE `apm_formula2`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `apm_formula2`
--
ALTER TABLE `apm_formula2`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
