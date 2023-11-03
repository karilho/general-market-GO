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
* [] Make CI/CD deploy with some features like Github Actions
* [] Improve API with some features like JWT, Auth, and unity tests
* [] Create some AWS funtions like Lambda, SQS
* [] A simple frontend to communicate with the API.





