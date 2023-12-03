# General Market GO
This is my personal project to learn Golang and some concepts of microservices, DDD, Cloud Deployment and others.

## How it was planned?

<a href="https://ibb.co/3yHYMNN"><img src="https://i.ibb.co/fY7Gnpp/example.png" alt="example" border="0"></a>


<img src="https://a.imagem.app/oX9way.png" alt="oX9way.png" border="0"/>


The central point is a BUYER, that can buy products from the market.

When a buyer buys a product, he generate a buy_order, with the total value and order date.

We have too a purchased_products table, that stores the products that were purchased in each buy_order, individual price and quantity

### Desnormalization Decision

I choose this because i think make relationships between tables it's a not necessary for now, but probably i will add this in the future.

## Goals

* [x] Make API WORK doing some CRUD operations.
* [x] Deploy on the Dockerhub
* [x] Deploy on the AWS using Localstack - CHANGED - Deploy on k8s using KIND
* [x] Make CI/CD deploy with some features like Github Actions
* [x] Improve API with some features, like: CRUD operations for buyer.
* [x] Implement logs
* [x] Create some AWS funtions: 
1. [x] S3 Bucket to store JSON + endpoint to call this (register buyer)
2. [x] Lambda to get from S3 the JSON file and send to SQS Queue.
3. [] Worker that runs with Goroutines to get from SQS and send to Postgres.
* [] Create unity tests.
* [] Validation using a middleware to give user permissions to call some endpoints.
* [] A simple frontend to communicate with the API.


## About Service 

<a href="https://ibb.co/n7CQ6Bh"><img src="https://i.ibb.co/Cb2Vm6x/flow.png" alt="flow" border="0"></a>

### Buyer Service Payload:

***EndPoint (EP)->  localhost:80/api/createBuyer*** ::::::::: Create a Buyer

    {
        "current_type": "buyer",
        "username": "TESTES3--ALPHA",
        "email": "johndoe@example.com",
        "password_hash": "hashedpassword",
        "full_name": "John Doe",
        "phone_number": "1234567890",
        "registration_date": "2022-12-01T14:15:22Z",
        "street_address": "123 Main St",
        "place_number": "1A",
        "city": "City",
        "state_province": "State",
        "postal_code": "12345"
    }

 ***Curl:***
````
    curl --location 'localhost:3000/createBuyer' \
    --header 'Content-Type: application/json' \
    --data-raw '{
      "current_type": "buyer",
      "username": "TESTES3--ALPHA",
      "email": "johndoe@example.com",
      "password_hash": "hashedpassword",
      "full_name": "John Doe",
      "phone_number": "1234567890",
      "registration_date": "2022-12-01T14:15:22Z",
      "street_address": "123 Main St",
      "place_number": "1A",
      "city": "City",
      "state_province": "State",
      "postal_code": "12345"
    }
    '
````

***EndPoint (EP)->  localhost:80/api/getBuyer/{buyerId}*** ::::: Retrieve a buyer info


