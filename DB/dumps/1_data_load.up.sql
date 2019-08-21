-- MySQL dump 10.13  Distrib 8.0.17, for macos10.14 (x86_64)
--
-- Host: localhost    Database: Restaurant
-- ------------------------------------------------------
-- Server version	8.0.16

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `Admin`
--

DROP TABLE IF EXISTS `Admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Admin` (
  `Name` varchar(20) NOT NULL,
  `Pass` varchar(100) DEFAULT NULL,
  `Email` varchar(50) NOT NULL,
  `Adder` int(11) DEFAULT NULL,
  `AdderRole` int(1) DEFAULT NULL,
  `ID` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `Status` int(1) DEFAULT '1',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `Email_UNIQUE` (`Email`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Admin`
--

LOCK TABLES `Admin` WRITE;
/*!40000 ALTER TABLE `Admin` DISABLE KEYS */;
INSERT INTO `Admin` VALUES ('admin1','$2a$04$AKe7E84SCtQZCjPYpRCN6OrdMS/VnV0Qx93Os.TU8nO71Tt67ysmG','admin1@gmail.com',1,2,1,1,'2019-08-08 15:47:35'),('admin2','$2a$05$qcrRmI4s9tU6UGxQHlb0O.y6cN4j2Ire.dj/fe5aKFiYaGCPWPk2G','admin2@gmail.com',1,2,5,1,'2019-08-09 13:42:44');
/*!40000 ALTER TABLE `Admin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DeletedTokens`
--

DROP TABLE IF EXISTS `DeletedTokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DeletedTokens` (
  `token` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DeletedTokens`
--

LOCK TABLES `DeletedTokens` WRITE;
/*!40000 ALTER TABLE `DeletedTokens` DISABLE KEYS */;
INSERT INTO `DeletedTokens` VALUES ('eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXI0QGdtYWlsLmNvbSIsImV4cCI6MTU2NTM1MDA0NiwiaWQiOjE3LCJuYW1lIjoidXNlcjQiLCJyYW5rIjowfQ.2okVxjRJ9V_3hyumBwaVdW94M6laruppfQvrNSseZ6w');
/*!40000 ALTER TABLE `DeletedTokens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Dish`
--

DROP TABLE IF EXISTS `Dish`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Dish` (
  `Name` varchar(50) NOT NULL,
  `Price` int(11) DEFAULT NULL,
  `RID` int(11) NOT NULL,
  `Adder` int(11) DEFAULT NULL,
  `AdderRole` int(1) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Status` int(1) DEFAULT '1',
  `Created_At` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Dish`
--

LOCK TABLES `Dish` WRITE;
/*!40000 ALTER TABLE `Dish` DISABLE KEYS */;
INSERT INTO `Dish` VALUES ('burger',100,14,1,1,5,1,'2019-08-08 17:16:25'),('chocoshake',100,14,1,1,6,1,'2019-08-08 17:17:47'),('fried chickens',100,14,1,1,7,1,'2019-08-08 17:17:59'),('fried chickens',100,13,1,1,8,1,'2019-08-08 17:18:32'),('garlic bread',100,13,1,1,9,1,'2019-08-08 17:18:45'),('vanilla shake',100,13,1,1,10,0,'2019-08-08 17:18:54'),('chicken burger',100,20,1,2,11,1,'2019-08-09 14:44:51');
/*!40000 ALTER TABLE `Dish` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Rest`
--

DROP TABLE IF EXISTS `Rest`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Rest` (
  `Name` varchar(50) NOT NULL,
  `latitude` varchar(10) NOT NULL,
  `longitude` varchar(10) NOT NULL,
  `Owner` varchar(50) DEFAULT NULL,
  `Adder` int(11) DEFAULT NULL,
  `AdderRole` int(1) DEFAULT NULL,
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Status` int(1) DEFAULT '1',
  `Created_At` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `RestId_UNIQUE` (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Rest`
--

LOCK TABLES `Rest` WRITE;
/*!40000 ALTER TABLE `Rest` DISABLE KEYS */;
INSERT INTO `Rest` VALUES ('McDonalds','125.287','123.223','user1@gmail.com',1,1,13,1,'2019-08-08 16:33:26'),('KFC','125','125','user1@gmail.com',1,1,14,1,'2019-08-08 16:34:00'),('CCD','12','14','user1@gmail.com',1,2,15,1,'2019-08-08 16:35:21'),('Starbucks','56.2','76','user1@gmail.com',1,2,16,1,'2019-08-08 16:35:52'),('Mainland China','72.123','78.987','user4@gmail.com',1,1,17,1,'2019-08-09 14:09:41'),('Oh Calcutta','125.123','23.987','user2@gmail.com',5,1,18,1,'2019-08-09 14:14:40'),('The Wall','12.123','123.987','user2@gmail.com',5,1,19,0,'2019-08-09 14:15:49'),('Burger Bae','102.123','43.987','user4@gmail.com',5,1,20,1,'2019-08-09 14:16:30');
/*!40000 ALTER TABLE `Rest` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `SuperAdmin`
--

DROP TABLE IF EXISTS `SuperAdmin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `SuperAdmin` (
  `Name` varchar(20) NOT NULL,
  `Pass` varchar(100) DEFAULT NULL,
  `Email` varchar(50) NOT NULL,
  `Adder` int(11) DEFAULT NULL,
  `AdderRole` int(1) DEFAULT NULL,
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Status` int(1) DEFAULT '1',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `Email_UNIQUE` (`Email`),
  UNIQUE KEY `id_UNIQUE` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `SuperAdmin`
--

LOCK TABLES `SuperAdmin` WRITE;
/*!40000 ALTER TABLE `SuperAdmin` DISABLE KEYS */;
INSERT INTO `SuperAdmin` VALUES ('Sourav','$2a$04$CAIRyD6NVptXbB25PEUFJeYEsBwIouUYLigBhWocbcmvOZrc7.OV.','sourav241196@gmail.com',0,2,1,1,'2019-08-08 14:14:10');
/*!40000 ALTER TABLE `SuperAdmin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `User`
--

DROP TABLE IF EXISTS `User`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `User` (
  `Name` varchar(20) NOT NULL,
  `PassHash` varchar(100) DEFAULT NULL,
  `email` varchar(50) NOT NULL,
  `Adder` int(11) DEFAULT NULL,
  `AdderRole` int(1) DEFAULT NULL,
  `ID` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `Status` int(1) DEFAULT '1',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `User`
--

LOCK TABLES `User` WRITE;
/*!40000 ALTER TABLE `User` DISABLE KEYS */;
INSERT INTO `User` VALUES ('user2','$2a$04$m9YYYPnPwJljNVRGNTHjru5WVj7rEg18eOxHRJH/8cKJvtTDq773W','user2@gmail.com',1,2,14,1,'2019-08-08 15:16:29'),('user4','$2a$04$qViSgQavHTsvE36TsgxbOOmpoXunLe3WJaFYbb6DdZ3mhIdsWrBuC','user4@gmail.com',1,2,17,1,'2019-08-09 13:38:17'),('user1','$2a$04$y9ExoM60HYRfNxm8N/tE3.YHVS/RhHB/6eaztdwVYhoRPspofsmk2','user1@gmail.com',1,1,18,1,'2019-08-09 14:12:25');
/*!40000 ALTER TABLE `User` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'Restaurant'
--

--
-- Dumping routines for database 'Restaurant'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-08-21 19:23:42
