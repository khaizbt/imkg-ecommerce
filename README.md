# Golang Clean Arch Example

Here are the results example for Golang Clean Architecture project

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

1. Clone this repo using https

```bash
git clone https://github.com/khaizbt/imkg-ecommerce warung-pintar-test
```

2. in your bash, run `cd warung-pintar-test`
3. if everything is configured, run

```bash
docker compose up -d
```

4. In MySQL, you can configure database name, host, user, password adjusted with .env and docker credentials. and Import SQL file [imkg-ecmommerce](https://github.com/khaizbt/imkg-ecommerce/blob/main/imkg-ecommerce.sql)
5. In MongoDB, you can make database Eg.: imkg_database and create 2 collection name Eg.: carts and transactions(customized with .env files)
6. In Elastic search you can create index Eg: products (customized with .env files)
7. If Elasticsearch can't run due to bootstrap check, you can run the following command

```bash
wsl -d docker-desktop
sysctl -w vm.max_map_count=262144
```

8. If you have configured the Database and Secret Key on [.env](https://github.com/khaizbt/imkg-ecommerce/blob/main/.env) you can immediately run `go run main.go ` use bash in Project Directory

9. If it's run successfully, you can import postman.json on your postman in the file [postman](https://github.com/khaizbt/imkg-ecommerce/blob/main/IMKG%20ECommerce.postman_collection.json) or via the link [here](https://documenter.getpostman.com/view/12945074/UVeCQoNQ)

## Rule of Thumb

1. run `docker-compose up -d` in your root repository
2. run `go run main.go`
3. If it is running well and the databases are connected to each other, you can register or login with the following account

```bash
Seller
email : seller@example.com
password : password123

Buyer
Email : buyer@example.com
Password : password123
```

As a Seller you can Create Product, Update, Delete, Check All Transactions, Update Transaction Status; and
As a Buyer you can Add To Cart, Update Cart, Delete Cart, Checkout, Get Purchase List

4. As a Seller, you must Post Products first so that Buyers can make purchases (Add To Cart) [End Point Product/Create Product]
5. Buyers can add to cart(End Point Cart/\*) and do Checkout at Endpoint(Transaction/Checkout) and see a list of purchases at (Transaction/List Transaction)
6. If the Buyer has already placed an Order (Checkout) you can change the order status to Sent, Confirmed or Canceled

## Unit Test

Due to time constraints, I haven't had time to create Unit Tests yet, but I will create a dev_unit_test Branch, you can use the Branch to see the unit tests. I'm not going to add unit tests to the main branch because I respect the expired deadline

## Thanks

Thank you, sorry for the shortcomings in this project, If you have feedback or improvement in this project, Just PR! :)
