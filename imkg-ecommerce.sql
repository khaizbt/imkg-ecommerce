-- MySQL dump 10.13  Distrib 8.0.28, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: imkg-ecommerce
-- ------------------------------------------------------
-- Server version	5.7.37

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
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `categories` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `slug` longtext,
  `image` longtext,
  `created_by` bigint(20) DEFAULT NULL,
  `updated_by` bigint(20) DEFAULT NULL,
  `deleted_by` bigint(20) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES (1,'Category 1','category-1','Storage/avatar/category-1.png',1,NULL,NULL,NULL,'2021-02-27 22:20:26.689','2021-02-27 22:20:26.689');
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cities`
--

DROP TABLE IF EXISTS `cities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cities` (
  `id_city` varchar(30) COLLATE latin1_general_ci NOT NULL,
  `id_province` varchar(30) COLLATE latin1_general_ci NOT NULL,
  `city_code` varchar(10) COLLATE latin1_general_ci DEFAULT NULL,
  `city_name` varchar(100) COLLATE latin1_general_ci NOT NULL,
  `city_type` varchar(10) COLLATE latin1_general_ci NOT NULL,
  `city_postal_code` varchar(5) COLLATE latin1_general_ci DEFAULT NULL,
  `is_active` tinyint(4) DEFAULT '1',
  `is_used` tinyint(4) DEFAULT '0',
  `created_by` varchar(30) COLLATE latin1_general_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_by` varchar(30) COLLATE latin1_general_ci DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id_city`),
  KEY `id_province` (`id_province`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cities`
--

LOCK TABLES `cities` WRITE;
/*!40000 ALTER TABLE `cities` DISABLE KEYS */;
INSERT INTO `cities` VALUES ('1101','11',NULL,'Simeulue','Kabupaten','23891',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1102','11',NULL,'Aceh Singkil','Kabupaten','24785',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1103','11',NULL,'Aceh Selatan','Kabupaten','23719',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1104','11',NULL,'Aceh Tenggara','Kabupaten','24611',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1105','11',NULL,'Aceh Timur','Kabupaten','24454',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1106','11',NULL,'Aceh Tengah','Kabupaten','24511',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1107','11',NULL,'Aceh Barat','Kabupaten','23681',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1108','11',NULL,'Aceh Besar','Kabupaten','23951',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1109','11',NULL,'Pidie','Kabupaten','24116',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1110','11',NULL,'Bireuen','Kabupaten','24219',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1111','11',NULL,'Aceh Utara','Kabupaten','24382',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1112','11',NULL,'Aceh Barat Daya','Kabupaten','23764',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1113','11',NULL,'Gayo Lues','Kabupaten','24653',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1114','11',NULL,'Aceh Tamiang','Kabupaten','24476',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1115','11',NULL,'Nagan Raya','Kabupaten','23674',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1116','11',NULL,'Aceh Jaya','Kabupaten','23654',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1117','11',NULL,'Bener Meriah','Kabupaten','24581',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1118','11',NULL,'Pidie Jaya','Kabupaten','24186',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1171','11',NULL,'Banda Aceh','Kota','23238',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1172','11',NULL,'Sabang','Kota','23512',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1173','11',NULL,'Langsa','Kota','24412',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1174','11',NULL,'Lhokseumawe','Kota','24352',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1175','11',NULL,'Subulussalam','Kota','24882',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1201','12',NULL,'Nias','Kabupaten','22876',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1202','12',NULL,'Mandailing Natal','Kabupaten','22916',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1203','12',NULL,'Tapanuli Selatan','Kabupaten','22742',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1204','12',NULL,'Tapanuli Tengah','Kabupaten','22611',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1205','12',NULL,'Tapanuli Utara','Kabupaten','22414',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1206','12',NULL,'Toba Samosir','Kabupaten','22316',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1207','12',NULL,'Labuhan Batu','Kabupaten','21412',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1208','12',NULL,'Asahan','Kabupaten','21214',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1209','12',NULL,'Simalungun','Kabupaten','21162',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1210','12',NULL,'Dairi','Kabupaten','22211',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1211','12',NULL,'Karo','Kabupaten','22119',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1212','12',NULL,'Deli Serdang','Kabupaten','20511',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1213','12',NULL,'Langkat','Kabupaten','20811',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1214','12',NULL,'Nias Selatan','Kabupaten','22865',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1215','12',NULL,'Humbang Hasundutan','Kabupaten','22457',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1216','12',NULL,'Pakpak Bharat','Kabupaten','22272',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1217','12',NULL,'Samosir','Kabupaten','22392',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1218','12',NULL,'Serdang Bedagai','Kabupaten','20915',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1219','12',NULL,'Batu Bara','Kabupaten','21655',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1220','12',NULL,'Padang Lawas Utara','Kabupaten','22753',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1221','12',NULL,'Padang Lawas','Kabupaten','22763',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1222','12',NULL,'Labuhan Batu Selatan','Kabupaten','21511',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1223','12',NULL,'Labuhan Batu Utara','Kabupaten','21711',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1224','12',NULL,'Nias Utara','Kabupaten','22856',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1225','12',NULL,'Nias Barat','Kabupaten','22895',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1271','12',NULL,'Sibolga','Kota','22522',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1272','12',NULL,'Tanjung Balai','Kota','21321',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1273','12',NULL,'Pematang Siantar','Kota','21126',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1274','12',NULL,'Tebing Tinggi','Kota','20632',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1275','12',NULL,'Medan','Kota','20228',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1276','12',NULL,'Binjai','Kota','20712',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1277','12',NULL,'Padang Sidempuan','Kota','22727',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1278','12',NULL,'Gunungsitoli','Kota','22813',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1301','13',NULL,'Kepulauan Mentawai','Kabupaten','25771',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1302','13',NULL,'Pesisir Selatan','Kabupaten','25611',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1303','13',NULL,'Solok','Kabupaten','27365',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1304','13',NULL,'Sijunjung (Sawah Lunto Sijunjung)','Kabupaten','27511',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1305','13',NULL,'Tanah Datar','Kabupaten','27211',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1306','13',NULL,'Padang Pariaman','Kabupaten','25583',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1307','13',NULL,'Agam','Kabupaten','26411',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1308','13',NULL,'Lima Puluh Koto/Kota','Kabupaten','26671',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1309','13',NULL,'Pasaman','Kabupaten','26318',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1310','13',NULL,'Solok Selatan','Kabupaten','27779',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1311','13',NULL,'Dharmasraya','Kabupaten','27612',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1312','13',NULL,'Pasaman Barat','Kabupaten','26511',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1371','13',NULL,'Padang','Kota','25112',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1372','13',NULL,'Solok','Kota','27315',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1373','13',NULL,'Sawah Lunto','Kota','27416',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1374','13',NULL,'Padang Panjang','Kota','27122',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1375','13',NULL,'Bukittinggi','Kota','26115',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1376','13',NULL,'Payakumbuh','Kota','26213',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1377','13',NULL,'Pariaman','Kota','25511',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1401','14',NULL,'Kuantan Singingi','Kabupaten','29519',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1402','14',NULL,'Indragiri Hulu','Kabupaten','29319',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1403','14',NULL,'Indragiri Hilir','Kabupaten','29212',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1404','14',NULL,'Pelalawan','Kabupaten','28311',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1406','14',NULL,'Kampar','Kabupaten','28411',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1407','14',NULL,'Rokan Hulu','Kabupaten','28511',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1408','14',NULL,'Bengkalis','Kabupaten','28719',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1409','14',NULL,'Rokan Hilir','Kabupaten','28992',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1410','14',NULL,'Kepulauan Meranti','Kabupaten','28791',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1471','14',NULL,'Pekanbaru','Kota','28112',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1473','14',NULL,'Dumai','Kota','28811',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1501','15',NULL,'Kerinci','Kabupaten','37167',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1502','15',NULL,'Merangin','Kabupaten','37319',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1503','15',NULL,'Sarolangun','Kabupaten','37419',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1504','15',NULL,'Batang Hari','Kabupaten','36613',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1505','15',NULL,'Muaro Jambi','Kabupaten','36311',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1506','15',NULL,'Tanjung Jabung Timur','Kabupaten','36719',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1507','15',NULL,'Tanjung Jabung Barat','Kabupaten','36513',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1508','15',NULL,'Tebo','Kabupaten','37519',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1509','15',NULL,'Bungo','Kabupaten','37216',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1571','15',NULL,'Jambi','Kota','36111',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1572','15',NULL,'Sungai Penuh','Kota','37113',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1601','16',NULL,'Ogan Komering Ulu','Kabupaten','32112',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1602','16',NULL,'Ogan Komering Ilir','Kabupaten','30618',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1603','16',NULL,'Muara Enim','Kabupaten','31315',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1604','16',NULL,'Lahat','Kabupaten','31419',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1605','16',NULL,'Musi Rawas','Kabupaten','31661',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1606','16',NULL,'Musi Banyuasin','Kabupaten','30719',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1607','16',NULL,'Banyuasin','Kabupaten','30911',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1608','16',NULL,'Ogan Komering Ulu Selatan','Kabupaten','32211',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1609','16',NULL,'Ogan Komering Ulu Timur','Kabupaten','32312',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1610','16',NULL,'Ogan Ilir','Kabupaten','30811',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1611','16',NULL,'Empat Lawang','Kabupaten','31811',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1671','16',NULL,'Palembang','Kota','31512',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1672','16',NULL,'Prabumulih','Kota','31121',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1673','16',NULL,'Pagar Alam','Kota','31512',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1674','16',NULL,'Lubuk Linggau','Kota','31614',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1701','17',NULL,'Bengkulu Selatan','Kabupaten','38519',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1702','17',NULL,'Rejang Lebong','Kabupaten','39112',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1703','17',NULL,'Bengkulu Utara','Kabupaten','38619',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1704','17',NULL,'Kaur','Kabupaten','38911',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1705','17',NULL,'Seluma','Kabupaten','38811',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1706','17',NULL,'Muko Muko','Kabupaten','38715',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1707','17',NULL,'Lebong','Kabupaten','39264',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1708','17',NULL,'Kepahiang','Kabupaten','39319',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1709','17',NULL,'Bengkulu Tengah','Kabupaten','38319',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1771','17',NULL,'Bengkulu','Kota','38229',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1801','18',NULL,'Lampung Barat','Kabupaten','34814',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1802','18',NULL,'Tanggamus','Kabupaten','35619',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1803','18',NULL,'Lampung Selatan','Kabupaten','35511',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1804','18',NULL,'Lampung Timur','Kabupaten','34319',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1805','18',NULL,'Lampung Tengah','Kabupaten','34212',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1806','18',NULL,'Lampung Utara','Kabupaten','34516',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1807','18',NULL,'Way Kanan','Kabupaten','34711',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1808','18',NULL,'Tulang Bawang','Kabupaten','34613',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1809','18',NULL,'Pesawaran','Kabupaten','35312',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1810','18',NULL,'Pringsewu','Kabupaten','35719',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1811','18',NULL,'Mesuji','Kabupaten','34911',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1812','18',NULL,'Tulang Bawang Barat','Kabupaten','34419',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1813','18',NULL,'Pesisir Barat','Kabupaten','35974',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1871','18',NULL,'Bandar Lampung','Kota','35139',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1872','18',NULL,'Metro','Kota','34111',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1901','19',NULL,'Bangka','Kabupaten','33212',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1902','19',NULL,'Belitung','Kabupaten','33419',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1903','19',NULL,'Bangka Barat','Kabupaten','33315',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1904','19',NULL,'Bangka Tengah','Kabupaten','33613',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1905','19',NULL,'Bangka Selatan','Kabupaten','33719',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1906','19',NULL,'Belitung Timur','Kabupaten','33519',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('1971','19',NULL,'Pangkal Pinang','Kota','33115',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('2101','21',NULL,'Karimun','Kabupaten','29611',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('2102','21',NULL,'Bintan','Kabupaten','29135',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('2103','21',NULL,'Natuna','Kabupaten','29711',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('2104','21',NULL,'Lingga','Kabupaten','29811',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('2105','21',NULL,'Kepulauan Anambas','Kabupaten','29991',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('2171','21',NULL,'Batam','Kota','29413',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('2172','21',NULL,'Tanjung Pinang','Kota','29111',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3101','31',NULL,'Kepulauan Seribu','Kabupaten','14550',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3171','31',NULL,'Jakarta Selatan','Kota','12230',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3172','31',NULL,'Jakarta Timur','Kota','13330',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3173','31',NULL,'Jakarta Pusat','Kota','10540',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3174','31',NULL,'Jakarta Barat','Kota','11220',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3175','31',NULL,'Jakarta Utara','Kota','14140',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3201','32',NULL,'Bogor','Kabupaten','16911',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3202','32',NULL,'Sukabumi','Kabupaten','43311',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3203','32',NULL,'Cianjur','Kabupaten','43217',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3204','32',NULL,'Bandung','Kabupaten','40311',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3205','32',NULL,'Garut','Kabupaten','44126',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3206','32',NULL,'Tasikmalaya','Kabupaten','46411',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3207','32',NULL,'Ciamis','Kabupaten','46211',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3208','32',NULL,'Kuningan','Kabupaten','45511',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3209','32',NULL,'Cirebon','Kabupaten','45611',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3210','32',NULL,'Majalengka','Kabupaten','45412',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3211','32',NULL,'Sumedang','Kabupaten','45326',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3212','32',NULL,'Indramayu','Kabupaten','45214',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3213','32',NULL,'Subang','Kabupaten','41215',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3214','32',NULL,'Purwakarta','Kabupaten','41119',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3215','32',NULL,'Karawang','Kabupaten','41311',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3216','32',NULL,'Bekasi','Kabupaten','17837',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3217','32',NULL,'Bandung Barat','Kabupaten','40721',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3218','32',NULL,'Pangandaran','Kabupaten','46511',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3271','32',NULL,'Bogor','Kota','16119',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3272','32',NULL,'Sukabumi','Kota','43114',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3273','32',NULL,'Bandung','Kota','40111',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3274','32',NULL,'Cirebon','Kota','45116',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3275','32',NULL,'Bekasi','Kota','17121',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3276','32',NULL,'Depok','Kota','16416',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3277','32',NULL,'Cimahi','Kota','40512',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3278','32',NULL,'Tasikmalaya','Kota','46116',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3279','32',NULL,'Banjar','Kota','46311',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3301','33',NULL,'Cilacap','Kabupaten','53211',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3302','33',NULL,'Banyumas','Kabupaten','53114',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3303','33',NULL,'Purbalingga','Kabupaten','53312',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3304','33',NULL,'Banjarnegara','Kabupaten','53419',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3305','33',NULL,'Kebumen','Kabupaten','54319',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3306','33',NULL,'Purworejo','Kabupaten','54111',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3307','33',NULL,'Wonosobo','Kabupaten','56311',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3308','33',NULL,'Magelang','Kabupaten','56519',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3309','33',NULL,'Boyolali','Kabupaten','57312',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3310','33',NULL,'Klaten','Kabupaten','57411',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3311','33',NULL,'Sukoharjo','Kabupaten','57514',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3312','33',NULL,'Wonogiri','Kabupaten','57619',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3313','33',NULL,'Karanganyar','Kabupaten','57718',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3314','33',NULL,'Sragen','Kabupaten','57211',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3315','33',NULL,'Grobogan','Kabupaten','58111',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3316','33',NULL,'Blora','Kabupaten','58219',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3317','33',NULL,'Rembang','Kabupaten','59219',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3318','33',NULL,'Pati','Kabupaten','59114',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3319','33',NULL,'Kudus','Kabupaten','59311',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3320','33',NULL,'Jepara','Kabupaten','59419',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3321','33',NULL,'Demak','Kabupaten','59519',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3322','33',NULL,'Semarang','Kabupaten','50511',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3323','33',NULL,'Temanggung','Kabupaten','56212',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3324','33',NULL,'Kendal','Kabupaten','51314',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3325','33',NULL,'Batang','Kabupaten','51211',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3326','33',NULL,'Pekalongan','Kabupaten','51161',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3327','33',NULL,'Pemalang','Kabupaten','52319',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3328','33',NULL,'Tegal','Kabupaten','52419',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3329','33',NULL,'Brebes','Kabupaten','52212',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3371','33',NULL,'Magelang','Kota','56133',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3372','33',NULL,'Surakarta (Solo)','Kota','57113',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3373','33',NULL,'Salatiga','Kota','50711',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3374','33',NULL,'Semarang','Kota','50135',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3375','33',NULL,'Pekalongan','Kota','51122',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3376','33',NULL,'Tegal','Kota','52114',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3401','34',NULL,'Kulon Progo','Kabupaten','55611',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3402','34',NULL,'Bantul','Kabupaten','55715',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3403','34',NULL,'Gunung Kidul','Kabupaten','55812',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3404','34',NULL,'Sleman','Kabupaten','55513',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3471','34',NULL,'Yogyakarta','Kota','55222',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3501','35',NULL,'Pacitan','Kabupaten','63512',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3502','35',NULL,'Ponorogo','Kabupaten','63411',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3503','35',NULL,'Trenggalek','Kabupaten','66312',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3504','35',NULL,'Tulungagung','Kabupaten','66212',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3505','35',NULL,'Blitar','Kabupaten','66171',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3506','35',NULL,'Kediri','Kabupaten','64184',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3507','35',NULL,'Malang','Kabupaten','65163',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3508','35',NULL,'Lumajang','Kabupaten','67319',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3509','35',NULL,'Jember','Kabupaten','68113',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3510','35',NULL,'Banyuwangi','Kabupaten','68416',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3511','35',NULL,'Bondowoso','Kabupaten','68219',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3512','35',NULL,'Situbondo','Kabupaten','68316',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3513','35',NULL,'Probolinggo','Kabupaten','67282',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3514','35',NULL,'Pasuruan','Kabupaten','67153',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3515','35',NULL,'Sidoarjo','Kabupaten','61219',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3516','35',NULL,'Mojokerto','Kabupaten','61382',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3517','35',NULL,'Jombang','Kabupaten','61415',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3518','35',NULL,'Nganjuk','Kabupaten','64414',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3519','35',NULL,'Madiun','Kabupaten','63153',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3520','35',NULL,'Magetan','Kabupaten','63314',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3521','35',NULL,'Ngawi','Kabupaten','63219',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3522','35',NULL,'Bojonegoro','Kabupaten','62119',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3523','35',NULL,'Tuban','Kabupaten','62319',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3524','35',NULL,'Lamongan','Kabupaten','64125',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3525','35',NULL,'Gresik','Kabupaten','61115',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3526','35',NULL,'Bangkalan','Kabupaten','69118',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3527','35',NULL,'Sampang','Kabupaten','69219',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3528','35',NULL,'Pamekasan','Kabupaten','69319',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3529','35',NULL,'Sumenep','Kabupaten','69413',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3571','35',NULL,'Kediri','Kota','64125',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3572','35',NULL,'Blitar','Kota','66124',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3573','35',NULL,'Malang','Kota','65112',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3574','35',NULL,'Probolinggo','Kota','67215',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3575','35',NULL,'Pasuruan','Kota','67118',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3576','35',NULL,'Mojokerto','Kota','61316',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3577','35',NULL,'Madiun','Kota','63122',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3578','35',NULL,'Surabaya','Kota','60119',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3579','35',NULL,'Batu','Kota','65311',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3601','36',NULL,'Pandeglang','Kabupaten','42212',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3602','36',NULL,'Lebak','Kabupaten','42319',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3603','36',NULL,'Tangerang','Kabupaten','15914',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3604','36',NULL,'Serang','Kabupaten','42182',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3671','36',NULL,'Tangerang','Kota','15111',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3672','36',NULL,'Cilegon','Kota','42417',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3673','36',NULL,'Serang','Kota','42111',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('3674','36',NULL,'Tangerang Selatan','Kota','15332',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5101','51',NULL,'Jembrana','Kabupaten','82251',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5102','51',NULL,'Tabanan','Kabupaten','82119',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5103','51',NULL,'Badung','Kabupaten','80351',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5104','51',NULL,'Gianyar','Kabupaten','80519',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5105','51',NULL,'Klungkung','Kabupaten','80719',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5106','51',NULL,'Bangli','Kabupaten','80619',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5107','51',NULL,'Karangasem','Kabupaten','80819',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5108','51',NULL,'Buleleng','Kabupaten','81111',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5171','51',NULL,'Denpasar','Kota','80227',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5201','52',NULL,'Lombok Barat','Kabupaten','83311',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5202','52',NULL,'Lombok Tengah','Kabupaten','83511',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5203','52',NULL,'Lombok Timur','Kabupaten','83612',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5204','52',NULL,'Sumbawa','Kabupaten','84315',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5205','52',NULL,'Dompu','Kabupaten','84217',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5206','52',NULL,'Bima','Kabupaten','84171',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5207','52',NULL,'Sumbawa Barat','Kabupaten','84419',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5208','52',NULL,'Lombok Utara','Kabupaten','83711',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5271','52',NULL,'Mataram','Kota','83131',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5272','52',NULL,'Bima','Kota','84139',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5301','53',NULL,'Sumba Barat','Kabupaten','87219',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5302','53',NULL,'Sumba Timur','Kabupaten','87112',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5303','53',NULL,'Kupang','Kabupaten','85362',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5304','53',NULL,'Timor Tengah Selatan','Kabupaten','85562',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5305','53',NULL,'Timor Tengah Utara','Kabupaten','85612',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5306','53',NULL,'Belu','Kabupaten','85711',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5307','53',NULL,'Alor','Kabupaten','85811',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5308','53',NULL,'Lembata','Kabupaten','86611',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5309','53',NULL,'Flores Timur','Kabupaten','86213',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5310','53',NULL,'Sikka','Kabupaten','86121',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5311','53',NULL,'Ende','Kabupaten','86351',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5312','53',NULL,'Ngada','Kabupaten','86413',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5313','53',NULL,'Manggarai','Kabupaten','86551',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5314','53',NULL,'Rote Ndao','Kabupaten','85982',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5315','53',NULL,'Manggarai Barat','Kabupaten','86711',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5316','53',NULL,'Sumba Tengah','Kabupaten','87358',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5317','53',NULL,'Sumba Barat Daya','Kabupaten','87453',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5318','53',NULL,'Nagekeo','Kabupaten','86911',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5319','53',NULL,'Manggarai Timur','Kabupaten','86811',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5320','53',NULL,'Sabu Raijua','Kabupaten','85391',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('5371','53',NULL,'Kupang','Kota','85119',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6101','61',NULL,'Sambas','Kabupaten','79453',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6102','61',NULL,'Bengkayang','Kabupaten','79213',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6103','61',NULL,'Landak','Kabupaten','78319',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6104','61',NULL,'Pontianak','Kabupaten','78971',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6105','61',NULL,'Sanggau','Kabupaten','78557',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6106','61',NULL,'Ketapang','Kabupaten','78874',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6107','61',NULL,'Sintang','Kabupaten','78619',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6108','61',NULL,'Kapuas Hulu','Kabupaten','78719',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6109','61',NULL,'Sekadau','Kabupaten','79583',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6110','61',NULL,'Melawi','Kabupaten','78619',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6111','61',NULL,'Kayong Utara','Kabupaten','78852',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6112','61',NULL,'Kubu Raya','Kabupaten','78311',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6171','61',NULL,'Pontianak','Kota','78112',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6172','61',NULL,'Singkawang','Kota','79117',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6201','62',NULL,'Kotawaringin Barat','Kabupaten','74119',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6202','62',NULL,'Kotawaringin Timur','Kabupaten','74364',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6203','62',NULL,'Kapuas','Kabupaten','73583',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6204','62',NULL,'Barito Selatan','Kabupaten','73711',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6205','62',NULL,'Barito Utara','Kabupaten','73881',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6206','62',NULL,'Sukamara','Kabupaten','74712',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6207','62',NULL,'Lamandau','Kabupaten','74611',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6208','62',NULL,'Seruyan','Kabupaten','74211',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6209','62',NULL,'Katingan','Kabupaten','74411',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6210','62',NULL,'Pulang Pisau','Kabupaten','74811',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6211','62',NULL,'Gunung Mas','Kabupaten','74511',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6212','62',NULL,'Barito Timur','Kabupaten','73671',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6213','62',NULL,'Murung Raya','Kabupaten','73911',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6271','62',NULL,'Palangka Raya','Kota','73112',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6301','63',NULL,'Tanah Laut','Kabupaten','70811',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6302','63',NULL,'Kotabaru','Kabupaten','72119',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6303','63',NULL,'Banjar','Kabupaten','70619',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6304','63',NULL,'Barito Kuala','Kabupaten','70511',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6305','63',NULL,'Tapin','Kabupaten','71119',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6306','63',NULL,'Hulu Sungai Selatan','Kabupaten','71212',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6307','63',NULL,'Hulu Sungai Tengah','Kabupaten','71313',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6308','63',NULL,'Hulu Sungai Utara','Kabupaten','71419',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6309','63',NULL,'Tabalong','Kabupaten','71513',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6310','63',NULL,'Tanah Bumbu','Kabupaten','72211',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6311','63',NULL,'Balangan','Kabupaten','71611',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6371','63',NULL,'Banjarmasin','Kota','70117',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6372','63',NULL,'Banjarbaru','Kota','70712',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6401','64',NULL,'Paser','Kabupaten','76211',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6402','64',NULL,'Kutai Barat','Kabupaten','75711',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6403','64',NULL,'Kutai Kartanegara','Kabupaten','75511',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6404','64',NULL,'Kutai Timur','Kabupaten','75611',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6405','64',NULL,'Berau','Kabupaten','77311',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6409','64',NULL,'Penajam Paser Utara','Kabupaten','76311',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6471','64',NULL,'Balikpapan','Kota','76111',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6472','64',NULL,'Samarinda','Kota','75133',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6474','64',NULL,'Bontang','Kota','75313',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6501','65',NULL,'Malinau','Kabupaten','77511',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6502','65',NULL,'Bulungan (Bulongan)','Kabupaten','77211',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6503','65',NULL,'Tana Tidung','Kabupaten','77611',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6504','65',NULL,'Nunukan','Kabupaten','77421',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('6571','65',NULL,'Tarakan','Kota','77114',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7101','71',NULL,'Bolaang Mongondow (Bolmong)','Kabupaten','95755',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7102','71',NULL,'Minahasa','Kabupaten','95614',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7103','71',NULL,'Kepulauan Sangihe','Kabupaten','95819',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7104','71',NULL,'Kepulauan Talaud','Kabupaten','95885',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7105','71',NULL,'Minahasa Selatan','Kabupaten','95914',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7106','71',NULL,'Minahasa Utara','Kabupaten','95316',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7107','71',NULL,'Bolaang Mongondow Utara','Kabupaten','95765',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7108','71',NULL,'Kepulauan Siau Tagulandang Biaro (Sitaro)','Kabupaten','95862',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7109','71',NULL,'Minahasa Tenggara','Kabupaten','95995',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7110','71',NULL,'Bolaang Mongondow Selatan','Kabupaten','95774',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7111','71',NULL,'Bolaang Mongondow Timur','Kabupaten','95783',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7171','71',NULL,'Manado','Kota','95247',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7172','71',NULL,'Bitung','Kota','95512',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7173','71',NULL,'Tomohon','Kota','95416',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7174','71',NULL,'Kotamobagu','Kota','95711',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7201','72',NULL,'Banggai Kepulauan','Kabupaten','94881',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7202','72',NULL,'Banggai','Kabupaten','94711',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7203','72',NULL,'Morowali','Kabupaten','94911',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7204','72',NULL,'Poso','Kabupaten','94615',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7205','72',NULL,'Donggala','Kabupaten','94341',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7206','72',NULL,'Toli-Toli','Kabupaten','94542',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7207','72',NULL,'Buol','Kabupaten','94564',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7208','72',NULL,'Parigi Moutong','Kabupaten','94411',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7209','72',NULL,'Tojo Una-Una','Kabupaten','94683',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7210','72',NULL,'Sigi','Kabupaten','94364',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7271','72',NULL,'Palu','Kota','94111',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7301','73',NULL,'Selayar (Kepulauan Selayar)','Kabupaten','92812',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7302','73',NULL,'Bulukumba','Kabupaten','92511',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7303','73',NULL,'Bantaeng','Kabupaten','92411',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7304','73',NULL,'Jeneponto','Kabupaten','92319',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7305','73',NULL,'Takalar','Kabupaten','92212',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7306','73',NULL,'Gowa','Kabupaten','92111',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7307','73',NULL,'Sinjai','Kabupaten','92615',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7308','73',NULL,'Maros','Kabupaten','90511',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7309','73',NULL,'Pangkajene Dan Kepulauan','Kabupaten','90611',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7310','73',NULL,'Barru','Kabupaten','90719',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7311','73',NULL,'Bone','Kabupaten','92713',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7312','73',NULL,'Soppeng','Kabupaten','90812',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7313','73',NULL,'Wajo','Kabupaten','90911',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7314','73',NULL,'Sidenreng Rappang/Rapang','Kabupaten','91613',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7315','73',NULL,'Pinrang','Kabupaten','91251',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7316','73',NULL,'Enrekang','Kabupaten','91719',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7317','73',NULL,'Luwu','Kabupaten','91994',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7318','73',NULL,'Tana Toraja','Kabupaten','91819',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7322','73',NULL,'Luwu Utara','Kabupaten','92911',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7325','73',NULL,'Luwu Timur','Kabupaten','92981',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7326','73',NULL,'Toraja Utara','Kabupaten','91831',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7371','73',NULL,'Makassar','Kota','90111',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7372','73',NULL,'Parepare','Kota','91123',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7373','73',NULL,'Palopo','Kota','91911',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7401','74',NULL,'Buton','Kabupaten','93754',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7402','74',NULL,'Muna','Kabupaten','93611',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7403','74',NULL,'Konawe','Kabupaten','93411',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7404','74',NULL,'Kolaka','Kabupaten','93511',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7405','74',NULL,'Konawe Selatan','Kabupaten','93811',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7406','74',NULL,'Bombana','Kabupaten','93771',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7407','74',NULL,'Wakatobi','Kabupaten','93791',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7408','74',NULL,'Kolaka Utara','Kabupaten','93911',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7409','74',NULL,'Buton Utara','Kabupaten','93745',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7410','74',NULL,'Konawe Utara','Kabupaten','93311',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7471','74',NULL,'Kendari','Kota','93126',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7472','74',NULL,'Bau-Bau','Kota','93719',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7501','75',NULL,'Boalemo','Kabupaten','96319',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7502','75',NULL,'Gorontalo','Kabupaten','96218',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7503','75',NULL,'Pohuwato','Kabupaten','96419',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7504','75',NULL,'Bone Bolango','Kabupaten','96511',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7505','75',NULL,'Gorontalo Utara','Kabupaten','96611',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7571','75',NULL,'Gorontalo','Kota','96115',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7601','76',NULL,'Majene','Kabupaten','91411',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7602','76',NULL,'Polewali Mandar','Kabupaten','91311',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7603','76',NULL,'Mamasa','Kabupaten','91362',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7604','76',NULL,'Mamuju','Kabupaten','91519',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('7605','76',NULL,'Mamuju Utara','Kabupaten','91571',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('8101','81',NULL,'Maluku Tenggara Barat','Kabupaten','97465',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('8102','81',NULL,'Maluku Tenggara','Kabupaten','97651',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('8103','81',NULL,'Maluku Tengah','Kabupaten','97513',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('8104','81',NULL,'Buru','Kabupaten','97371',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('8105','81',NULL,'Kepulauan Aru','Kabupaten','97681',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('8106','81',NULL,'Seram Bagian Barat','Kabupaten','97561',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('8107','81',NULL,'Seram Bagian Timur','Kabupaten','97581',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('8108','81',NULL,'Maluku Barat Daya','Kabupaten','97451',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('8109','81',NULL,'Buru Selatan','Kabupaten','97351',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('8171','81',NULL,'Ambon','Kota','97222',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('8172','81',NULL,'Tual','Kota','97612',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('8201','82',NULL,'Halmahera Barat','Kabupaten','97757',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('8202','82',NULL,'Halmahera Tengah','Kabupaten','97853',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('8203','82',NULL,'Kepulauan Sula','Kabupaten','97995',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('8204','82',NULL,'Halmahera Selatan','Kabupaten','97911',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('8205','82',NULL,'Halmahera Utara','Kabupaten','97762',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('8206','82',NULL,'Halmahera Timur','Kabupaten','97862',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('8207','82',NULL,'Pulau Morotai','Kabupaten','97771',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('8271','82',NULL,'Ternate','Kota','97714',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('8272','82',NULL,'Tidore Kepulauan','Kota','97815',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9101','91',NULL,'Fakfak','Kabupaten','98651',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9102','91',NULL,'Kaimana','Kabupaten','98671',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9103','91',NULL,'Teluk Wondama','Kabupaten','98591',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9104','91',NULL,'Teluk Bintuni','Kabupaten','98551',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9105','91',NULL,'Manokwari','Kabupaten','98311',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9106','91',NULL,'Sorong Selatan','Kabupaten','98454',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9107','91',NULL,'Sorong','Kabupaten','98431',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9108','91',NULL,'Raja Ampat','Kabupaten','98489',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9109','91',NULL,'Tambrauw','Kabupaten','98475',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9110','91',NULL,'Maybrat','Kabupaten','98051',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9171','91',NULL,'Sorong','Kota','98411',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9401','94',NULL,'Merauke','Kabupaten','99613',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9402','94',NULL,'Jayawijaya','Kabupaten','99511',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9403','94',NULL,'Jayapura','Kabupaten','99352',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9404','94',NULL,'Nabire','Kabupaten','98816',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9408','94',NULL,'Kepulauan Yapen (Yapen Waropen)','Kabupaten','98211',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9409','94',NULL,'Biak Numfor','Kabupaten','98119',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9410','94',NULL,'Paniai','Kabupaten','98765',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9411','94',NULL,'Puncak Jaya','Kabupaten','98979',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9412','94',NULL,'Mimika','Kabupaten','99962',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9413','94',NULL,'Boven Digoel','Kabupaten','99662',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9414','94',NULL,'Mappi','Kabupaten','99853',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9415','94',NULL,'Asmat','Kabupaten','99777',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9416','94',NULL,'Yahukimo','Kabupaten','99041',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9417','94',NULL,'Pegunungan Bintang','Kabupaten','99573',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9418','94',NULL,'Tolikara','Kabupaten','99411',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9419','94',NULL,'Sarmi','Kabupaten','99373',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9420','94',NULL,'Keerom','Kabupaten','99461',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9426','94',NULL,'Waropen','Kabupaten','98269',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9427','94',NULL,'Supiori','Kabupaten','98164',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9428','94',NULL,'Mamberamo Raya','Kabupaten','99381',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9429','94',NULL,'Nduga','Kabupaten','99541',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9430','94',NULL,'Lanny Jaya','Kabupaten','99531',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9431','94',NULL,'Mamberamo Tengah','Kabupaten','99553',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9432','94',NULL,'Yalimo','Kabupaten','99481',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9433','94',NULL,'Puncak','Kabupaten','98981',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9434','94',NULL,'Dogiyai','Kabupaten','98866',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9435','94',NULL,'Intan Jaya','Kabupaten','98771',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9436','94',NULL,'Deiyai (Deliyai)','Kabupaten','98784',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL),('9471','94',NULL,'Jayapura','Kota','99114',1,0,NULL,'2020-11-19 10:34:01',NULL,NULL);
/*!40000 ALTER TABLE `cities` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `couriers`
--

DROP TABLE IF EXISTS `couriers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `couriers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  `slug` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `couriers`
--

LOCK TABLES `couriers` WRITE;
/*!40000 ALTER TABLE `couriers` DISABLE KEYS */;
INSERT INTO `couriers` VALUES (1,'JNE','jne'),(2,'JNT','jnt'),(3,'SiCepat','sicepat'),(4,'IDExpress','ide'),(5,'AnterAja','ant');
/*!40000 ALTER TABLE `couriers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `features`
--

DROP TABLE IF EXISTS `features`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `features` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `key` longtext,
  `value` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `features`
--

LOCK TABLES `features` WRITE;
/*!40000 ALTER TABLE `features` DISABLE KEYS */;
INSERT INTO `features` VALUES (1,'product','create_product',NULL,NULL),(2,'product','update_product',NULL,NULL),(3,'product','delete_product',NULL,NULL),(4,'product','list_product',NULL,NULL),(5,'cart','add_to_cart',NULL,NULL),(6,'cart','update_cart',NULL,NULL),(7,'cart','delete_cart',NULL,NULL),(8,'transaction','checkout',NULL,NULL),(9,'transaction','list_transaction',NULL,NULL),(10,'transaction','list_sales_transaction',NULL,NULL),(11,'cart','list_cart',NULL,NULL),(12,'transaction','update status',NULL,NULL);
/*!40000 ALTER TABLE `features` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_brands`
--

DROP TABLE IF EXISTS `product_brands`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_brands` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(120) DEFAULT NULL,
  `slug` varchar(120) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_brands`
--

LOCK TABLES `product_brands` WRITE;
/*!40000 ALTER TABLE `product_brands` DISABLE KEYS */;
INSERT INTO `product_brands` VALUES (1,'Rexus','rexus'),(2,'Gateron','slug'),(3,'Outemu','outemu'),(4,'Noir','noir'),(5,'Vortex','vortex');
/*!40000 ALTER TABLE `product_brands` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_categories`
--

DROP TABLE IF EXISTS `product_categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_categories` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(120) DEFAULT NULL,
  `slug` varchar(120) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_categories`
--

LOCK TABLES `product_categories` WRITE;
/*!40000 ALTER TABLE `product_categories` DISABLE KEYS */;
INSERT INTO `product_categories` VALUES (1,'Metcanical Keyboard','mechanical-keyboard'),(2,'Switch','switch'),(3,'Foam','foam'),(4,'Keycaps','keycaps'),(5,'Other','other');
/*!40000 ALTER TABLE `product_categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_images`
--

DROP TABLE IF EXISTS `product_images`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_images` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `product_id` bigint(20) DEFAULT NULL,
  `image_name` longtext,
  `is_primary` tinyint(1) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_products_product_image` (`product_id`),
  CONSTRAINT `fk_product_images_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  CONSTRAINT `fk_products_product_image` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_images`
--

LOCK TABLES `product_images` WRITE;
/*!40000 ALTER TABLE `product_images` DISABLE KEYS */;
/*!40000 ALTER TABLE `product_images` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `category_id` bigint(20) DEFAULT NULL,
  `title` longtext,
  `slug` longtext,
  `description` longtext,
  `address` longtext,
  `phone` longtext,
  `instagram_account` longtext,
  `purchase_price` bigint(20) DEFAULT NULL,
  `selling_price` bigint(20) DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_products_category` (`category_id`),
  KEY `fk_products_user` (`user_id`),
  CONSTRAINT `fk_products_category` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`),
  CONSTRAINT `fk_products_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `provinces`
--

DROP TABLE IF EXISTS `provinces`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `provinces` (
  `id_province` varchar(30) COLLATE latin1_general_ci NOT NULL,
  `id_country` varchar(30) COLLATE latin1_general_ci DEFAULT NULL,
  `province_code` varchar(30) COLLATE latin1_general_ci DEFAULT NULL,
  `province_name` varchar(200) COLLATE latin1_general_ci NOT NULL,
  `is_active` tinyint(4) DEFAULT '1',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `is_used` tinyint(4) DEFAULT '0',
  `created_by` varchar(30) COLLATE latin1_general_ci DEFAULT NULL,
  `updated_by` varchar(30) COLLATE latin1_general_ci DEFAULT NULL,
  PRIMARY KEY (`id_province`),
  KEY `id_country` (`id_country`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `provinces`
--

LOCK TABLES `provinces` WRITE;
/*!40000 ALTER TABLE `provinces` DISABLE KEYS */;
INSERT INTO `provinces` VALUES ('11','102','11','Nanggroe Aceh Darussalam (NAD)',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('12','102','12','Sumatera Utara',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('13','102','13','Sumatera Barat',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('14','102','14','Riau',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('15','102','15','Jambi',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('16','102','16','Sumatera Selatan',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('17','102','17','Bengkulu',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('18','102','18','Lampung',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('19','102','19','Bangka Belitung',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('21','102','21','Kepulauan Riau',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('31','102','31','DKI Jakarta',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('32','102','32','Jawa Barat',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('33','102','33','Jawa Tengah',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('34','102','34','Daerah Istimewa Yogyakarta',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('35','102','35','Jawa Timur',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('36','102','36','Banten',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('51','102','51','Bali',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('52','102','52','Nusa Tenggara Barat (NTB)',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('53','102','53','Nusa Tenggara Timur (NTT)',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('61','102','61','Kalimantan Barat',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('62','102','62','Kalimantan Tengah',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('63','102','63','Kalimantan Selatan',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('64','102','64','Kalimantan Timur',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('65','102','65','Kalimantan Utara',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('71','102','71','Sulawesi Utara',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('72','102','72','Sulawesi Tengah',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('73','102','73','Sulawesi Selatan',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('74','102','74','Sulawesi Tenggara',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('75','102','75','Gorontalo',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('76','102','76','Sulawesi Barat',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('81','102','81','Maluku',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('82','102','82','Maluku Utara',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('91','102','92','Papua Barat',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL),('94','102','91','Papua',1,'2020-11-19 10:34:01','2020-11-19 10:34:01',0,NULL,NULL);
/*!40000 ALTER TABLE `provinces` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_type_features`
--

DROP TABLE IF EXISTS `user_type_features`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_type_features` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `id_user_type` bigint(20) DEFAULT NULL,
  `id_feature` bigint(20) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_user_type_features_feature` (`id_feature`),
  KEY `fk_user_type_features_user_type` (`id_user_type`),
  CONSTRAINT `fk_user_type_features_feature` FOREIGN KEY (`id_feature`) REFERENCES `features` (`id`),
  CONSTRAINT `fk_user_type_features_user_type` FOREIGN KEY (`id_user_type`) REFERENCES `user_types` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_type_features`
--

LOCK TABLES `user_type_features` WRITE;
/*!40000 ALTER TABLE `user_type_features` DISABLE KEYS */;
INSERT INTO `user_type_features` VALUES (1,1,1,NULL,NULL),(2,1,2,NULL,NULL),(3,1,3,NULL,NULL),(4,1,4,NULL,NULL),(5,2,9,NULL,NULL),(6,2,8,NULL,NULL),(7,2,7,NULL,NULL),(8,2,6,NULL,NULL),(9,2,5,NULL,NULL),(10,1,10,NULL,NULL),(11,1,12,NULL,NULL),(12,2,11,NULL,NULL);
/*!40000 ALTER TABLE `user_type_features` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_types`
--

DROP TABLE IF EXISTS `user_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_types` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_types`
--

LOCK TABLES `user_types` WRITE;
/*!40000 ALTER TABLE `user_types` DISABLE KEYS */;
INSERT INTO `user_types` VALUES (1,'Super Admin',NULL,NULL),(2,'Admin',NULL,NULL),(3,'User',NULL,NULL);
/*!40000 ALTER TABLE `user_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `email` longtext,
  `password` longtext,
  `username` longtext,
  `id_user_type` bigint(20) DEFAULT NULL,
  `address` longtext,
  `phone` longtext,
  `avatar` longtext,
  `status` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_users_user_type` (`id_user_type`),
  CONSTRAINT `fk_users_user_type` FOREIGN KEY (`id_user_type`) REFERENCES `user_types` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'Seller','seller@example.com','$2a$04$dCjR1rfu0e7CmhUrLHAPA.abcr6GaoVDgUIjpO.S7eHzc0YVbjqyi','seller',1,'Yogyakarta',NULL,'','Active','2021-02-27 21:55:09.568','2021-02-27 21:55:09.568'),(2,'Buyer','buyer@example.com','$2a$04$dCjR1rfu0e7CmhUrLHAPA.abcr6GaoVDgUIjpO.S7eHzc0YVbjqyi','buyer',2,'Yogyakarta',NULL,'','Active','2021-02-27 21:55:09.571','2021-02-27 21:55:09.571'),(5,'Khaiz Badaru Tammam','khaiz@gmail.com','$2a$04$REm/DjZiSPdMDhnXolOrNOZYlTNCF4usCgEoyjhQeXQeJlHt2IjWS','tammam',2,NULL,'',NULL,'Active',NULL,NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-01  8:52:24
