[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-wait-group/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-wait-group/blob/main/README.md)

## Wait Group
Method to synchronize ```go routines``` which would be the ```wait group```, start with the initial code:  
```go
...
func main() {
  var wait_group sync.WaitGroup // declaration
  wait_group.Add(2) // number of goroutines that will be running at the same time
  writePhrase("Hello world")
  writePhrase("Olá mundo")
}
...
```

Here we execute the ```writePhrase``` function with a loop to print the ```text``` 5 times, after that we go to the second execution of it, and when printing + 5 times "Hello world" the program ends.  
How would we make one function depend on the other to execute? How to do it in a faster, more dynamic way, but this time synchronizing them, so that they both finish before my program finishes executing with ```sync.WaitGroup```, go's default package.  
### Usando WaitGroup
```go
...
func main() {
	var wait_group sync.WaitGroup // declaration
	wait_group.Add(2) // number of goroutines that will be running at the same time
  // 2 wait group goroutines
//anonymous func
  go func() { // go routine
		writePhrase("Hello world")
		wait_group.Done() //when finished
	}()
	go func() {
		writePhrase("Olá mundo")
		wait_group.Done()
	}()

	wait_group.Wait() // wait for both to finish until you reach zero routines running

}
...
```
What we did here was create anonymous functions with ```goroutine```, when the function finishes ```wait_group.Done()``` it will do a "-1" in ```wait_group.Add(2)`` `, the two functions will be executed concurrently.
We will only be blocked when we reach ```wait_group.Wait()``` if there is still some routine in ```wait_group```, when it reaches zero.
