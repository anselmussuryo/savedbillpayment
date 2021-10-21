CREATE TABLE `saved_bill_payment` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `merchant_id` bigint NOT NULL,
  `customer_id` bigint NOT NULL,
  `subsciber_no` bigint NOT NULL,
  `description` varchar(255),
  `isshowib` varchar(5),
  `isshowmobile` varchar(5),
  `amount` decimal(20,4),
  `isfavorite` varchar(10)
);
