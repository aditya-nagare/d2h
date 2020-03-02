--
-- Database: `d2h`
--

-- --------------------------------------------------------

--
-- Table structure for table `channels`
--

CREATE TABLE `channels` (
  `id` int(20) NOT NULL,
  `name` varchar(500) NOT NULL,
  `price` int(20) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `channels`
--

INSERT INTO `channels` (`id`, `name`, `price`) VALUES
(1, 'Zee', 10),
(2, 'Sony', 15),
(3, 'Star Plus', 20),
(4, 'Discovery', 10),
(5, 'NatGeo', 20);

-- --------------------------------------------------------

--
-- Table structure for table `packs`
--

CREATE TABLE `packs` (
  `id` int(20) NOT NULL,
  `name` varchar(200) NOT NULL,
  `price` int(20) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `packs`
--

INSERT INTO `packs` (`id`, `name`, `price`) VALUES
(1, 'SILVER', 50),
(2, 'GOLD', 100);

-- --------------------------------------------------------

--
-- Table structure for table `pack_compositions`
--

CREATE TABLE `pack_compositions` (
  `id` int(20) NOT NULL,
  `pack_id` int(20) NOT NULL,
  `channel_id` int(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `pack_compositions`
--

INSERT INTO `pack_compositions` (`id`, `pack_id`, `channel_id`) VALUES
(1, 1, 1),
(2, 1, 2),
(3, 1, 3),
(4, 2, 1),
(5, 2, 2),
(6, 2, 3),
(7, 2, 4),
(8, 2, 5);

-- --------------------------------------------------------

--
-- Table structure for table `services`
--

CREATE TABLE `services` (
  `id` int(20) NOT NULL,
  `name` varchar(500) NOT NULL,
  `price` int(20) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `services`
--

INSERT INTO `services` (`id`, `name`, `price`) VALUES
(1, 'LearnEnglish Service', 200),
(2, 'LearnCooking Service', 100);

-- --------------------------------------------------------

--
-- Table structure for table `subscriptions`
--

CREATE TABLE `subscriptions` (
  `id` int(20) NOT NULL,
  `user_id` int(20) NOT NULL,
  `service_id` int(20) DEFAULT NULL,
  `pack_id` int(20) DEFAULT NULL,
  `channel_id` int(20) DEFAULT NULL,
  `duration` int(10) UNSIGNED NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `subscriptions`
--

INSERT INTO `subscriptions` (`id`, `user_id`, `service_id`, `pack_id`, `channel_id`, `duration`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 1, 0, 1, 0, 3, '2020-02-23 11:11:59', '2020-02-23 11:11:59', NULL),
(3, 1, 2, 0, 0, 1, '2020-02-23 11:22:57', '2020-02-23 11:22:57', NULL),
(4, 1, 0, 0, 1, 1, '2020-02-23 11:56:37', '2020-02-23 11:56:37', NULL),
(5, 2, 0, 2, 0, 3, '2020-02-25 08:58:58', '2020-02-25 08:58:58', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(10) NOT NULL,
  `name` varchar(500) NOT NULL,
  `balance` int(20) UNSIGNED NOT NULL DEFAULT '100',
  `email` varchar(100) NOT NULL,
  `phone` varchar(10) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `balance`, `email`, `phone`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'aditya', 1000, 'adi@adi.com', '8149507255', '2020-02-23 07:02:41', '2020-02-25 08:54:42', NULL),
(2, 'ganesh', 1000, 'ganesh@gh.com', '1234567890', '2020-02-23 07:02:41', '2020-02-25 11:42:56', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `channels`
--
ALTER TABLE `channels`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `packs`
--
ALTER TABLE `packs`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `pack_compositions`
--
ALTER TABLE `pack_compositions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `pack_id_fk` (`pack_id`),
  ADD KEY `channel_id_fk` (`channel_id`);

--
-- Indexes for table `services`
--
ALTER TABLE `services`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `subscriptions`
--
ALTER TABLE `subscriptions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `channel_id_fkey` (`channel_id`),
  ADD KEY `pack_id_fkey` (`pack_id`),
  ADD KEY `service_id_fk` (`service_id`),
  ADD KEY `user_id_fk` (`user_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `channels`
--
ALTER TABLE `channels`
  MODIFY `id` int(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
--
-- AUTO_INCREMENT for table `packs`
--
ALTER TABLE `packs`
  MODIFY `id` int(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
--
-- AUTO_INCREMENT for table `pack_compositions`
--
ALTER TABLE `pack_compositions`
  MODIFY `id` int(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;
--
-- AUTO_INCREMENT for table `services`
--
ALTER TABLE `services`
  MODIFY `id` int(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
--
-- AUTO_INCREMENT for table `subscriptions`
--
ALTER TABLE `subscriptions`
  MODIFY `id` int(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
--
-- Constraints for dumped tables
--

--
-- Constraints for table `pack_compositions`
--
ALTER TABLE `pack_compositions`
  ADD CONSTRAINT `channel_id_fk` FOREIGN KEY (`channel_id`) REFERENCES `channels` (`id`) ON UPDATE CASCADE,
  ADD CONSTRAINT `pack_id_fk` FOREIGN KEY (`pack_id`) REFERENCES `packs` (`id`) ON UPDATE CASCADE;