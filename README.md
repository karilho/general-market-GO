# General Market GO
This is my personal project to learn Golang and some concepts of microservices, DDD, Cloud Deployment and others.

## How it was planned?

<a href="https://ibb.co/3yHYMNN"><img src="https://i.ibb.co/fY7Gnpp/example.png" alt="example" border="0"></a>

The central point is a BUYER, that can buy products from the market.

When a buyer buys a product, he generate a buy_order, with the total value and order date.

We have too a purchased_products table, that stores the products that were purchased in each buy_order, individual price and quantity

### Desnormalization Decision

I choose this because i think make relationships between tables it's a not necessary for now, but probably i will add this in the future.

## Goals

* [] Make API WORK doing some CRUD operations.
* [] A simple frontend to communicate with the API.
* [] Deploy on the Dockerhub
* [] Deploy on the AWS using Localstack
* [] Create some funtions to use some main AWS features like Lambda, SQS

``` 
CURL EXAMPLES

curl --location --request POST 'http://localhost:8080/buyers' \
```

-- > SERVICE - AGRUPA PODS NA REDE INTERNA DO K8S E DA NOME DE DOMINIO PARA ELES PERTIMINDO COMUNICAÇÃO ENTRE ELES

--> CRIAMOS UM SERVICE PARA O POSTGRES
DEPOIS USAMOS OS SEGUINTES COMANDOS ->

k apply -f service-postgres.yaml 
## para rodar o service criando um hostname para oS podS do postgres

## depois matamos o pod com:::: é seguro pois o deployment vai recria-lo
k delete pod nomedopoddaapi

## depois redirecionamos utilizando o port-forward COM O NOME DO POD, inclusive pro banco de dados
k port-forward nomedopoddaapi 8080:8080

## pegar as configurações de todos os clusters existentes (isso que é deployado no k8s) 
 k config get-contexts

## troca o contexto do kubectl para o cluster que queremos
k config use-context nomedocontexto


