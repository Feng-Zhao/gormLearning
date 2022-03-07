/*
 Navicat MySQL Data Transfer

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 07/03/2022 17:55:04
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for book
-- ----------------------------
DROP TABLE IF EXISTS `book`;
CREATE TABLE `book` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `belong_sid` int unsigned NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `modify_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of book
-- ----------------------------
BEGIN;
INSERT INTO `book` VALUES (1, 'book1', 2022010101, 'a test book', '2022-02-15 14:15:16', '2022-02-15 14:33:51');
COMMIT;

-- ----------------------------
-- Table structure for class
-- ----------------------------
DROP TABLE IF EXISTS `class`;
CREATE TABLE `class` (
  `id` int NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `teacher_id` int unsigned NOT NULL,
  `teacher_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of class
-- ----------------------------
BEGIN;
INSERT INTO `class` VALUES (20220101, '22级1年1班', 1, 't_1');
COMMIT;

-- ----------------------------
-- Table structure for student
-- ----------------------------
DROP TABLE IF EXISTS `student`;
CREATE TABLE `student` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `class_id` int unsigned DEFAULT NULL,
  `class_name` varchar(255) DEFAULT NULL,
  `teacher_id` int unsigned DEFAULT NULL,
  `teacher_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2022010103 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of student
-- ----------------------------
BEGIN;
INSERT INTO `student` VALUES (2022010101, 's1', 20220101, '22-1-1', 1, 't_1');
INSERT INTO `student` VALUES (2022010102, 's2', 20220102, '22-1-2', 2, 't_2');
COMMIT;

-- ----------------------------
-- Table structure for teacher
-- ----------------------------
DROP TABLE IF EXISTS `teacher`;
CREATE TABLE `teacher` (
  `id` int unsigned NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of teacher
-- ----------------------------
BEGIN;
INSERT INTO `teacher` VALUES (1, 't_1');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
