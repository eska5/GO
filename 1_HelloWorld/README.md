# Excercise 1.Hello World

--

# Instructions
The classical introductory exercise. Just say "Hello, World!".

"Hello, World!" is the traditional first program for beginning programming in a new language or environment.

The objectives are simple:

- Write a function that returns the string "Hello, World!".
- Run the test suite and make sure that it succeeds.
- Submit your solution and check it at the website.

# Solution
``` 
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

```

# Knowledge gained

We need to include our code to a package, so for now we used package main and created a function main. Which means that entrypoint of our application is a main function.
Package main creates a directory instead our Go container/project. Next line is import "fmt". To clarify this one we should start with import which simply imports packages. This time we 
are using "fmt" which stands for "format" to use `fmt.Println("Hello, World!")` function which returns `Hello, World!` in console. 
Next up is our `func main()` which is a main of our file in which we will be writing our code. We use brackets(`{ }`) to show where the function starts and where it ends. 
Finally when we run this file the output should be `Hello, World!`.


# Valuable Keywords
- package -> directory inside of our Go container
- import -> imports packages that we want to use
- func -> basically a function declaration
- "fmt" -> format package used to use `fmt.function()` for example `fmt.Println("Hello, World!")`