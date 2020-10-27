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
