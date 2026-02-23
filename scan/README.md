# Scan

Functions for reading user input from stdin.

[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square;logo=opensource)](../LICENSE)

## Sponsorship

[![Boosty](https://img.shields.io/badge/Boosty-F15F2C?style=for-the-badge&logo=boosty&logoColor=white)![Support](https://img.shields.io/badge/Support%20me-grey?style=for-the-badge)](https://boosty.to/andmitr/donate) 

![Bitcoin](https://img.shields.io/badge/Bitcoin-F7931A?style=flat&logo=bitcoin&logoColor=white&logoSize=auto) 
```
1CCnwAvJYEoDVGM7vsBg2Q99cF9EHtBVaY
```

![Tether](https://img.shields.io/badge/Tether%20(USDT%20ETH)-168363?style=flat&logo=tether&logoColor=white&logoSize=auto) 
```
0x54f0ccc6b2987de454f69f2814fc9202bcfb74fe
```

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [License](#license)

## Features

- Simple single-line input reading
- Returns `io.EOF` when input is closed (Ctrl+D)
- Type conversion with error handling

## Installation

```bash
go get github.com/andmitr/pkg/scan
```

## Usage

### func String

```go
func String() (string, error)
```
String reads a single line from stdin and returns it as a string.

#### Example

```go
fmt.Print("Enter your name: ")
name, err := scan.String()
if errors.Is(err, io.EOF) {
    fmt.Println("Input closed")
    return
}
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Hello, %s!\n", name)
```

### func Int

```go
func Int() (int, error)
```
Int reads a single line from stdin and returns it as an int.

#### Example

```go
fmt.Print("Enter your age: ")
age, err := scan.Int()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("You are %d years old\n", age)
```

### func Float

```go
func Float() (float64, error)
```
Float reads a single line from stdin and returns it as a float64.

#### Example
```go
fmt.Print("Enter price: ")
price, err := scan.Float()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Price: %.2f\n", price)
```

## License
MIT Licensed. See [LICENSE](../LICENSE) for details.
