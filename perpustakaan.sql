-- MySQL dump 10.13  Distrib 8.0.46, for Linux (x86_64)
--
-- Host: localhost    Database: perpustakaan
-- ------------------------------------------------------
-- Server version	8.0.46-0ubuntu0.24.04.3

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `anggota`
--

DROP TABLE IF EXISTS `anggota`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `anggota` (
  `id` int NOT NULL AUTO_INCREMENT,
  `no_anggota` int DEFAULT NULL,
  `nama` varchar(200) DEFAULT NULL,
  `klasifikasi_anggota_id` int DEFAULT NULL,
  `alumni` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `no_anggota` (`no_anggota`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `anggota`
--

LOCK TABLES `anggota` WRITE;
/*!40000 ALTER TABLE `anggota` DISABLE KEYS */;
INSERT INTO `anggota` VALUES (1,1,'naji',1,0),(3,2,'hanzo',1,1),(5,3,'majdi',2,0),(6,4,'mila',3,0);
/*!40000 ALTER TABLE `anggota` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `buku`
--

DROP TABLE IF EXISTS `buku`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `buku` (
  `id` int NOT NULL AUTO_INCREMENT,
  `judul` text,
  `list_kategori_id` int DEFAULT NULL,
  `stock` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `buku`
--

LOCK TABLES `buku` WRITE;
/*!40000 ALTER TABLE `buku` DISABLE KEYS */;
INSERT INTO `buku` VALUES (2,'BUMI MANUSIA',2,5);
/*!40000 ALTER TABLE `buku` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `buku_hub`
--

DROP TABLE IF EXISTS `buku_hub`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `buku_hub` (
  `id` int NOT NULL AUTO_INCREMENT,
  `barcode` varchar(50) DEFAULT NULL,
  `buku_id` int DEFAULT NULL,
  `list_kondisi_id` int DEFAULT NULL,
  `anggota_id` int DEFAULT NULL,
  `rak_id` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `buku_hub`
--

LOCK TABLES `buku_hub` WRITE;
/*!40000 ALTER TABLE `buku_hub` DISABLE KEYS */;
INSERT INTO `buku_hub` VALUES (1,'f0001',1,2,NULL,3),(2,'h0001',2,1,NULL,2),(3,'m0001',4,3,NULL,1),(4,'f0002',1,2,NULL,3),(5,'f0003',1,2,NULL,3),(6,'f0004',1,2,NULL,3),(7,'f0005',1,2,NULL,3);
/*!40000 ALTER TABLE `buku_hub` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `denda`
--

DROP TABLE IF EXISTS `denda`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `denda` (
  `id` int NOT NULL AUTO_INCREMENT,
  `petugas_id` int DEFAULT NULL,
  `buku_id` int DEFAULT NULL,
  `list_kondisi_id` int DEFAULT NULL,
  `total_hari` int DEFAULT NULL,
  `denda` int DEFAULT NULL,
  `harga_buku_id` int DEFAULT NULL,
  `total_denda` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `denda`
--

LOCK TABLES `denda` WRITE;
/*!40000 ALTER TABLE `denda` DISABLE KEYS */;
/*!40000 ALTER TABLE `denda` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `harga_buku`
--

DROP TABLE IF EXISTS `harga_buku`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `harga_buku` (
  `id` int NOT NULL AUTO_INCREMENT,
  `buku_id` int DEFAULT NULL,
  `harga` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `harga_buku`
--

LOCK TABLES `harga_buku` WRITE;
/*!40000 ALTER TABLE `harga_buku` DISABLE KEYS */;
/*!40000 ALTER TABLE `harga_buku` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `klasifikasi_anggota`
--

DROP TABLE IF EXISTS `klasifikasi_anggota`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `klasifikasi_anggota` (
  `id` int NOT NULL AUTO_INCREMENT,
  `klasifikasi` char(50) DEFAULT NULL,
  `maks_buku` int DEFAULT NULL,
  `maks_hari` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `klasifikasi_anggota`
--

LOCK TABLES `klasifikasi_anggota` WRITE;
/*!40000 ALTER TABLE `klasifikasi_anggota` DISABLE KEYS */;
INSERT INTO `klasifikasi_anggota` VALUES (1,'siswa',3,7),(2,'guru',5,14),(3,'staf TU',2,7);
/*!40000 ALTER TABLE `klasifikasi_anggota` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `list_kategori`
--

DROP TABLE IF EXISTS `list_kategori`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `list_kategori` (
  `id` int NOT NULL AUTO_INCREMENT,
  `kategori` varchar(300) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `list_kategori`
--

LOCK TABLES `list_kategori` WRITE;
/*!40000 ALTER TABLE `list_kategori` DISABLE KEYS */;
INSERT INTO `list_kategori` VALUES (1,'horror'),(2,'fantasy'),(3,'romance'),(4,'mystery'),(5,'comedy');
/*!40000 ALTER TABLE `list_kategori` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `list_kondisi`
--

DROP TABLE IF EXISTS `list_kondisi`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `list_kondisi` (
  `id` int NOT NULL AUTO_INCREMENT,
  `kondisi` char(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `list_kondisi`
--

LOCK TABLES `list_kondisi` WRITE;
/*!40000 ALTER TABLE `list_kondisi` DISABLE KEYS */;
INSERT INTO `list_kondisi` VALUES (1,'baik'),(2,'rusak ringan'),(3,'rusak berat'),(4,'hilang');
/*!40000 ALTER TABLE `list_kondisi` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `petugas`
--

DROP TABLE IF EXISTS `petugas`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `petugas` (
  `id` int NOT NULL AUTO_INCREMENT,
  `anggota_id` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `petugas`
--

LOCK TABLES `petugas` WRITE;
/*!40000 ALTER TABLE `petugas` DISABLE KEYS */;
INSERT INTO `petugas` VALUES (1,6),(2,5);
/*!40000 ALTER TABLE `petugas` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pinjaman`
--

DROP TABLE IF EXISTS `pinjaman`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pinjaman` (
  `id` int NOT NULL AUTO_INCREMENT,
  `anggota_id` int DEFAULT NULL,
  `buku_id` int DEFAULT NULL,
  `tgl_pinjaman` date DEFAULT NULL,
  `tgl_balik` date DEFAULT NULL,
  `petugas_pinajam_id` int DEFAULT NULL,
  `list_kondisi_id` int DEFAULT NULL,
  `petugas_balik_id` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pinjaman`
--

LOCK TABLES `pinjaman` WRITE;
/*!40000 ALTER TABLE `pinjaman` DISABLE KEYS */;
/*!40000 ALTER TABLE `pinjaman` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rak`
--

DROP TABLE IF EXISTS `rak`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `rak` (
  `id` int NOT NULL AUTO_INCREMENT,
  `no_rak` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `no_rak` (`no_rak`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rak`
--

LOCK TABLES `rak` WRITE;
/*!40000 ALTER TABLE `rak` DISABLE KEYS */;
INSERT INTO `rak` VALUES (1,1),(2,2),(3,3);
/*!40000 ALTER TABLE `rak` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `reservasi`
--

DROP TABLE IF EXISTS `reservasi`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `reservasi` (
  `id` int NOT NULL AUTO_INCREMENT,
  `anggota_id` int DEFAULT NULL,
  `buku_id` int DEFAULT NULL,
  `batas_pengambilan` date DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reservasi`
--

LOCK TABLES `reservasi` WRITE;
/*!40000 ALTER TABLE `reservasi` DISABLE KEYS */;
/*!40000 ALTER TABLE `reservasi` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-07-13 18:01:07
