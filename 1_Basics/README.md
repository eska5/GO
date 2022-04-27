# Basics

Go is a statically typed, compiled programming language. Go syntax is simillar to C and Python.
Go is also known as Golang.

In this readme we will cover topics such as: Packages, Variables, Constants and Functions.

# Packages
Go applications are organized in packages. Packages are collections of files in the same directory.
if we want to use package from the directory for example `SomePackage/subPackage/ThePack`, we should declare `package ThePack`.
imports from package such as (functions, types, variables, constants) that start with a capital letter are `public and visible`, 
on the other hand when we won't declare them with a capital letter, they are `private and inaccessible`.
- camelCase -> `private and inaccessible`
- PascalCase -> `public and visible` 

Example:
```
package main
```
# Variables
Go is a statically-typed language, which means that all variables `must have a defined type` at compilation.

Variables can be defined by explicitly specifying a type:
```
var variableName int
```

We can also initialize our variable without giving a type, compiler will assing the variable type for us to match the given value.
```
variableName := 10
```

Assigning values are declared with `=` operator. 
Variables once declared are not capable of changing its type.

# Constants
Constant values are unchangeable. To define constant variable we use the `const` keyword.
```
const variableName = 1234
```

# Functions
Go functions can have 0 or more parameters. Parameters must have type declared when being passed for example:
```
func start(variable string){}
```

Values are returned from functions with a keyword `return`

To invoke a function such as start given above we simply type
```
start(variable)
```

Comments are declared in the same way as in c++ which means:
- single line -> `// code`
- multiple lines -> `/* code */`

Complete example of Go function:
```
// public function
func PublicFunction (variable string) string {
    return variable
}

// private function
func publicFunction (variable string) string {
    return variable
}
```