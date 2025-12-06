
DROP TABLE IF EXISTS `comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `cultural_id` int DEFAULT NULL,
  `cultural_type` enum('event','tourist_attraction') NOT NULL DEFAULT 'event',
  `user_id` int DEFAULT NULL,
  `content` text NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comments`
--

LOCK TABLES `comments` WRITE;
/*!40000 ALTER TABLE `comments` DISABLE KEYS */;
INSERT INTO `comments` VALUES (7,2,'event',1,'Adorei a feira! Muito organizado.','2025-11-05 00:44:23','2025-11-05 00:44:23'),(8,2,'event',1,'Ótima estrutura.','2025-11-05 00:48:26','2025-11-05 00:48:26'),(9,2,'event',1,'A','2025-11-05 00:48:42','2025-11-05 00:48:42'),(10,2,'event',1,'B','2025-11-05 00:48:47','2025-11-05 00:48:47'),(11,2,'event',1,'C','2025-11-05 00:48:51','2025-11-05 00:48:51'),(12,2,'event',1,'D','2025-11-05 00:52:04','2025-11-05 00:52:04'),(13,11,'tourist_attraction',10,'Ótimo para passear!','2025-11-11 22:05:45','2025-11-11 22:05:45'),(14,14,'tourist_attraction',10,'Adorei a visita, ótimo para crianças!','2025-11-17 00:01:12','2025-11-17 00:01:12');
/*!40000 ALTER TABLE `comments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `events`
--

DROP TABLE IF EXISTS `events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `events` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `location` varchar(255) DEFAULT NULL,
  `description` text,
  `start_date` varchar(255) NOT NULL,
  `end_date` varchar(255) NOT NULL,
  `duration_hours` varchar(255) NOT NULL,
  `is_accessible` tinyint(1) DEFAULT '1',
  `price` varchar(255) NOT NULL DEFAULT 'R$0,00',
  `image` varchar(255) DEFAULT NULL,
  `organizer_id` int DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events`
--

LOCK TABLES `events` WRITE;
/*!40000 ALTER TABLE `events` DISABLE KEYS */;
INSERT INTO `events` VALUES (1,'I Wanna Be Tour','Estádio Independência - R. Pitangui, 3230 - Horto, Belo Horizonte - MG, 31030-066',NULL,'10/03/2024','11/03/2024','12:00 - 01:00',1,'R$125 - 650','iwnb-event.png',13,'2025-08-04 23:48:14','2025-08-20 20:15:57'),(2,'Bienal do Livro','Av. Amazonas, 6200 - Gameleira, Belo Horizonte - MG, 30510-000','A maior feira de livros de Minas Gerais.','03/03/2024','10/03/2024','8:00 - 20:00',1,'R$0,00','bienal-event.png',10,'2025-08-24 21:30:45','2025-10-05 16:54:30'),(3,'MCR Ao Vivo','Av. Antônio Abrahão Caram, 1001 - São José, Belo Horizonte - MG, 31275-000','Uma das maiores bandas de todos os tempos em show inédito. Pela primeira vez em BH, My Chemical Romance promete uma noite inesquecível.','23/10/2024','','20:00 - 23:00',1,'R$400 - 1000','mcr-event.png',5,'2025-10-07 23:07:27','2025-10-07 23:07:27'),(4,'Jogo do Cruzeiro','Av. Antônio Abrahão Caram, 1001 - São José, Belo Horizonte - MG, 31275-000','Bora torcer? Semi-final da Sul-Americana, Cruzeiro x Lanús.','23/10/2024','','19:30 - 23:00',0,'R$50 - 300','cru-event.png',27,'2025-10-07 23:07:27','2025-10-07 23:07:27'),(5,'DCC Week 2023','Av. Pres. Antônio Carlos, 6627 - Pampulha, Belo Horizonte - MG, 31270-901','Venha ampliar seu conhecimento e fazer contatos importantes em nossa feira de oportunidades, aberta ao público.','2024-04-08T00:00','2024-06-08T00:00','8:30 - 18:30',1,'R$0,00','dccweek-event.png',10,'2025-10-07 23:07:27','2025-11-28 13:32:44');
/*!40000 ALTER TABLE `events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `notifications`
--

DROP TABLE IF EXISTS `notifications`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `notifications` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `cultural_id` int DEFAULT NULL,
  `cultural_type` enum('event','tourist_attraction') NOT NULL DEFAULT 'event',
  `user_id` int DEFAULT NULL,
  `type` enum('updated','canceled','commented','closed') NOT NULL DEFAULT 'updated',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `title` varchar(255) DEFAULT NULL,
  `seen` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `notifications`
--

LOCK TABLES `notifications` WRITE;
/*!40000 ALTER TABLE `notifications` DISABLE KEYS */;
INSERT INTO `notifications` VALUES (1,2,'event',10,'updated','2025-09-28 19:21:39',NULL,1),(2,2,'event',10,'updated','2025-10-02 00:46:06',NULL,1),(3,2,'event',10,'updated','2025-10-02 00:47:27',NULL,1),(4,2,'event',10,'updated','2025-10-02 00:50:00',NULL,1),(5,2,'event',10,'updated','2025-10-02 00:50:39',NULL,1),(6,2,'event',10,'updated','2025-10-02 00:54:35',NULL,1),(7,2,'event',10,'updated','2025-10-02 00:57:32',NULL,1),(8,2,'event',10,'updated','2025-10-02 00:57:48',NULL,1),(9,2,'event',10,'updated','2025-10-02 00:58:54',NULL,1),(10,2,'event',10,'updated','2025-10-02 01:01:55',NULL,1),(11,2,'event',10,'updated','2025-10-02 01:02:33',NULL,1),(12,2,'event',10,'updated','2025-10-02 01:02:42',NULL,1),(13,2,'event',10,'updated','2025-10-02 01:02:44',NULL,1),(14,2,'event',10,'updated','2025-10-05 16:54:33',NULL,1),(15,2,'event',10,'updated','2025-10-05 16:54:52',NULL,1),(16,2,'event',10,'updated','2025-10-05 16:55:03',NULL,1),(17,2,'event',10,'updated','2025-10-05 16:56:25',NULL,1),(18,2,'event',10,'updated','2025-10-05 16:56:29',NULL,1),(19,2,'event',10,'updated','2025-10-05 16:56:46',NULL,1),(20,14,'tourist_attraction',10,'updated','2025-11-20 16:55:27',NULL,1),(21,14,'tourist_attraction',10,'updated','2025-11-20 16:56:24',NULL,1),(22,14,'tourist_attraction',10,'updated','2025-11-20 16:56:46',NULL,1),(23,14,'tourist_attraction',10,'updated','2025-11-20 16:58:24',NULL,1),(24,14,'tourist_attraction',10,'updated','2025-11-20 16:58:45',NULL,1),(25,5,'event',10,'updated','2025-11-28 13:32:48',NULL,1);
/*!40000 ALTER TABLE `notifications` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tourist_attractions`
--

DROP TABLE IF EXISTS `tourist_attractions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tourist_attractions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `location` varchar(255) DEFAULT NULL,
  `description` text,
  `working_hours` varchar(255) NOT NULL,
  `is_accessible` tinyint(1) DEFAULT '1',
  `price` varchar(255) NOT NULL DEFAULT 'R$0,00',
  `image` varchar(255) DEFAULT NULL,
  `organizer_id` int DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tourist_attractions`
--

LOCK TABLES `tourist_attractions` WRITE;
/*!40000 ALTER TABLE `tourist_attractions` DISABLE KEYS */;
INSERT INTO `tourist_attractions` VALUES (3,'Praça Raul Soares','Av. Amazonas, Centro, MG - 30180-001','Praça de livre acesso o tempo todo.','',1,'R$0,00','e4482ac2-d79f-4efe-b8c4-38070204e410.webp',10,'2025-09-28 21:36:56','2025-09-28 21:36:56'),(11,'Praça da Liberdade','Av. Bias Fortes, Savassi, Belo Horizonte - MG','Uma praça histórica cercada por museus e edifícios icônicos. Parte do Circuito Cultural Praça da Liberdade e um dos pontos turísticos mais visitados de BH.','',1,'R$0,00','liberdade-attraction.png',10,'2025-11-11 21:58:19','2025-11-20 00:18:53'),(12,'Mercado Central','Av. Augusto de Lima, 744 - Centro, Belo Horizonte - MG','Famoso mercado com mais de 400 lojas, vendendo queijos, doces, artesanato e produtos típicos mineiros.','',1,'R$0,00','mercado-attraction.png',10,'2025-11-11 21:58:19','2025-11-11 21:58:19'),(13,'Igrejinha da Pampulha','Av. Otacílio Negrão de Lima, 3000 - Pampulha, Belo Horizonte - MG','Obra-prima de Oscar Niemeyer com painéis de Portinari. Parte do Conjunto Arquitetônico da Pampulha, Patrimônio da UNESCO.','Domingo	08:00–17:00\nSegunda-feira 08:00-17:00\nTerça-feira	08:00–17:00\nQuarta-feira	08:00–17:00\nQuinta-feira	08:00–17:00\nSexta-feira	08:00–17:00\nSábado	08:00–17:00\n',1,'R$30,00','igrejinha-attraction.png',10,'2025-11-11 21:58:19','2025-11-20 13:59:31'),(14,'Parque das Mangabeiras','Av. José do Patrocínio Pontes, 580 - Mangabeiras, Belo Horizonte - MG','Grande parque na Serra do Curral com mirantes, trilhas e áreas de lazer. Oferece uma vista espetacular da cidade.','Domingo	08:00–17:00\nSegunda-feira	Fechado\nTerça-feira	08:00–17:00\nQuarta-feira	08:00–17:00\nQuinta-feira	08:00–17:00\nSexta-feira	08:00–17:00\nSábado	08:00–17:00\n',1,'R$0,00','mangabeiras-attraction.png',10,'2025-11-11 21:58:19','2025-11-20 16:58:09'),(15,'Pirulito da Praça Sete','Praça Sete de Setembro - Centro, Belo Horizonte - MG','Obelisco icônico doado pelo povo de Betim, localizado no cruzamento das avenidas Afonso Pena e Amazonas. Marco zero da cidade.','',1,'R$0,00','psete-attraction.png',10,'2025-11-11 21:58:19','2025-11-11 21:58:19');
/*!40000 ALTER TABLE `tourist_attractions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_favorites`
--

DROP TABLE IF EXISTS `user_favorites`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_favorites` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `favorite_type` enum('event','tourist_attraction') NOT NULL DEFAULT 'event',
  `favorite_id` bigint unsigned NOT NULL,
  `favorited_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_seen_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_favorites`
--

LOCK TABLES `user_favorites` WRITE;
/*!40000 ALTER TABLE `user_favorites` DISABLE KEYS */;
INSERT INTO `user_favorites` VALUES (20,10,'event',5,'2025-10-07 23:48:26','2025-11-28 13:32:55'),(27,10,'event',2,'2025-10-12 21:37:32','2025-11-20 17:23:58'),(28,1,'event',1,'2025-10-13 21:53:51','2025-11-11 21:14:56'),(29,1,'event',3,'2025-10-20 23:04:29','2025-10-20 23:12:14'),(30,1,'event',2,'2025-11-05 00:49:42','2025-11-11 20:08:21'),(31,10,'tourist_attraction',11,'2025-11-11 22:05:52','2025-11-20 14:00:37'),(32,10,'tourist_attraction',14,'2025-11-11 23:14:01','2025-11-21 22:24:03'),(33,10,'tourist_attraction',3,'2025-11-12 00:40:44','2025-11-20 17:18:14'),(36,10,'tourist_attraction',13,'2025-11-23 15:34:48','2025-11-28 13:31:08'),(38,10,'event',4,'2025-11-24 19:16:16','2025-11-24 19:16:20'),(39,10,'event',1,'2025-11-28 13:28:46','2025-11-28 13:30:07');
/*!40000 ALTER TABLE `user_favorites` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `document` varchar(50) NOT NULL,
  `company_name` varchar(100) DEFAULT NULL,
  `type` varchar(50) NOT NULL,
  `password` varchar(100) NOT NULL,
  `document_type` enum('CPF','CNPJ') NOT NULL DEFAULT 'CPF',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'Clara :)','clara004.costa@gmail.com','09801160667','','common','hellohello','CPF','2025-08-20 20:14:45','2025-08-20 20:14:47'),(2,'Clara Costa da Fonseca','clara.costa@studiosol.com.br',' 62173620000180','','common','People','CPF','2025-08-20 20:14:45','2025-08-20 20:14:47'),(3,'Clara Costa da Fonseca','clarafonseca@dcc.ufmg.br',' 62173620000180','','common','People','CPF','2025-08-20 20:14:45','2025-08-20 20:14:47'),(5,'Clara Costa da Fonseca','claraufmg4@gmail.br','62173620000180','set it up','organizer','People','CPF','2025-08-20 20:14:45','2025-08-20 20:14:47'),(10,'Clara Costa da Fonseca','claraufmg4@gmail.com','62173620000189','Sit&Cry','organizer','bolinhoo','CNPJ','2025-08-20 20:14:45','2025-08-20 20:14:47'),(11,'IZA CHATA','izachata@chata.com','00000000000','','common','Chatachata','CPF','2025-08-20 20:14:45','2025-08-20 20:14:47'),(13,'Vin D. Jr.','triplex@hotmail.com','000000000000000','Con&Diesel','organizer','Caveends','CNPJ','2025-08-20 20:14:45','2025-08-20 20:14:47'),(14,'Fala','serio@serio.com','00000000000','set it up','organizer','capecape','CNPJ','2025-08-20 20:14:45','2025-08-20 20:14:47'),(27,'Clara Costa da Fonseca','clara004.costa@gmail.com.br','09801160667333','Claraboia','organizer','capecape','CNPJ','2025-08-20 20:14:45','2025-08-20 20:14:47');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

-- Dump completed on 2025-12-05 18:54:30
