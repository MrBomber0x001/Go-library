## Intro

functional programming is the process of buildign software by using `pure` functions, avoiding `shared` state, `mutable` data, and `side effects`

The thing to keep on your mind, is functions in Go are `first-class` functions

## simple overview of HOF

It can be:

- Passed as parameters
- Returened from other functions
to simplify this concept, I will try to write the equivalent js code

```go

// takes a callback function, and a parameter
func showName(f func(string), name string) {
    f(name)
}

func printInConsole(name string){
    fmt.Printf("The name is %s.\n", name)
}

func getClicker() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main(){
    // 1. function as parameter
    showName(printInConsole, "Forest")

    // 2; returning a function
    click := getClicker()
    fmt.Println(click()) // 1
    fmt.Println(click()) // 2
    fmt.Println(click()) // 3

}
```

## Pure functions

The output of a pure function depends entirely on the input and nothing else, so there is nothing which can change it's output
it's always going to return the same output for the same input
And there is no hidden state in between that can change it  

- **Idempotent** - it doesn't matter how many time we call the function
- Race conditions due to side effects are not possible with pure functions, so it makes pure functions a great feature to use in `concurrent` apps

**An example of pure vs impure functions**

| Pure | Impure |
| -- | -- |
| sha256.New() | rand.Intn() |
| len([]{1,3,4}) | time.Now() |
| math.sqrt(4) | fmt.Printf() |

```go
// pure
func isMinor(age int) bool {
    if age >= 18 {
        return false
    }
    return true
}

func main(){
    rand.Seed(time.Now().UTC().UnixNano())
    fmt.Printf("isMinor(20): %v\n", isMinor(20)) // pure
    fmt.Printf("rand.Intn(10): %v\n", rand.Intn(10)) // impure
}

```

## Encapsulting external state

## Closures

Anonymous function that references variables from outside it's body
The most trivial way to understand closures is the `click counter` example

- Are safer to use by isolating data
- Don't have to use global variables
- You can referene a variable without being passed as parameter and change its state

```go
func main(){
    clicker := newClick()
    fmt.Println(clickMe()) // 1
    fmt.Println(clickMe()) // 2
    fmt.Println(clickMe()) // 3
}
func newClick() int {
    click := 0
    return func() int {
        return ++click
    }
}
```

```go
func main(){
    ages := []int{20, 16, 14, 10, 24, 35, 50, 35, 36}
    sort.Ints(ages)
    fmt.Println("sorted ages: ", ages)

    index := sort.Search(len(ages), func(i int ) bool {
        return ages[i] >= 18
    })
    fmt.Println(">=18 years old:",  ages[index:])
}
```

TODO: understand how Search is working behind the scenes
