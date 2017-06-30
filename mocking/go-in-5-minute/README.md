# Mocking go-in-5-minutes example

[![Source code](../../src/reference.png)](https://github.com/arschles/go-in-5-minutes/tree/master/episode0)

Is showing a way to mock in go

## Example

go-in-5-minutes example implement a HashTable structure in two different modes. A in memory HashTable
and a second one using Redis Database.

The HashTable.go file holds an interface for methods that will be implemented in those to modes.

There will be a Get and a Set methods. Using the common interface golang allows to use an implementation
in testing and the other one in main file (production).

Same implementation can be used in both places, for integration testing.
