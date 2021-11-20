CREATE  DATABASE banking;
USE banking;

DROP TABLE IF EXISTS `customers`;
CREATE TABLE `customers` (
	`customer_id` int(11) NOT NULL AUTO_INCREMENT,
	`name` varchar(100) NOT NULL,
	`date_of_birth` date NOT NULL,
	`city` varchar(100) NOT NULL,
	`zipcode` varchar(100) NOT NULL,
	`status` tinyint(1) NOT NULL DEFAULT '1',
	PRIMARY KEY (`customer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;

INSERT INTO `customers` VALUES
	(2000, 'Danilo', '1960-5-18', 'New York', '12314856', 1),
	(2001, 'Janana', '1960-5-18', 'New York', '45645546', 1),
	(2002, 'Lavínia', '1960-5-18', 'New York', '89745602', 0),
	(2003, 'Cleunir', '1960-5-18', 'New York', '55889599', 1),
	(2004, 'Lucio', '1939-7-24', 'New York', '12318489', 1), 
    (2005, 'Ed', '1964-12-10', 'New York', '12318489', 0);
	
DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
	`account_id` int(11) NOT NULL AUTO_INCREMENT,
	`customer_id` int(11) NOT NULL,
	`opening_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`account_type` varchar(10) NOT NULL,
	`amount` int(11) NOT NULL,
	`pin` varchar(10) NOT NULL,
	`status` tinyint(4) NOT NULL DEFAULT '1',
	PRIMARY KEY (`account_id`),
	KEY `accounts_FK` (`customer_id`),
	CONSTRAINT `accounts_FK` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`customer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=95476 DEFAULT CHARSET=latin1;


INSERT INTO `accounts` VALUES
	(95470, 2000, '1960-5-18 10:20:06', 'Saving', '1075', 1),
	(95471, 2001, '1960-5-18 11:56:59', 'Saving', '4596', 1),
	(95472, 2002, '1960-5-18 11:56:59', 'Checking', '3258', 1),
	(95473, 2003, '1960-5-18 11:56:59', 'Saving', '7852', 1),
	(95474, 2004, '1939-7-24 11:56:59', 'Checking', '9513', 1),
	(95475, 2005, '1939-7-24 11:56:59', 'Saving', '4568', 0);


DROP TABLE IF EXISTS `transactions`;
CREATE TABLE `transactions` (
	`transaction_id` int(11) NOT NULL AUTO_INCREMENT,
	`account_id` int(11) NOT NULL ,
	`amount` int(11) NOT NULL,
	`transaction_type` varchar(10) NOT NULL,
	`transaction_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (`transaction_id`),
	KEY `transactions_FK` (`account_id`),
	CONSTRAINT `transactions_FK` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


SELECT * FROM 