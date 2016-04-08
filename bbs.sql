/*
Navicat MySQL Data Transfer

Source Server         : hysql
Source Server Version : 50627
Source Host           : localhost:3306
Source Database       : bbs

Target Server Type    : MYSQL
Target Server Version : 50627
File Encoding         : 65001

Date: 2016-04-08 18:59:38
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `blog`
-- ----------------------------
DROP TABLE IF EXISTS `blog`;
CREATE TABLE `blog` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `blogname` varchar(255) NOT NULL,
  `content` longtext NOT NULL,
  `writername` varchar(255) NOT NULL,
  `submittime` datetime NOT NULL,
  `updatetime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  `tagid` int(11) NOT NULL DEFAULT '1',
  `month` varchar(255) NOT NULL,
  `year` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of blog
-- ----------------------------
INSERT INTO `blog` VALUES ('38', '朱阿姨', '<p>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;这里写你的博； ‘ &nbsp;; &#39; &quot; cfwe客内容\n &nbsp; &nbsp; &nbsp; &nbsp;</p>', 'huangyang', '2016-04-06 17:01:56', '2016-04-06 18:08:24', '9', '', '');
INSERT INTO `blog` VALUES ('39', '朱阿姨', '<p>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;这里写你的博； ‘ &nbsp;; &#39; &quot; cfwe客内容gtrg &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;</p>', 'huangyang', '2016-04-06 17:33:31', '2016-04-06 18:38:07', '10', '', '');
INSERT INTO `blog` VALUES ('40', '哈哈哈哈', '<p>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;这里写你“”的博客</p><p style=\"text-align:center\">fewfcewrweerfregwefwefdscewc</p><p>内容\r\n &nbsp; &nbsp; &nbsp; &nbsp;<br/></p>', 'huangyang', '2016-04-06 17:57:29', '2016-04-08 15:53:27', '1', '', '');
INSERT INTO `blog` VALUES ('41', 'wef', '<p>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;这里写你的博客<img src=\"http://localhost:4001/upload/scrawl/562e2014-fbe3-11e5-8273-989096bbd3b2.png\" alt=\"562e2014-fbe3-11e5-8273-989096bbd3b2.png\"/>内容\n &nbsp; &nbsp; &nbsp; &nbsp;</p>', 'huangyang', '2016-04-06 18:35:54', '2016-04-06 18:35:54', '12', '', '');
INSERT INTO `blog` VALUES ('42', 'dududu', '<p>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;这里写你<img src=\"http://localhost:4001/upload/scrawl/1fb53d28-fd5f-11e5-aa11-989096bbd3b2.png\" alt=\"1fb53d28-fd5f-11e5-aa11-989096bbd3b2.png\"/>的博客内容<img src=\"http://localhost:4001/upload/scrawl/127758a8-fd5f-11e5-aa11-989096bbd3b2.png\" alt=\"127758a8-fd5f-11e5-aa11-989096bbd3b2.png\"/>\r\n &nbsp; &nbsp; &nbsp; &nbsp;</p>', 'huangyang', '2016-04-08 15:54:10', '2016-04-08 15:54:31', '11', '', '');
INSERT INTO `blog` VALUES ('43', 'testtime', '<p>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;这里写你的博客<img src=\"http://localhost:4001/upload/scrawl/0c5a5b18-fd74-11e5-a17c-989096bbd3b2.png\" alt=\"0c5a5b18-fd74-11e5-a17c-989096bbd3b2.png\"/>内容\r\n &nbsp; &nbsp; &nbsp; &nbsp;</p>', 'huangyang', '2016-04-08 18:24:18', '2016-04-08 18:24:18', '1', 'April', '2016年');

-- ----------------------------
-- Table structure for `comment`
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `commenterid` int(11) NOT NULL,
  `blogid` int(11) NOT NULL,
  `superid` int(11) DEFAULT NULL,
  `content` varchar(255) DEFAULT NULL,
  `cmttime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES ('1', '2', '19', '-1', '这是哪里啊！', '2015-12-02 11:19:44');
INSERT INTO `comment` VALUES ('2', '2', '26', '-1', '第一条评论', '2015-12-09 16:33:49');
INSERT INTO `comment` VALUES ('3', '2', '26', '-1', '第二条评论', '2015-12-09 16:34:07');
INSERT INTO `comment` VALUES ('4', '2', '26', '-1', '第三', '2015-12-09 16:34:21');

-- ----------------------------
-- Table structure for `itemtag`
-- ----------------------------
DROP TABLE IF EXISTS `itemtag`;
CREATE TABLE `itemtag` (
  `tagid` int(11) unsigned zerofill NOT NULL AUTO_INCREMENT,
  `tagname` varchar(48) NOT NULL,
  `tagdescribe` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`tagid`,`tagname`),
  UNIQUE KEY `tagname` (`tagname`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of itemtag
-- ----------------------------
INSERT INTO `itemtag` VALUES ('00000000001', '生活', '生活摘要');
INSERT INTO `itemtag` VALUES ('00000000002', '技术', '技术笔记');
INSERT INTO `itemtag` VALUES ('00000000009', '银魂', '寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无');
INSERT INTO `itemtag` VALUES ('00000000010', '超级寿限无', '超级寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无寿限无');
INSERT INTO `itemtag` VALUES ('00000000011', 'gintama', 'fwfowef');
INSERT INTO `itemtag` VALUES ('00000000012', 'gintama2', 'fwfowef');
INSERT INTO `itemtag` VALUES ('00000000015', '技术2', '违反');

-- ----------------------------
-- Table structure for `usermsg`
-- ----------------------------
DROP TABLE IF EXISTS `usermsg`;
CREATE TABLE `usermsg` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `sex` varchar(255) NOT NULL,
  `email` varchar(255) DEFAULT NULL,
  `authority` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of usermsg
-- ----------------------------
INSERT INTO `usermsg` VALUES ('1', 'huangyang', '123', '男', '503582@we.com', '会员');
INSERT INTO `usermsg` VALUES ('2', '黄扬', '123', '男', 'fef@weaf.cn', '管理员');
INSERT INTO `usermsg` VALUES ('3', 'li', '123', '', '', '会员');
INSERT INTO `usermsg` VALUES ('4', 'liu', '123', '', '', '会员');
INSERT INTO `usermsg` VALUES ('5', '5', 'admin', '', '', '会员');
