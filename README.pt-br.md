[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-wait-group/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-wait-group/blob/main/README.md)

## Wait Group
Método pra sincronizar ```go routines``` que seria o ```wait group```, comece com o código inicial:  
```go
...
func main() {
  var wait_group sync.WaitGroup // declaração
  wait_group.Add(2) // quantidade de goroutines que vão estar rodando ao mesmo tempo
  writePhrase("Hello world")
  writePhrase("Olá mundo")
}
...
```

Aqui executamos a função ```writePhrase``` com um laço para imprimir o ```text``` 5 vezes, após isso vamos pra segunda execução dela, e ao printar + 5 vezes "Olá mundo" o programa termina.  
Como fariamos para que uma função dependa da outra para executar? Como fazer de um jeito mais dinâmico mais rápido, ams dessa vez sincronizando elas, para que as duas tenham terminado antes do meu programa terminar de executar com ```sync.WaitGroup```, pacote padrão do go.  
### Usando WaitGroup
```go
...
func main() {
	var wait_group sync.WaitGroup // declaração
	wait_group.Add(2) // quantidade de goroutines que vão estar rodando ao mesmo tempo
  // 2 goroutines do grupo de espera
	//func anonima 
  go func() { // go routine
		writePhrase("Hello world")
		wait_group.Done() //quando terminar
	}()
	go func() {
		writePhrase("Olá mundo")
		wait_group.Done()
	}()

	wait_group.Wait() // aguarda as duas acabarem

}
...
```
O que fazemos aqui foi criar funções anônimas com ```goroutine```, ao terminar a função ```wait_group.Done()``` fará um "-1" em ```wait_group.Add(2)```, as duas funções serão executadas concorrentemente.  
Nós só vamos ser barrados quando chegar em ```wait_group.Wait()``` caso ainda haja alguma rotina em ```wait_group```, quando ele chegar a zero.  
