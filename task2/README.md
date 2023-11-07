# Random Number Generator and Checker
This Go program generates a random number between 1 and 100 and performs various checks based on the generated number. The code is organized into functions for modularity and readability.
## Code Explanation
### 1. `generateRandomNumber` Function
The `generateRandomNumber` function generates a random number between 1 and 100 using the `rand` package and returns it.
### 2. `checkNumber` Function
The `checkNumber` function takes a number as input and performs the following checks:

- If the number is 50, it prints "It's 50!".
- If the number is higher than 50 and even, it prints "It's closer to 100, and it's even!".
- If the number is higher than 50 (but not even), it prints "It's closer to 100".
- If the number is lower than 50, it prints "It's closer to 0".

## Running the Code
Ensure you have Go installed on your system. If not, you can download and install it from [the official Go website](https://golang.org/dl/)

1. Clone the repository to your local machine 
```
git clone https://github.com/grodideal/sre-gig-tasks.git
```
2. Navigate to the project directory:
```
cd sre-gig-tasks/task2
```
3. Run the Go program:
```
go run main.go
```
## When you run this code, it will produce the following output:
```
$ go run main.go
Generated Random Number: 59
It's closer to 100
```