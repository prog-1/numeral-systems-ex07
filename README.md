# Numeral Systems

> Warning: It is not allowed to use any external long maths / big numbers libraries.

Write a program with two functions.

## Task 1: Machine Representation of Signed Integers

### Definition

Write a function that returns a machine representation (returned as a value of
type `string`) of a signed integer of type `int64`. The function signature:

```go
func NumToBin(int64 num) string
```

Write tests for the function.

### Test examples

```go
{0, "0"},
{1, "1"},
{8, "1000"},
{-1, "1111111111111111111111111111111111111111111111111111111111111111"},
{-1, "1111111111111111111111111111111111111111111111111111111111111111"},
{1000_0000_0000_0000, "11100011010111111010100100110001101000000000000000"},
{-1000_0000_0000_0001, "1111111111111100011100101000000101011011001110010111111111111111"},
```

## Task 2: Sum of Binary Numbers

Write a function that returns a sum of two integers that are encoded as base-2
values of type `string`. Assume that the integers are represented using 64 bits.
The function signature:

```go
func BinNumSum(string a, b) string
```

Write tests for the function.

### Test examples

```go
{"1", "11", "100"},
{"11001" /* 25 */, "1111111111111111111111111111111111111111111111111111111111110010" /* -14 */, "1011" /* 11 */},
{"1111111111111111111111111111111111111111111111111111111111110110", /* -10 */, "111" /* 7 */, "1111111111111111111111111111111111111111111111111111111111111101" /* -3 */},
{"11", /* 3 */, "1111111111111111111111111111111111111111111111111111111111111101" /* -3 */, "0" /* 0 */},
```
