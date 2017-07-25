CREATE TABLE `userinfo` (
	`uid` INT(10) NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(64) NULL DEFAULT NULL,
	`token` VARCHAR(64) NULL DEFAULT NULL,
	`created` DATE NULL DEFAULT NULL,
	`payed` DATE NULL DEFAULT NULL,
	PRIMARY KEY (`uid`)
);

CREATE TABLE `wallet` (
	`uid` INT(10) NOT NULL,
	`balance` FLOAT NULL,
	`freeze` FLOAT NULL,
	`amount` FLOAT NULL,
	PRIMARY KEY (`uid`)
)