<div align="center">
  <h1>
    <img src="https://raw.githubusercontent.com/ashleymcnamara/gophers/master/7TH_BIRTHDAY.png" width="192" /><br />
    Abstract
  </h1>
</div>

## Declaring functions

To declare a function in Go, you need to use the key-word `func` followed by its **name**, **list of arguments** and curly-braces to mark the **scope**:

```go
// receives two arguments of type int and the function itself returns an integer
func addTwoNumbers(a number, b: number) number {
  return a + B
}
```

The function's name also tells us about the function restrictions: those starting with a **lowercase letter are public** and those starting with a **uppercase letter are private**.

Also you can declare a function and declare a return variable which you can use inside your function body:

```go
func greetingPrefix(language string) (prefix string) {
  // Do whatever you want with `prefix` here and then you just need to return

  return
}
```

## How to write tests

Go already has a built-in testing tool which can be accessed by importing `testing`. However in order to make it work, you need to create a file called **xxxxx_test.go** and inside of it, you must create a funcion starting with `Test` prefix on its name.

After that, make sure of passing the following arguments: t \*testing.T

This `t` thing will allow you to access everything from go's testing tool.

If you need to split part of the code to create some helper function, you need to create it inside the Test funcion scope and receive the same `t *testing.T` argument. Inside of it, you must be explicit about that function being a helper one by using `t.Helper()`.

This `Helper` method will redirect the error to where the helper function is called and not its own scope.

To create subtests (remember `it` from jest), there's this `t.Run` method:

```go
func TestSubtraction(t *testing.T) {
  t.Run("a little description of what this is going to test exactly", func(t *testing.T) {
    // testing logic
  })
}
```

## Variables and Constants

Constants are declared the same way you do that in JavaScript:

`const nameOfTheConstant = "This is an example"`

Also they optimize your performance a little bit more than common variables.

To declare and initialize variables, you can use the operator `:=`.

`variableName := "Test"` this will automatically infer the type string to this variable.

Declaring variables without initialization will be _zero-valued_, which means that if you declare something like:

`var hey int`, this will be automatically starting with the int zero-value (0).

check [variables](https://gobyexample.com/variables).

## Conditionals

**if statements** are simple as you imagine:

```go
if something != otherThing {
  // Do this
}
```

yeah, you don't need parenthesis here.

**Switch statements** are also very simple:

```go
switch condition {
  case something:
    // Do something
  default:
    // do something if nothing is none of the expressions above are evaluated as true
}
```
