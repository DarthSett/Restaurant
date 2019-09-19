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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

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
