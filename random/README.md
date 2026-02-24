# random

Cryptographically secure random value generation.

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

- Cryptographically secure using `crypto/rand`
- Random integers with configurable range
- Random strings with optional special characters

## Installation

```bash
go get github.com/andmitr/pkg/random
```

## Usage

### func Int

```go
func Int(max int) (int, error)
```
Int returns a random integer between 0 and max (inclusive).

#### Example
```go
n, err := random.Int(100)
if err != nil {
    log.Fatal(err)
}
fmt.Println(n) // 0..100
```

### func MaxInt

```go
func MaxInt() (int, error)
```

MaxInt returns a random integer between 0 and math.MaxInt - 1.

#### Example

```go
n, err := random.MaxInt()
if err != nil {
    log.Fatal(err)
}
fmt.Println(n)
```

### func String

```go
func String(length int, specialChars bool) (string, error)
```

String generates a random string of the specified length. If specialChars is true, includes special characters in the output.

#### Example

```go
// Without special characters
s, err := random.String(16, false)
if err != nil {
    log.Fatal(err)
}
fmt.Println(s) // e.g. "aBcDeFgHiJkLmNoP"

// With special characters
s, err = random.String(16, true)
if err != nil {
    log.Fatal(err)
}
fmt.Println(s) // e.g. "aB#dE&gH!JkL$NoP"
```
## License
MIT Licensed. See [LICENSE](../LICENSE) for details.
