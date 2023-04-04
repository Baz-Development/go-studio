## Info

Se você rodar go vet no seu código, deve receber um erro similar ao seguinte:

```
sync/v2/sync_test.go:16: call of verificaContador copies lock valor: v1.Contador contains sync.Mutex
sync/v2/sync_test.go:39: verificaContador passes lock by valor: v1.Contador contains sync.Mutex
```

Uma rápida olhada na documentação do sync.Mutex nos diz o porquê:

`Um Mutex não deve ser copiado depois do primeiro uso.`

Mutex nos permite adicionar travas aos nossos dados
WaitGroup é uma maneira de esperar as goroutines terminarem suas tarefas.


## WaitGroup

```
Um WaitGroup aguarda por uma coleção de goroutines terminar seu processamento.
A goroutine principal faz a chamada para o Add definir o número de goroutines que serão esperadas. 
Então, cada uma das goroutines é executada e chama Done quando termina sua execução. 
o mesmo tempo, Wait pode ser usado para bloquear a execução até que todas as goroutines tenham terminado.
```

Ao esperar por wg.Wait() terminar sua execução antes de fazer nossas asserções, podemos ter certeza que todas as nossas goroutines tentaram chamar o Incrementa no Contador.