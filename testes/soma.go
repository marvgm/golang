//ESte é o pacote de exemplo de teste
package testes

//Teste para a função Soma
func Soma(a int, b int) int {
	if a > 0 {
		return a + b
	}
	return a + 0
}

//Fibboncia func para benchmark
func Fib(n int) int  {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
