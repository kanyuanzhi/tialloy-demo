/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : web_service

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 11/12/2021 13:26:03
*/

SET NAMES utf8mb4;
SET
    FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for terminals
-- ----------------------------
DROP TABLE IF EXISTS `terminals`;
CREATE TABLE `terminals`
(
    `id`                    int UNSIGNED                                                  NOT NULL AUTO_INCREMENT,
    `name`                  varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
    `manager`               varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
    `host_hostname`         varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
    `host_boot_time`        bigint UNSIGNED                                               NULL DEFAULT NULL,
    `host_os`               varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
    `host_platform`         varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
    `host_platform_family`  varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
    `host_platform_version` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
    `host_kernel_version`   varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
    `host_kernel_arch`      varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
    `host_user`             varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
    `cpu_model_name`        varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
    `cpu_physical_cores`    int                                                           NULL DEFAULT NULL,
    `cpu_logical_cores`     int                                                           NULL DEFAULT NULL,
    `mem_total`             bigint UNSIGNED                                               NULL DEFAULT NULL,
    `net_name`              varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
    `net_ip`                varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
    `net_mac`               varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
    `disk_path`             varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
    `disk_fstype`           varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
    `disk_total`            bigint UNSIGNED                                               NULL DEFAULT NULL,
    `created_at`            bigint                                                        NULL DEFAULT NULL,
    `updated_at`            bigint                                                        NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = Dynamic;

SET
    FOREIGN_KEY_CHECKS = 1;
