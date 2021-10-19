CREATE TABLE `saved_bill_payment` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `merchant_id` int NOT NULL,
  `customer_id` int NOT NULL,
  `subsciber_no` int NOT NULL,
  `description` varchar(255),
  `isshowib` varchar(5),
  `isshowmobile` varchar(5),
  `amount` decimal(20,4),
  `isfavorite` varchar(10)
);
