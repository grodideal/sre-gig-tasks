# Ga Array Operations
This Go program demonstrates basic array manipulation and iteration. It performs the following operations:

## 1. Menu Array
A string array named `Menu` is defined with the items "hamburger" and "salad". The program then iterates over the `Menu` array and prints each item with the format `Food: <Food name>.`
## Output:
```
Food: hamburger
Food: salad
```
## 2. Items Array
An array named `items` is defined with 5 items ("apple", "banana", "orange", "grape", "pear"). The program iterates over the `items` array and prints a message for each item in the format:
```
This is <ITEM> and its index in the array is <INDEX>
```
## Output:
```
This is apple and its index in the array is 0
This is banana and its index in the array is 1
This is orange and its index in the array is 2
This is grape and its index in the array is 3
This is pear and its index in the array is 4

```
## Running the Code
Ensure you have Go installed on your system. If not, you can download and install it from [the official Go website](https://golang.org/dl/)

1. Clone the repository to your local machine 
```
git clone https://github.com/grodideal/sre-gig-tasks.git
```
2. Navigate to the project directory:
```
cd sre-gig-tasks/task1
```
3. Run the Go program:
```
go run main.go
```
## When you run this code, it will produce the following output:
```
$ go run main.go
Food: hamburger
Food: salad
This is apple and its index in the array is 0
This is banana and its index in the array is 1
This is orange and its index in the array is 2
This is grape and its index in the array is 3
This is pear and its index in the array is 4
```