-- phpMyAdmin SQL Dump
-- version 4.6.5.2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: May 02, 2017 at 10:54 PM
-- Server version: 5.6.34
-- PHP Version: 5.6.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- Database: `testdb`
--

-- --------------------------------------------------------

--
-- Table structure for table `Test_Table`
--

CREATE TABLE `Test_Table` (
  `id` int(11) NOT NULL,
  `email` char(60) DEFAULT NULL,
  `name` char(60) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `Test_Table`
--
ALTER TABLE `Test_Table`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `Test_Table`
--
ALTER TABLE `Test_Table`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;