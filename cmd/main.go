package main

import (
	"context"
	"fmt"
	"go_redis_concurrency/internal/go_redis_concurrency/redis"
	"os"
	"strconv"
	"sync"
)

const (
	total_clients = 30
)

func main() {
	const companyId = "TestCompanySL"

	// --- (0) ----
	// Recover implementation method
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		panic("missing implementation method")
	}
	selectedImpl, err := strconv.Atoi(argsWithoutProg[0])
	if err != nil {
		panic(err)
	}
	redis.SelectedConcurrencyImplementation = redis.ConcurrencyImplementation(selectedImpl)

	// --- (1) ----
	// Get the redis config and init the repository
	repository := redis.NewRepository("0.0.0.0:20003")

	// --- (2) ----
	// Publish available shares
	repository.PublishShares(context.Background(), companyId, 1000)

	// --- (3) ----
	// Run concurrent clients that buy shares
	var wg sync.WaitGroup
	wg.Add(total_clients)

	for idx := 1; idx <= total_clients; idx++ {
		userId := fmt.Sprintf("user%d", idx)
		go repository.BuyShares(context.Background(), userId, companyId, 100, &wg)
	}
	wg.Wait()

	// --- (3) ----
	// Get the remaining company shares
	shares, err := repository.GetCompanyShares(context.Background(), companyId)
	if err != nil {
		panic(err)
	}
	fmt.Printf("the number of free shares the company %s has is: %d", companyId, shares)
}
