# Learning the Go

## Variables

Testing out the different variable declarations and initializations!

    $ go run variables.go
    initial
    1 2
    true
    0
    apple

## Functions

Let's see how functions work, aye?

    $ go run funcs.go
    1 + 2 = 3
    1 + 2 + 3 = 6
    3
    7
    7

## Encoding

Playing around with encoding a bit:

    $ go run encoding.go
    YWJjMTIzIT8kKiYoKSctPUB+
    abc123!?$*&()'-=@~
    
    YWJjMTIzIT8kKiYoKSctPUB-
    abc123!?$*&()'-=@~
    
This is neat because the URL-compatible base64 example (the second one) will use `-` instead of `+` by default!

## JSON Data

Toying with JSON a bit:

    $ go run json.go
    true
    1
    2.34
    "gopher"
    ["apple","peach","pear"]
    {"apple":5,"lettuce":7}
    {"Page":1,"Fruits":["apple","peach","pear"]}
    {"page":1,"fruits":["apple","peach","pear"]}
    map[num:6.13 strs:[a b]]
    6.13
    a
    {1 [apple peach]}
    apple
    {"apple":5,"lettuce":7}

This one was a real learning experience! I found out that a multi-object (such as `&response1{Page: 1, Fruits: []string{"..."}}` can't have the ending `}` on a new line... which is super annoying. Also, initializing a `byte` array uses parentheses for its value opposed to curly braces -- but maybe that's because it's not really a collection of multiple different things but just each byte of a single "string" represented as an array (i.e. `char foo[4] = "bar";` in C...).
