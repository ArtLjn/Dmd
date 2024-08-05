# ************************************************************
# Sequel Ace SQL dump
# 版本号： 20064
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# 主机: localhost (MySQL 8.3.0)
# 数据库: xcr_person_health
# 生成时间: 2024-08-04 13:57:47 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# 转储表 health_info
# ------------------------------------------------------------

DROP TABLE IF EXISTS `health_info`;

CREATE TABLE `health_info` (
  `id` bigint unsigned NOT NULL,
  `height` bigint unsigned DEFAULT NULL,
  `weight` bigint unsigned DEFAULT NULL,
  `bmi` bigint unsigned DEFAULT NULL,
  `age` bigint unsigned DEFAULT NULL,
  `heart_rate` bigint unsigned DEFAULT NULL,
  `vision` bigint unsigned DEFAULT NULL,
  `blood_pressure` bigint unsigned DEFAULT NULL,
  `blood_sugar` bigint unsigned DEFAULT NULL,
  `upload_time` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
