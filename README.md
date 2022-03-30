# Warung Pintar Pre Test Backend Engineer

Here are the results of the Warung Pintar Backend Engineer pre test project

## Overview

1. This Repo using [Gin-Gonic](https://github.com/gin-gonic/gin)
2. Gorm for query
3. MySQL for master Data User, Brand, Product Category, Kecamatan & Kabupaten,
4. Elastic Search for save Master Data Product
5. MongoDB use for save data Cart and Transaction
6. JWT for Authentication
7. Middleware Using [Secure Middleware](https://github.com/unrolled/secure)
8. Docker

## Installation

To Run this Repo, you need to install [Go](https://golang.org/dl/) and Have a Docker in your Operating System

## Usage

1. Clone this repo using https `git clone https://gitlab.com/khaizbt/privy-test.git warung-pintar-test`
2. in your bash, run `cd warung-pintar-test`
3. You can configure and run docker-compose.yml `docker-compose up -d`
4. You can use any RDBMS Database, for configuration the database is contained in the file [.env](https://gitlab.com/khaizbt/privy-test/-/blob/master/.env) and Docs of database in [here](https://gorm.io/docs/connecting_to_the_database.html)
5. If you have configured the Database and Secret Key on [.env](https://gitlab.com/khaizbt/privy-test/-/blob/master/.env) you can immediately run `go run main.go ` use bash in Project Directory
6. You can Import Database to MongoDB, Elastic Search and MySQL
7. If it's run successfully, you can import postman.json on your postman in the file [postman](https://gitlab.com/khaizbt/privy-test/-/blob/master/Privy.postman_collection.json) or via the link [here](https://documenter.getpostman.com/view/12945074/UVeCQoNQ)

## Thanks

Thank you Warung Pintar team, sorry for the shortcomings in this project, I am waiting for good news from the PrivyID team ;)
