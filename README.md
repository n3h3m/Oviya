# Oviya Programming Language

A programming language designed for expressive and intuitive programming. With its unique blend of mathematical notation, reactive programming, and English-like syntax, Oviya aims to make code more readable and maintainable while providing powerful features for both beginners and advanced developers .

## Philosophy

Oviya is designed to be expressive like English, compact like math, and reactive like modern UI frameworks. It prioritizes:

- Readability over verbosity
- Expressiveness over boilerplate
- Reactive programming as a first-class citizen
- A syntax that matches how humans think, not just how machines run

## Features

### Core Language Features 

- **Dynamic Typing**: Variables can hold any type of value and be reassigned freely
- **Compiled language**: Compiles to C and the compiler was built on Go for performance and reliability
- **Mathematical Notation**: Natural mathematical expressions like `2x + y`
- **Method Chaining**: Elegant data processing with the `->` operator
- **Reactive Programming**: `when` statements for event-driven programming
- **Parallel Processing**: Background execution with `~` operator
- **English-like Syntax**: Readable conditions using words like `is`, `above`, `between`

### Advanced Features 

- **First-class Control Flow**: Assign `if`, `while`, `for` statements to variables
- **Complex Numbers**: Built-in support for complex numbers and quaternions
- **Time Arithmetic**: Natural time expressions like `3days from now`
- **String Interpolation**: Variable substitution in double-quoted strings
- **Array Operations**: Advanced slicing with `L:start:end:step:filter:map`
- **Autovivification**: Automatic creation of nested data structures
- **Closures & Decorators**: Functional programming capabilities

## Quick Start

Create your first Oviya program in a file named `hello.ov`:

```oviya
// Basic printing
print "Hello, Oviya!"

// Variables and arithmetic
x = 10
y = 5
result = 2x + y  // Natural mathematical notation
print "Result: {result}"

// Function definition
square(n) = n^2
print "Square of 7: {square(7)}"

// Method chaining
data = "  apple, banana, cherry  "
cleaned = data -> split(",") -> map(trim)
print cleaned  // ["apple", "banana", "cherry"]
```

## Language Syntax

### Comments 

```oviya
// Single line comment

///
Multi-line 
comments
///
```

### Literals
```
// Below are valid literals
t = 6pm // time literal
d = 27jun // date literal

```

### Function Literals
```
// Expressions like `+2`, `-2`, `/5`, `^2`, `*10`, `%=3`
// are treated as **anonymous functions** or **function literals**.
//
// These can be assigned to variables and used like regular functions.
// The `$` symbol is implicitly used as the argument.
//
// For example:
square = ^2        // same as: func(x) = x^2
double = *2         
half = /2          
add3 = +3           
mod5 = %5           

// They can be used in map, filter, reduce, etc.
nums = [1, 2, 3, 4, 5]
squares = nums.map(square)        // [1, 4, 9, 16, 25]
doubled = double nums        // [2, 4, 6, 8, 10]
doubled = nums -> double
even = nums.filter(%2 = 0)    // [2, 4]

// You can also assign and call them directly
f = -1
print f(10)    // 9
```

### Default arguement
```
// $ is usually considered as default arguement
// Upon receiving multiple arguments $ becomes the list of arguments
// In that scenario, Use `$1`, `$2`, `$3`, etc. for positional arguments in anonymous functions.
//
// Examples:
add = $1 + $2         // same as: func(a, b): return a + b
power = $1 ^ $2       // same as: func(base, exp): return base ^ exp
max = $1 > $2 : $1 : $2 // ternary operator cond:true_value:false_value

print add(3, 5)       // 8
print power(2, 3)     // 8
print max(10, 20)     // 20
```

### Variable Assignment 

```oviya
// Basic assignments
x = 5
name = "Alice"
active = true

// Constants (cannot be reassigned)
const PI = 3.14159
const MAX_USERS = 1000

// Numeric literals with commas
const POPULATION = 7,800,000,000

// Mathematical expressions
z = 10x + y    // Multiplication without explicit operator
const tau = 2pi // Using pi constant
```

## Data Types

Oviya supports a rich set of built-in data types:

    `null`, `bool`, `str`,
    `int`,  `ratio`, `real`, `complex`, `quat` ,
    `ts`, // timestamp
    `list`, `set`, `dict` ,
    `inter`, `seq`, `func`, // interval, sequence (generators)
    `url`, `email`, `err`


#### Comparison Operators
```oviya
a = 5
b = 10

// Standard comparisons
result1 = a = 5     // Equality (prints true)
result2 = b != 10   // Inequality
result3 = a  a     // Greater than
result5 = a = 10   // Greater than or equal

// Negated comparisons
result7 = a !> 3    // Not greater than
result8 = b != 15 // Not greater than or equal

// Double exclamation for equality
print b !! 10       // prints true (alternative equality)
print a!5          // prints false (inequality shorthand)
```


### Collection Examples 

```oviya
// Lists (1-indexed)
numbers = [1, 2, 3, 4, 5]
print numbers[1]  // Prints: 1

// Mixed-type arrays
mixed = [null, true, 42, "hello", [1,2], {a: 1}]

// Sets (ordered, unique values)
unique_nums = {1, 2, 3, 2, 1}  // {1, 2, 3}

// Maps (ordered key-value pairs)
person = {
    name: "John",
    age: 30,
    city: "New York"
}
person["occupation"] = "Developer"
```

### Complex Numbers and Quaternions 

```oviya
// Complex numbers
z1 = 2 + 3i
z2 = 1 - 4i
result = z1 * z2

// Quaternions
q = 1 + 2i + 3j + 4k
print q.r  // Real part: 1
print q.i  // i component: 2
print q.j  // j component: 3  
print q.k  // k component: 4

// Alternative notation when i,j,k are already taken ariables
z = 2 + 3i_
q = 1 + 2i_ + 3j_ + 4k_
```

## Functions

### Function Declaration 

```oviya
// Below ways to define functions

func f(n)
    print x

twice() = 2$ //$ stands for default arguement

twice(x) = 2x // Also valid

fib(n) = 0 if n < 1
         1 if n = 1
         fib(n-2) + fib(n+1)
         
fib(n) = 0 if n < 1 1 if n = 1 fib(n-2) + fib(n+1) // Also valid    
```

### Function Calls 

```oviya
// Three ways to calls functions
f(x)
f x
x -> f

// Traditional syntax
result1 = square(5)
result2 = fibonacci(10)

// Space-based calling (parentheses optional)
result3 = twice 5
result4 = fibonacci 10

// Multiple arguments
func add(a, b, c)
    return a + b + c

sum1 = add(1, 2, 3)
sum2 = add 1, 2, 3  // Same as above
sum2 = 1,2,3 -> sum // Also valid
```

### Typing
```
// Oviya allows struct which contains a declaration as below
// struct name must follow CamelCase while other variables must follow snake_case
// _ in struct name is not allowed, and uppercase in variable names are not allowed
// UPPER_CASE in variable name is allowed if all of them are upper case
// Variable names cannot start with or end with digits

// Struct defining syntax

struct StructName
    type field_name  validator_function
    
// Below are common validator functions, often times the type itself is a function, in generall in Oviya everything is functions in fact every single tokens are functions
// Enum: {'ACT', 'NSW', 'NT', 'QLD', 'SA', 'TAS', 'VIC', 'WA'}
// Less than or greater than, or between

// Struct Example
struct Address
    str street
    str city
    str state  {'ACT', 'NSW', 'NT', 'QLD', 'SA', 'TAS', 'VIC', 'WA'},
    int zip len = 4

struct Hotel    
    str name    len < 100
    int price   %50   
    email email
    url web
    Address address
    real rating 1<=$<=5
    
    
// Type will be enfored at the run time, also will be checked by static code analysers as well
// The below object will be failed to instantiated since Texas is not a valid austrlian state
hotel = Hotel {
    name: "Park Shereton",
    price: 270,
    email: "contact@shereton.com",
    web: "www.shereton.com",
    address: {
        street: "128 George St",
        city: "Sydney",
        state: "TX",
        zip: 2000
    }
}
// Output
// address.state is failing to be enum of {'ACT', 'NSW', 'NT', 'QLD', 'SA', 'TAS', 'VIC', 'WA'}


Here are some example values
x = 10             // int
name = "James"     // str
x = 1/3            // ratio


z = 2 + 3i         // complex
q = 4 + 4i + 3j + 7k // quaternion

// For complex and quaternions the variables i,j,k can be used, however Oviya allows redeclaring i,j as variables.
In Those circumstances you can declare them above as below

z = 2 + 3i_         // complex
q = 4 + 4i_ + 3j_ + 7k_ // quaternion

You can access them as usual
q.k // returns 3
q.r // would return 4

// Interval takes in this format
// start..end function
// If end is missing it's considered unbounded
// Function can be any function or a function literal
// In all practicality intervals and se


group_a = 1..10  // intergers    from 1 to 10
group_b = 1..100 +2
N = 1.. // natual numbers

odds = 1.. +2 // Set of odd numbers to infinity
odds = 1..100 +2 // Odd numbers under 100

squares = 1.. ^2 // infinite sequence of squares  1, 4, 9, 16...

// An interval is just a list with 2 items where the left one is less than right one
// interval can be any type of values, as longs as the values are comparable
bound = 1.2,,2.8
lunch = 1200,,1300

// For all purposes, intervals and sequnces can be iterated as lists, but can't be mutated as it would make the sqeunce rule controdictory

By design Oviya is a dynamically typed language, different types of variables can be assgined reassgined with different types of values.

However, for functions Oviya allows type hinting and those types are indeed enforced at runtime
x = 10
x = 'hello' // valid

int f(int x) = 2x // Optional hints, But they get enforced at runtime

// Partial type hinting is allowed
f(int a, b, str c)  // b has no type check, only `a` and `c` are enforced
    return "lorem ipsum" // return value not enforced


// Type casting works just like a function call
s = str 123    // s = "123"
n = int "456" // n = 456
r = ratio 0.75 // r = 3/4

// Treat types as functions
age_str = str(25)       // "25"
age_str = str 25        // "25"
age_str = 25 -> str     // "25"

```


### Type Hints 

```oviya
// Function with type hints (enforced at runtime)
int multiply(int a, int b) = a * b

// Partial type hinting
func process(int count, data, str format)
    // Only 'count' and 'format' are type-checked
    return format_data(data, count, format)

// Type casting
age_str = str 25        // "25"
age_num = int "25"      // 25
fraction = ratio 0.75   // 3/4

// Treat types as functions
age_str = str(25)       // "25"
age_str = str 25        // "25"
age_str = 25 -> str     // "25"

```

## Control Flow

### Conditional Statements 

```oviya
// Traditional if-else
age = 18
if age >= 18
    print "Adult"
else
    print "Minor"

// English-like conditions
// `is` `are` keyword often doesn't carry much meaning, just a place holder to make code like natural language
temperature = 75
if temperature is above 70 and time is after 6pm
    print "Perfect evening weather"
    
// The above is actually the below implementation 
temperature = 75
if temperature > 70 & time > 6pm
    print "Perfect evening weather"

// Multiple conditions
if all of username, password, email are not null
    create_account(username, password, email)
```

### Loops 

```oviya
// Numeric for loop
for 5 i
    print i  // Prints 1, 2, 3, 4, 5

// Range-based loops
for 1..10 i
    print "Number: {i}"

// Collection iteration
fruits = ["apple", "banana", "cherry"]
for fruits fruit
    print "I like {fruit}"

// Dictionary iteration
person = {name: "Alice", age: 30}
for person key, value
    print "{key}: {value}"
```

### `when` keyword

```oviya
// when key word
// Keep watching a variable and executes the body upon condition met
// when is active only on the block it was declared
// goes out of scope when the declared body goes out of scope

// Basic when statement
x = 5
when x > 90
  print "🔥 CPU overheating!",
```

### Parallel Processing 

```oviya
// Background function execution
func slow_task(name)
    sleep(2000)  // 2 seconds
    print "Task {name} completed"

// Non-blocking calls
~slow_task("A")
    // This body will run after successful completion
    print "Reached here After A is completed"
~slow_task("B") 
~slow_task("C")

// Parallel loops
urls = ["http://api1.com", "http://api2.com", "http://api3.com"]

// Sequential processing
for urls url
    ~fetch_data(url)  // Each call is non-blocking

// Parallel loop (waits for all to complete)
~for urls url
    fetch_data(url)
```

### Time and Date Arithmetic 

```oviya
// Time arithmetic
deadline = now + 3days
meeting = 2hours from deadline
reminder = 30minutes before meeting

// Duration calculations  
project_start = now
project_end = 6months after project_start
sprint_duration = 2weeks

// Time comparisons
if current_time is after deadline
    print "Project is overdue!"
    
if meeting_time is between 9am and 5pm
    print "Business hours meeting"
```

### Arrays

```oviya

//  Array notations
L = [1,2,3,4,5,6,7,8,9,10] // indexes start with 1

// Access array element
L[1] // returns 1
L[-1] // Returns 10

// Alternate notation to access element
(L)1 // returns 1
(L)-1 // returns 10
(L)5 //returns 5

// Use the mathematical notation if integer literal is available
X = [1,2,3,5]
print X1, X2 // prints 1,2

// Arrays can be splitted with a division operators

L/2 // returns [[1,2,3,4,5],[6,7,8,9,10]]

// The (paren)n notation would be useful in below situation

first_half = (L/2)1 // Returns [1,2,3,4,5], same as (L/2)[1]
first_third = (L/3)[1] // Returns [1,2,3] since L/3 = [[1,2,3],[4,5,6],[7,8,9],[10]]


// Array slicing
// L:start:end:step_function // Indices are inclusive
L:: // Returns L

// Array Slicing (inclusive, 1-based)

L = [1,2,3,4,5,6,7,8,9,10]

L::         // → [1,2,3,4,5,6,7,8,9,10]
L:1:5       // → [1,2,3,4,5]
L::5        // → [1,2,3,4,5]
L:8         // → [8,9,10]
L:8:10      // → [8,9,10]


I = [  1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

J = [-10,-9,-8,-7,-6,-5,-4,-3,-2, -1]
L = [  1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

// Negative indices are supported
L:-3:       // last 3 elements, [8,9,10]
L:-3:-1     // [8,9,10]

// L:start:end

//start..end are just interval,,, they are mathematical
// For ascending sequence increment 1 and descending decrement by 1 are automically supplied step function

// 10..1 is a valid sequence, which essentially reverses the list
L:10:1 //  [10, 9, 8, 7, 6, 5, 4, 3, 2, 1]

L:1:10:+2      // → [1,3,5,7,9]
L:10:1:-3      // → [10,7,4,1]
L:::^2      // → [1,4,9]

// Slices are first class
s1 = 1:5
s2 = 1:10:2

// Apply the slice like below
L:s1 // [1,2,3,4,5]
L:s2 // [1,3,5,7,9]

// Slicing also supports filtering
L:start:end:step_function:filter

// Slicing also supports mapping
L:start:end:step_function:filter_function:map_funtion

s3 = 1:10::is_prime:*2
L:s3 // [1,4,9,25,49]

```

### String Interpolation
```
// Literal strings (single quotes - no interpolation)
literal = 'Variables like {name} are not expanded here'
```

### Angle Bracket Functions

```oviya
// Predicate functions with natural language
// Must always return boolean and usually used in if, when, while

<very_high>(temp) = temp > 90
needs_review(comment) = comment contains "TODO"

// Usage in conditions
if temperature very_high
    print "Overheating!"
    
if comment needs_review
    print "Needs attention"

// Combining predicates
<concerning> = $ <very_high> or $ <urgent>
if temperature is not <concerning>
   print "OK"
```

### Autovivification 

```oviya
// Automatic creation of nested structures
config = {}
config["database"]["host"] = "localhost"
config["database"]["port"] = 5432
config["cache"]["redis"]["url"] = "redis://localhost"

print config
// Output:
// {
//   database: {host: "localhost", port: 5432},
//   cache: {redis: {url: "redis://localhost"}}
// }

// Array autofill
list = [1, 2, 3]
list[10] = "hello"  // Fills indices 4-9 with null
print list  // [1, 2, 3, null, null, null, null, null, null, "hello"] //null values don't cost memory
```

### Closures and Decorators 

```oviya
// Builtin suppport for closures

func counter()
    count = 0
    func inc()
        count++
        return count
    return inc

c = counter()
print c() // 1
print c() // 2

// Decorators are allowed syntax
func log(f)
    func inner(x)
        print "calling with", x
        return f(x)
    return inner

@log square(x) = x^2
square(4) // logs and returns 16

// Decorators can be chained
@memoise @log square(x) = x^2
```
### Rate Expressions (per)

```
// Oviya allows natural expressions using `per` keyword:
// e.g., 5000 per month, 40 km per hour, 25 items per box

// These are converted internally to: rate(value, unit)

// Units are first-class and composable:
// e.g., km/hour, dollars/month, people/km^2

salary = 5000 per month
speed = 100 km per hour
cost = 5 usd per kg

print salary * 12  // 60000 per year
print 200 km / speed  // → 2 hours

```

### Advanced Filtering

```oviya
// Multiple filter syntaxes
data = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

// Using $ as default variable
evens = data.filter($ % 2 = 0)

// Natural language filtering
large = data.filter(> 5)
medium = data.filter(between 3 and 7)
small = data.filter($ less than 4)
tiny = data.filter($ is less than 3)

// Combining operations
result = data
    -> filter(> 2)
    -> map(* 3)
    -> filter( sort
```

### Reactive Assignments

```oviya
// Reactive assignments automatically update when dependencies change:


// Basic reactive assignment
price = 100
tax_rate = 0.18
total => price + (price * tax_rate)

print total  // 118

// When price changes, total updates automatically
price = 150
print total  // 177 (automatically recalculated)

// Complex reactive expressions
base_salary = 50000
bonus_rate = 0.15
tax_rate = 0.25

gross_income => base_salary + (base_salary * bonus_rate)
net_income => gross_income - (gross_income * tax_rate)

print net_income  // Updates when any dependency changes
```

### String Formatting

```oviya
// Interpolated strings (double quotes)
name = "Alice"
age = 30
message = "Hello, {name}! You are {age} years old."

// Literal strings (single quotes)
template = 'Name: {name}, Age: {age}'  // Literals, no interpolation

// Multi-line interpolated
info = """
User: {name}
Age: {age}
Status: {status}
"""

// Multi-line literal
raw_template = '''
User: {name}
Age: {age}
Status: {status}
'''

// Escape braces for literals
literal_braces = "Literal braces: {{name}}"  // Output: Literal braces: {name}
```

## Examples

### Data Processing Pipeline 

```oviya
// Process CSV-like data
raw_data = """
  John Doe, 30, Engineer  
  Jane Smith, 25, Designer
  Bob Johnson, 35, Manager  
"""

processed = raw_data
    -> split("\n")
    -> map(trim)
    -> filter(len > 0)
    -> map(split(","))
    -> map(map(trim))
    -> map({
        name: $[1],
        age: int($[2]), 
        role: $[3]
    })

for processed p
    print "{p.name} ({p.age}) - {p.role}"
    
// Inside functional contexts (like map, filter), object literals are auto-wrapped into lambda functions
// The symbol `$` refers to the input element

map({
    name: $[1],
    age: int($[2]),
    role: $[3]
})

// is equivalent to

map(func f(x)
 return {
    name: x[1],
    age: int(x[2]),
    role: x[3]
})
```

## Glossary

- `$` — default parameter inside lambdas and inline functions  
- `->` — pipe operator: passes output of one function into the next  
- `=>` — reactive assignment: updates automatically when dependencies change  
- `<predicate>` — natural-language-style boolean function  
- `~` — non-blocking / background execution  
- `when` — reactive trigger block that runs on condition or change  

## Get Started

Oviya is a compiled language that compiles to C. The compiler itself is written in Go for performance and portability.

To get started:

```bash
# Clone the repository
git clone https://github.com/n3h3m/Oviya.git

# Navigate to the project directory
cd Oviya

# Build the Oviya compiler (requires Go)
go build -o oviya ./cmd/oviya

# Compile your .ov file to C and build it
./oviya build examples/hello.ov

# Run the compiled binary
./hello

```

**Maintainers**: Nehemiah Jacob
**Repository**: [github.com/n3h3m/oviya](https://github.com/n3h3m/oviya)

[![](https://img.shields.io/badge/Made_with-Go-informational?style=flat)](https://golang.org)
****