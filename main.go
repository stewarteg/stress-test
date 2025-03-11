package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	//entradinhas
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 1, "Número total de requests")
	concurrency := flag.Int("concurrency", 1, "Número de chamadas simultâneas")
	flag.Parse()

	if *url == "" {
		fmt.Println("A URL é obrigatória.")
		return
	}
	if *requests <= 0 || *concurrency <= 0 {
		fmt.Println("Os valores de requests e concurrency devem ser maiores que zero.")
		return
	}

	var total200, totalOther int
	statusCodes := make(map[int]int)
	var mu sync.Mutex

	// as porra da goroutines
	sem := make(chan struct{}, *concurrency)
	var wg sync.WaitGroup

	startTime := time.Now()

	for i := 0; i < *requests; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-sem }()

			resp, err := http.Get(*url)
			if err != nil {
				fmt.Println("Erro ao realizar o request:", err)
				return
			}
			defer resp.Body.Close()

			mu.Lock()
			statusCodes[resp.StatusCode]++
			if resp.StatusCode == 200 {
				total200++
			} else {
				totalOther++
			}
			mu.Unlock()
		}()
	}

	wg.Wait()
	totalTime := time.Since(startTime)

	fmt.Println("Relatório de Teste de Carga:")
	fmt.Printf("Tempo total gasto: %v\n", totalTime)
	fmt.Printf("Total de requests realizados: %d\n", *requests)
	fmt.Printf("Total de requests com status 200: %d\n", total200)
	fmt.Println("Distribuição de outros códigos de status HTTP:")
	for code, count := range statusCodes {
		fmt.Printf("Status %d: %d\n", code, count)
	}
}
