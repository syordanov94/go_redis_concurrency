# Go-Redis Concurrency Lab

This is a project, written in **Go**, that demostrates various solutions to a specific concurrency problem, using **Redis**. 

The problem states that users want to buy stock shares for a specific company. The stock shares the company has are 1000 and each of the 30 clients buys 100 shares. These clients can buy shares simultaneously and the application should assure that all the available shares are bought (no more, no less).

To achive this, this project makes use of some concurrency solutions that can be applied to **Redis**.

- *Atomic Operators*
- *Transactions*
- *Locks*
- *LUA Scripts*


## Prerequisites

- Golang 1.19 or higher installed 
- Docker installed
- _Recomended but not mandatory_ VS Code or a similiar IDE 

## How to install and Run the project

- First you will have to clone the project from this github repository.

```bash
git clone https://github.com/syordanov94/go_redis_concurrency.git
```

- Once cloned, access the root of the project and innitiate the local **Redis** instance using **Docker**. This can be done buy running the following command:

```bash
docker compose up
```

- Now you can run the _main.go_ to verify that the project is running correctly. In order to select which concurrency control solution you want to test, you will need to pass an additional argument. This argument is a number (from 0 to 4) indicating each one:
    - 0 -> No Concurrency Control
    - 1 -> Atomic Operators
    - 2 -> Transactions
    - 3 -> LUA Scripts
    - 4 -> Redis Locks

Once you have selected one, simply run:

```bash
go run cmd/main.go {concurrencySolutionSelected}
```

for example:

```bash
go run cmd/main.go 2
```

## How to test the project