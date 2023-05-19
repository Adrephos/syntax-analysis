# Syntax analysis

This project is an implementation of two parsing algorithms for context-free languages, a top-down parser and a bottom-up parser.
The top-down parser is made for LL(1) grammars and uses predictive parsing to know if a given string is part
of the language the given grammar describes. In the other hand, the bottom-up algorithm is made for LR(0) grammars,
it first constructs the LR(0) automaton and then uses an SLR parser to know if a given string is part of the
language the grammar describes.

## Usage

To use this tool, each line on the input is a production rule and is of the form:
```
A -> α | β
```
With 'A' being a single non-terminal and 'α' and 'β' being strings of terminals and/or non-terminals symbols.
To represent the empty string 'ε' is used and every symbol of the grammar must be represented by a single character
or the program won't understand your grammar.

If you want to get the input with a file, the syntax should be as follows:
```
A -> α | β

w
w
---
S -> α | β

w
w
```
Where 'w' represents a string to parse, and each grammar is separated by '---'

## Build

To build this project simply clone the repository and build it using go:
```
git clone https://github.com/Adrephos/syntax-analysis
cd syntax-analysis
go build
```
Now you can execute the program:
- Manual input
```
./syntax-analysis
```
- File input
```
./syntax-analysis -f "path-to-file"
```
- Generate grammar
```
./syntax-analysis -g
```
- Help command
```
./syntax-analysis -h
```

