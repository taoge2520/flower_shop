-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.7.16 - MySQL Community Server (GPL)
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  9.4.0.5125
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- 导出 flower_shop 的数据库结构
CREATE DATABASE IF NOT EXISTS `flower_shop` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `flower_shop`;

-- 导出  表 flower_shop.comment 结构
CREATE TABLE IF NOT EXISTS `comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user` varchar(50) NOT NULL,
  `content` varchar(255) NOT NULL,
  `commodity_id` int(11) NOT NULL,
  `orderid` int(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- 正在导出表  flower_shop.comment 的数据：~6 rows (大约)
DELETE FROM `comment`;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
INSERT INTO `comment` (`id`, `user`, `content`, `commodity_id`, `orderid`) VALUES
	(1, '无', '尚无用户评论', 0, 1),
	(2, '张三', '不错，很好', 1, 1),
	(3, '李四', '贼好看了', 1, 1),
	(4, 'test', 'this is good', 1, 2),
	(5, 'test1', 'dsfaf', 1, 3),
	(6, 'test1', 'sfasflsf', 1, 2),
	(7, 'test1', '很好，非常好', 1, 4);
/*!40000 ALTER TABLE `comment` ENABLE KEYS */;

-- 导出  表 flower_shop.commodity 结构
CREATE TABLE IF NOT EXISTS `commodity` (
  `id` int(11) NOT NULL,
  `name` varchar(50) NOT NULL,
  `price` int(11) NOT NULL,
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `count` int(11) NOT NULL,
  `picture` varchar(50) DEFAULT NULL,
  `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1为无存货'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  flower_shop.commodity 的数据：~8 rows (大约)
DELETE FROM `commodity`;
/*!40000 ALTER TABLE `commodity` DISABLE KEYS */;
INSERT INTO `commodity` (`id`, `name`, `price`, `add_time`, `count`, `picture`, `is_delete`) VALUES
	(1, '菊花1', 10, '2017-03-14 16:40:01', 10, 'static/img/juhua.jpg', 0),
	(2, '小兰花', 15, '2017-03-10 15:42:44', 20, 'static/img/lanhua.jpg', 0),
	(3, '兰花', 20, '2017-03-07 15:39:01', 30, 'static/img/lanhua.jpg', 0),
	(4, '菊花', 23, '2017-03-07 15:29:14', 12, 'static/img/juhua.jpg', 0),
	(5, '菊花', 16, '2017-03-07 15:29:30', 24, 'static/img/juhua.jpg', 0),
	(6, '菊花', 45, '2017-03-07 15:29:54', 64, 'static/img/juhua.jpg', 0),
	(7, '菊花', 45, '2017-03-07 15:30:19', 78, 'static/img/juhua.jpg', 0),
	(8, '菊花', 15, '2017-03-07 15:30:37', 50, 'static/img/juhua.jpg', 0);
/*!40000 ALTER TABLE `commodity` ENABLE KEYS */;

-- 导出  表 flower_shop.order_com 结构
CREATE TABLE IF NOT EXISTS `order_com` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `comm_name` varchar(50) NOT NULL COMMENT '商品名',
  `user_name` varchar(50) NOT NULL,
  `count_c` int(11) NOT NULL,
  `totle_c` int(11) NOT NULL,
  `is_commented` varchar(10) NOT NULL DEFAULT '未评价',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- 正在导出表  flower_shop.order_com 的数据：~3 rows (大约)
DELETE FROM `order_com`;
/*!40000 ALTER TABLE `order_com` DISABLE KEYS */;
INSERT INTO `order_com` (`id`, `comm_name`, `user_name`, `count_c`, `totle_c`, `is_commented`) VALUES
	(2, '菊花1', 'test1', 5, 50, '已评价'),
	(3, '菊花1', 'test1', 2, 20, '已评价'),
	(4, '菊花1', 'test1', 6, 60, '已评价');
/*!40000 ALTER TABLE `order_com` ENABLE KEYS */;

-- 导出  表 flower_shop.user 结构
CREATE TABLE IF NOT EXISTS `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uname` varchar(50) NOT NULL,
  `pwd` varchar(50) NOT NULL,
  `phone` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- 正在导出表  flower_shop.user 的数据：~2 rows (大约)
DELETE FROM `user`;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` (`id`, `uname`, `pwd`, `phone`) VALUES
	(1, 'test123', '123456', '18030131311'),
	(2, 'test1', '123', '13359072312');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
