
<a href= "https://raullesteves.medium.com/github-como-fazer-um-readme-md-bonitão-c85c8f154f8"> README LINDO </a>

## General Market GO, passo a passo criação:

1. Criei a API.
2. Criei o Compose pro banco e o Dockerfile da API.
3. Comecei a migração pro k8s, simulando local com o <a href="https://kind.sigs.k8s.io"> Kind </a>.
4. Criei o deployment e o service do Postgres nos arquivos .yaml.
5. Expus para o exterior fazendo a instalação do  <a href="https://kind.sigs.k8s.io/docs/user/ingress/"> INGRESS </a>.
6. Criei as migrations diretamente na API, falta implementar um LOCK. 
7. Criei o deployment e o service da API na porta 3000, mas o ingress vai fazer o redirecionamento para a porta 80.
8. Criei todo o processo de CI/CD com o Github Actions, fazendo o build da imagem e o push para o dockerhub.
9. Integrei todo o processo com o k8s, utilizando o Flux como auxiliar para deploy da imagem e atualização do mesmo.

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
docker rmi nomedaimagemg
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

## Flux

**Passo a passo do que foi feito:**

1. Instalar o fluxctl -><a href="https://fluxcd.io/flux/get-started/ "> Flux </a>
2. Scan na imagem, auto-atualizações etc -> <a href="https://fluxcd.io/flux/guides/image-update/"> Image update </a>

**Instalar o fluxctl**

Tipo de imagem:
* GITREPOSITORY: Em resumo, a criação desse objeto GitRepository é uma parte do processo de configuração do Flux CD para sincronizar automaticamente o código do repositório https://github.com/stefanprodan/podinfo com o cluster Kubernetes, garantindo que o cluster esteja sempre atualizado com o código mais recente desse repositório.
* KUSTOMIZATION:  é usado para definir como personalizar e implantar recursos no cluster Kubernetes com base nos arquivos de personalização presentes no repositório Git referenciado pelo GitRepository chamado "podinfo." Ele define várias configurações, incluindo o intervalo de sincronização, o caminho para os arquivos de personalização, a referência à fonte Git e outras configurações relacionadas à personalização.
* Validar se o flux tá rodando e olhando pro repo: 

````
flux get kustomizations --watch
````
* Instalar o bootstrap dele para que ele crie o repo com as permissões:
````
flux bootstrap github \
  --components-extra=image-reflector-controller,image-automation-controller \
  --owner=$GITHUB_USER \
  --repository=flux-image-updates \
  --branch=main \
  --path=clusters/my-cluster \
  --read-write-key \
  --personal
````

* Avisar ao flux para PUXAR a imagem e aplicar as mudanças para detectar suas modificações:
````
flux reconcile kustomization flux-system --with-source
````

* Configurar o image scan para que ele faça o scan da imagem e atualize o flux:
1. Crie o IMAGEREPOSITORY: Sua função é avisar ao flux qual registro de container ele deve scanear em busca de noas tagsde imagem. Ele também define o intervalo de verificação para o flux verificar se há novas tags de imagem.
2. Crie o IMAGEPOLICY: Sua função é definir uma política de imagem que o flux deve seguir. Ele define a política de atualização de imagem para o flux, que pode ser automática ou manual, também define a ordem que ele vai pegar as tags (útlima pra mais recente por exemplo).
3. Crie o IMAGEUPDATEAUTOMATION: Sua função é avisar qual ordem vai ser utilizada quando o FLUX for fazer o update da imagem (adicionar o $imagepolicy no manifesto lá da imagem que tá no repo).


* Pegar o arquivo lambda.go e fazer o build dele para o linux, para que o lambda consiga ler:
````
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main lambda.go
````

* Criar o zip do arquivo:

````
 zip nomedozip.zip main
````


1. Toda vez que fizer uma pretensão de compra, SALVAR O JSON NUM BUCKET DO S3 (só pra usar, poderia mandar direto pro SQS) (Posteriormente -> Adicionar serviço de envio de e-mail para cadastrados e sem compras.)
2. Quando for salvo no bucket do s3, trigger um LAMBDA que pega esse JSON DO S3 {} e joga pra uma fila SQS (somente pra evitar condição de corrida)
3. Componente do LAMBDA (que vai ser criado DENTRO DA AWS), 
Função = pegar o json do s3, jogar no SQS. 
Ativação = NOVA ENTRADA NO S3.

//Todo: Criar o WORKER


4. Ter um worker do go lendo dessa FILA DO SQS E salvando no DB os efetivados (campo buy order)


POR FIM GARANTIR QUE NÃO HAVERÁ FALHA DE VULNERABILIDADE PRA CRIAR O PONTO (muitas req)

SQS -> CASO DE USO MAIS COMUM - GARANTIR UM RETRY EM CASO DE FALHA (COM A IMPLEMENTAÇÃO DE UMA DEAD LETTER QUEUE {POSTERIOR})
LAMBDA -> No caso que eu to esperando um TSUNAMI DE REQUISIÇÃO, MAS moldar todos os prédios pra isso só 1x ao ano não compensa.
 -> No caso, eu poderia ter um LAMBDA que fica rodando o tempo todo, e quando eu precisar de mais poder de processamento, eu só aumento o número de LAMBDA que eu quero rodando.


### Anotações sobre AWS

* Lembrar de criar um usuário com permissões de administração, e pegar a chave de acesso e a chave secreta.
* Quando criar um LAMBDA, lembrar de criar uma role para ele, e dar permissão de acesso ao S3 e ao SQS.
* No Lambda, alterar o runtime settings e colocar MAIN para que ele leia o main do seu codigo.
* Quando criar, criar com runtime GO mesmo, arquitetura x86