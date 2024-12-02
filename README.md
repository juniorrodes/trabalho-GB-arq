# Trabalho Grau B Arquitetura e organização de computadores.

## Pré requisitos:
- [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)
- [helm](https://helm.sh/)
- minikube ou qualquer ferramenta de criação de cluster k8s
- [k6](https://k6.io/) 

## Como rodar o projeto.
Inicialmente será necessário iniciar um cluster K8s, para fazer isso você pode usar a sua ferramenta de escolha, no caso deste projeto foi escolhido usar [minikube](https://minikube.sigs.k8s.io/docs/) uma ferramenta de criação e gerenciamento de clusters K8s.

Comece rodando o comando `minikube addons enable metrics-server` para habilitar o sistema de métricas no k8s, que é necessário para o autoscaler funcionar

Comece criando um cluster rodando o comando:
``minikube start``
Para garantir que o cluster está rodando corretamente você pode rodar o comando `kubectl get nodes` se o comando listar pelo menos um nó o cluster está rodando corretamente.
Com o cluster rodando, rode o comando `helm install my-app app` isto irá criar os recursos necessários para a aplicação dentro do cluster. Depois disso rode o comando `minikube tunnel`(preferencialmente rode isso em outro terminal) este comando irá expor o load balancer criado no cluster para a sua máquina, para que seja possível chamar a aplicação como um serviço local.

## Rodando Stress test
Já que este projeto tem como propósito fazer uma demo de algumas funcionalidades do k8s, foi implementado um stress test que força o auto scale do k8s a escalar a nossa aplicação subindo automaticamente o número de pods com a nossa imagem. Para rodar este teste será necessário ter a ferramenta k6(mencionada nos [Pré requisitos](#pré-requisitos)). Antes de rodar os testes rode o comando `export APP_PORT=$(kubectl service my-app -o jsonpath='{.spec.ports[0].nodePort}')` e garanta também que o minikube tunnel está rodando. Para execução basta rodar a ferramenta `k6 run -e PORT=$APP_PORT script.js`.
