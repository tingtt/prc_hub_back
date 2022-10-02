SET
  SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";

START TRANSACTION;

SET
  time_zone = "+00:00";

--
-- Database: `prc_hub`
--
CREATE DATABASE IF NOT EXISTS `prc_hub` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

USE `prc_hub`;

-- --------------------------------------------------------
--
-- Table structure for table `users`
--
CREATE TABLE `users` (
  `id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL UNIQUE,
  `password` varchar(255) NOT NULL,
  `post_event_availabled` tinyint(1) NOT NULL DEFAULT 0,
  `manage` tinyint(1) NOT NULL DEFAULT 0,
  `admin` tinyint(1) NOT NULL DEFAULT 0,
  `twitter_id` varchar(255),
  `github_username` varchar(255),
  PRIMARY KEY (`id`)
);

--
-- Table structure for table `oauth2_tokens`
--
CREATE TABLE `oauth2_tokens` (
  `user_id` varchar(255) NOT NULL,
  `provider` varchar(255) NOT NULL,
  `token` varchar(255) NOT NULL,
  `owner_id` varchar(255) NOT NULL,
  PRIMARY KEY (`user_id`, `provider`)
);

--
-- Table structure for table `events`
--
CREATE TABLE `events` (
  `id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` varchar(255),
  `location` varchar(255),
  `published` tinyint(1) NOT NULL,
  `completed` tinyint(1) NOT NULL,
  `user_id` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
);

--
-- Table structure for table `event_datetimes`
--
CREATE TABLE `event_datetimes` (
  `event_id` varchar(255) NOT NULL,
  `start` datetime NOT NULL,
  `end` datetime NOT NULL,
  PRIMARY KEY (`event_id`, `start`),
  FOREIGN KEY (`event_id`) REFERENCES events(`id`) ON DELETE CASCADE
);

--
-- Table structure for table `documents`
--
CREATE TABLE `documents` (
  `id` varchar(255) NOT NULL,
  `event_id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `url` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`event_id`) REFERENCES events(`id`) ON DELETE CASCADE
);