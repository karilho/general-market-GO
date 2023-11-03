
<a href= "https://raullesteves.medium.com/github-como-fazer-um-readme-md-bonitão-c85c8f154f8"> README LINDO </a>

## General Market GO, passo a passo criação:

1. Criei a API.
2. Criei o Compose pro banco e o Dockerfile da API.
3. Comecei a migração pro k8s, simulando local com o <a href="https://kind.sigs.k8s.io"> Kind </a>.
4. Criei o deployment e o service do Postgres nos arquivos .yaml.
5. Expus para o exterior fazendo a instalação do  <a href="https://kind.sigs.k8s.io/docs/user/ingress/"> INGRESS </a>.
6. Criei as migrations diretamente na API, falta implementar um LOCK. 
7. Criei o deployment e o service da API na porta 3000, mas o ingress vai fazer o redirecionamento para a porta 80.


### Estudos K8S

**O que é um Cluster?**

 Kubernetes gerencia um GRUPO/CONJUNTO/CLUSTER de máquinas, chamadas de NODES.
Um Cluster de K8S possui dois tipos de recursos:
* Control plane -> Coordenação do cluster.
* Nodes -> Máquinas

**O que é um Node?**

 Um Node é uma máquina que faz parte do cluster, pode ser uma VM ou uma máquina física. 

**O que é um Pod?**

 Um POD é uma unidade básica de um cluster, é um grupo de um ou mais containers que compartilham armazenamento e rede.

## Lista de Comandos úteis:
# K8S
**Criar POD com base no Kind = Deployment**
````
k apply -f nomedoARQUIVO.yaml
````

**Matar um Pod com base no POD, se quiser, pode mudar pra service**
````
k delete pod nomedopodPODoUservico
````

**Listar os logs de um POD**
````
k logs nomedopodPOD
````

**Fazer redirecionamento com K8S para portas externas. (no caso, 8080 é a porta externa e 8080 é a porta interna)**
````
k port-forward nomedopoddaapiPOD 8080:8080
````
**Listar os contextos do CLUSTERS atuais do kubectl / Trocar para outro contexto.**
````
k config get-contexts
k config use-context NOMEDOCONTEXTO
````

# DOCKER

**Criar uma imagem com base no dockerfile do diretório // Empurrar pro dockerhub a bendita.**
````
docker build -t nomedockerhub/repodockerhub:VERSAOTAG .
docker push nomedockerhub/repodockerhub:VERSAOTAG

docker build -t karilho/generalmarket:1.0.0 .
docker push karilho/general-market-go:1.0.0
````

**Rodar uma imagem com base no nome da imagem**
````
docker run nomedaimagem
````

**Listar as imagens / Listar os containers "-a" ao final faz mostrar só os que tão rodando**
````
docker images
docker ps
````

**Parar um container / remover container**
````
docker stop nomedocontainer
docker rm nomedocontainer
````

**Remover uma imagem**
````
docker rmi nomedaimagem
````

**TAG imagem (não tão funcional)**
````
docker tag nomedaimagem espacoHUBREPO:versao
docker tag general-market-go karilho/general-market-go:v1.0.0
````

**Fazer o login no dockerhub**
````
docker login -u "login" -p "pass" docker.io
````

**Entrar dentro da máquina do container**
````
docker exec -it nomedocontainer bash
````


