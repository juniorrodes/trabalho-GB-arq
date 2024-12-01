# Trabalho Grau B Arquitetura e organização de computadores.

## Pré requisitos:
- kubectl
- helm
- minikube ou qualquer ferramenta de criação de cluster k8s


## Como rodar o projeto.
Inicialmente será necessário iniciar um cluster K8s, para fazer isso você pode usar a sua ferramenta de escolha, no caso deste projeto foi escolhido usar [minikube](https://minikube.sigs.k8s.io/docs/) uma ferramenta de criação e gerenciamento de clusters K8s.

Comece criando um cluster rodando o comando:
``minikube start``
Para garantir que o cluster está rodando corretamente você pode rodar o comando `kubectl get nodes` se o comando listar pelo menos um nó o cluster está rodando corretamente.
Com o cluster rodando, rode o comando `helm install my-app app` isto irá criar os recursos necessários para a aplicação dentro do cluster. Depois disso rode o comando `minikube tunnel`(preferencialmente rode isso em outro terminal) este comando irá expor o load balancer criado no cluster para a sua máquina, para que seja possível chamar a aplicação como um serviço local.
