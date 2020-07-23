-- phpMyAdmin SQL Dump
-- version 4.5.4.1deb2ubuntu2.1
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: Jul 08, 2020 at 08:12 AM
-- Server version: 5.7.30-0ubuntu0.16.04.1
-- PHP Version: 7.2.11-3+ubuntu16.04.1+deb.sury.org+1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `recruit_app_go`
--
CREATE DATABASE IF NOT EXISTS `recruit_app_go` DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;
USE `recruit_app_go`;

-- --------------------------------------------------------

--
-- Table structure for table `access_level`
--

CREATE TABLE `access_level` (
  `id` int(11) NOT NULL,
  `user_type` varchar(65) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `applications`
--

CREATE TABLE `applications` (
  `id` int(11) NOT NULL,
  `userID` int(11) NOT NULL,
  `postID` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `job_post_document_id` varchar(255) NOT NULL,
  `application_document_id` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `job_posts`
--

CREATE TABLE `job_posts` (
  `id` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `organization` varchar(255) NOT NULL,
  `created_by` int(11) NOT NULL,
  `job_post_document_id` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `job_posts`
--

INSERT INTO `job_posts` (`id`, `title`, `organization`, `created_by`, `job_post_document_id`) VALUES
(1, 'Vacancy for Java Developers', 'ButterBee IT Services', 1, NULL),
(2, 'Systems Engineer', 'ButterBee IT Services', 1, NULL),
(3, 'Technical Support Officer', 'ButterBee IT Services', 1, NULL);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `first_name` varchar(65) NOT NULL,
  `last_name` varchar(65) NOT NULL,
  `other_names` varchar(65) DEFAULT NULL,
  `contact` varchar(65) NOT NULL,
  `email` varchar(300) NOT NULL,
  `password` varchar(300) NOT NULL,
  `user_type` varchar(65) NOT NULL DEFAULT 'user',
  `created_on` datetime NOT NULL,
  `modified_on` datetime DEFAULT NULL,
  `user_document_id` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `first_name`, `last_name`, `other_names`, `contact`, `email`, `password`, `user_type`, `created_on`, `modified_on`, `user_document_id`) VALUES
(1, 'Knaomi', 'Ampofo', 'Amponfi', '0559633956', 'jigbaeboo@gmail.com', '$2a$04$9xQpsSZ8ZAYRv6W95qJTsuMVE05T9pxfKf6QAvHXnj7IwZDYQKP0i', 'recruiter', '2020-06-14 13:52:06', '2020-06-14 13:52:26', NULL),
(2, 'Naomi', 'Amponfi', '', '0559633956', 'naomibae@gmail.com', '$2a$04$khLInjuZ.ELifgSW8ZqV4OT9D6uOqoQCXRDa46v4WxoCad5wPhN1u', 'user', '2020-07-06 19:52:20', NULL, NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `access_level`
--
ALTER TABLE `access_level`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `applications`
--
ALTER TABLE `applications`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `job_posts`
--
ALTER TABLE `job_posts`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `access_level`
--
ALTER TABLE `access_level`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `applications`
--
ALTER TABLE `applications`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
--
-- AUTO_INCREMENT for table `job_posts`
--
ALTER TABLE `job_posts`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
