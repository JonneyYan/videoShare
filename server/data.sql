CREATE TABLE `userinfo` (
	`uid` INT(10) NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(64) NULL DEFAULT NULL,
	`token` VARCHAR(64) NULL DEFAULT NULL,
	`created` DATE NULL DEFAULT NULL,
	PRIMARY KEY (`uid`)
);

CREATE TABLE `wallet` (
	`uid` INT(10) NOT NULL,
	`count` FLOAT NULL,
	`freeze` FLOAT NULL,
	PRIMARY KEY (`uid`)
)