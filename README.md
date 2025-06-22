# Oviya Programming Language

**A dynamic, interpreted programming language with natural syntax and reactive programming features**

- [Overview](#overview)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Language Features](#language-features)
- [Syntax Reference](#syntax-reference)
- [Data Types](#data-types)
- [Functions](#functions)
- [Control Flow](#control-flow)
- [Collections](#collections)
- [Time and Dates](#time-and-dates)
- [Method Chaining](#method-chaining)
- [Reactive Programming](#reactive-programming)
- [Background Processing](#background-processing)
- [Advanced Features](#advanced-features)
- [Examples](#examples)

## Overview

Oviya is a modern, dynamic programming language designed for simplicity and expressiveness[1][2]. It combines familiar programming concepts with innovative features like reactive assignments, natural language operators, and powerful method chaining capabilities[3][4]. The language is currently interpreted and implemented in Go, with plans for compilation in future versions[1].

### Key Features

- **Dynamic typing** with intuitive type inference
- **Natural language syntax** supporting English-like expressions
- **Reactive programming** with automatic variable updates
- **Method chaining** with pipeline operators
- **Background processing** with simple async syntax
- **Comprehensive time handling** with built-in temporal operations
- **Flexible control structures** as first-class objects

## Installation

### Prerequisites

- Go 1.19 or higher
- Git

### Building from Source

```bash
git clone https://github.com/n3h3m/oviya.git
cd oviya
go build -o oviya cmd/main.go
```

### Running Oviya

```bash
# Run a file
./oviya script.ov

# Interactive mode
./oviya
```

## Quick Start

Create your first Oviya program:

```oviya
// hello.ov
print "Hello, World!"

// Variables and basic operations
name = "Oviya"
version = 1.0
print "Welcome to {name} v{version}"

// Function definition
greet(name) = "Hello, {name}!"
print greet("Developer")
```

Run it:
```bash
./oviya hello.ov
```

## Language Features

### File Extension
Oviya source files use the `.ov` extension[5].

### Comments
```oviya
// Single line comment

///
Multi-line
comment block
///
```

### Print Statement
```oviya
print "Hello, World!"
print "Value: {variable}"  // String interpolation
```

## Syntax Reference

### Variable Declarations

```oviya
// Basic assignment
x = 5
y = 7
name = "Alice"

// Variables can be reassigned
y = 10

// Constants (cannot be reassigned)
const PI = 3.14159
const MAX_SIZE = 1,000,000  // Commas allowed in numeric literals
```

### Operators

#### Arithmetic Operators
```oviya
a = x + y      // Addition
b = a - x      // Subtraction
c = x * y      // Multiplication
d = x / y      // Division

// Implicit multiplication
z = 10x + y    // Equivalent to z = 10 * x + y
const tau = 2pi  // Equivalent to tau = 2 * pi
```

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

## Data Types

Oviya supports the following built-in types:

- `null` - Represents absence of value
- `bool` - Boolean true/false
- `int` - Integer numbers
- `float` - Floating-point numbers
- `timestamp` - Date and time values
- `string` - Text data
- `list` - Ordered collections
- `set` - Unique, ordered collections
- `map` - Key-value pairs

### Type Examples

```oviya
// Basic types
nothing = null
flag = true
count = 42
price = 19.99
text = "Hello"

// Complex types
numbers = [1, 2, 3, 4]
unique_items = {1, 2, 3}
dictionary = {1: "one", 2: "two", 3: "three"}

// Mixed-type collections
mixed_list = [null, true, 1, 4.5, [1, 2], {1, 2, 3}, {"key": "value"}]
```

## Functions

### Function Declaration

```oviya
// Standard function
func calculate_area(length, width):
    return length * width

// Shorthand function
area(length, width) = length * width

// Conditional function
fib(n) = 0 if n (temp) = temp is between 35 and 45
(comment) = comment contains "TODO"
```

### Function Calling

```oviya
// Standard calls
result1 = area(10, 5)
result2 = calculate_area(10, 5)

// Space-separated calls (syntactic sugar)
result3 = area 10, 5
result4 = fib 8

// Natural language function calls
if temperature is very_high:
    print "Too hot!"

if comment needs_review:
    print "Action required"
```

## Control Flow

### Conditional Statements

```oviya
// Basic if statement
if temperature > 30:
    print "It's hot!"

// Natural language conditions
if temperature is above 30 and humidity is below 60:
    print "Perfect weather"

if all of a, b, c are not null:
    print "All values present"

if any of x, y, z are above 100:
    print "High value detected"

// English operators
if score is between 80 and 90:
    print "Good score"

if filename ends_with ".ov":
    print "Oviya source file"
```

### Loops

```oviya
// Numeric loop
for 5 i:
    print i  // Prints 1, 2, 3, 4, 5

// Collection iteration
numbers = [1, 2, 3, 4, 5]
for numbers item:
    print item

// With condition
for numbers item:
    if item > 3:
        break
    print item
```

### First-Class Control Structures

```oviya
// Assign control structures to variables
loop = for numbers item:
    print item

condition = if temperature > 25:
    print "Warm"

// Activate later
loop:      // Execute the loop
condition: // Execute the condition
```

## Collections

### Lists (1-indexed)

```oviya
// List creation and access
numbers = [1, 2, 3, 4, 5]
print numbers[1]  // Prints 1 (first element)
print numbers[5]  // Prints 5 (last element)

// List modification
numbers[2] = 10   // Replace second element
print numbers     // [1, 10, 3, 4, 5]
```

### Sets (Ordered and Unique)

```oviya
// Set creation
unique_numbers = {1, 2, 3, 2, 1}  // Results in {1, 2, 3}
mixed_set = {1, 2, null}

// Set operations
unique_numbers.add(4)
unique_numbers.remove(1)
```

### Maps (Insertion-Ordered)

```oviya
// Map creation and modification
person = {
    "name": "Alice",
    "age": 30,
    "city": "Sydney"
}

// Dynamic addition
person["email"] = "alice@example.com"
person[4] = "fourth_item"

// Access
print person["name"]  // "Alice"
print person[4]       // "fourth_item"
```

## Time and Dates

Oviya provides intuitive time handling with natural language keywords:
Oviya provides intuitive time handling with natural language keywords:

### Time Keywords
- `now` - Current timestamp
- `day`, `days` - Day units
- `hour`, `hours` - Hour units
- `minute`, `minutes` - Minute units
- `second`, `seconds` - Second units
- `week`, `weeks` - Week units
- `month`, `months` - Month units
- `year`, `years` - Year units
- `decade`, `decades` - Decade units
- `century`, `centuries` - Century units

### Time Operations

```oviya
// Time calculations
deadline = now + 3days
meeting_time = now + 2hours + 30minutes
project_end = now + 6months

// Time comparisons
if deadline is before now:
    print "Deadline passed"

if meeting_time is after 6pm:
    print "Evening meeting"

// Complex time logic
if time is after 9am and time is before 5pm:
    print "Business hours"
```

## Method Chaining

The pipeline operator (`->`) enables powerful method chaining:

### Basic Chaining

```oviya
// String processing pipeline
input = "  apple, banana  ,  cherry  "
result = input -> trim -> lower -> split(",") -> map(trim)
// Result: ["apple", "banana", "cherry"]

// Numeric processing
numbers = [1, 5, 3, 8, 2]
processed = numbers -> filter(> 3) -> sort -> map(* 2)
// Result: [10, 16] (filtered > 3, sorted, then doubled)
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

## Reactive Programming

### Reactive Assignments

Reactive assignments automatically update when dependencies change:

```oviya
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

### Watch Expressions

Monitor variable changes and execute code when conditions are met:

```oviya
// Basic watch
x = 5
when x  100:
    print "x is very high"
when x  90: print "🔥 CPU overload",
    "memory": when memory_usage > 80: print "💾 Memory warning",
    "disk": when disk_usage > 95: print "💽 Disk full"
}

// Control individual watches
pause alerts["cpu"]
resume alerts["memory"]
cancel alerts["disk"]

// Add new watches dynamically
alerts["network"] = when network_latency > 1000:
    print "🌐 Network slow"
```

## Background Processing

Execute functions asynchronously with the `~` operator:

```oviya
// Synchronous function call
func process_data(data):
    // Time-consuming operation
    return analyzed_data

result = process_data(large_dataset)  // Blocks execution

// Asynchronous function call
~process_data(large_dataset)  // Non-blocking

// Batch async processing
urls = [
    "https://api1.com/data",
    "https://api2.com/data", 
    "https://api3.com/data"
]

for urls url:
    ~fetch_data(url)  // All requests run in parallel
```

## Advanced Features

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

### State Debugging

```oviya
// Debug variable states
x = 42
y = 17

state x        // Output: the value of x = 42
state x, y     // Output: the value of x = 42, y = 17

// Useful in complex algorithms
func complex_algorithm(data):
    processed = data -> transform
    state processed
    
    result = processed -> analyze
    state result
    
    return result
```

### Natural Language Functions

```oviya
// Define phrase functions
(priority) = priority starts_with "!!!"
(task) = task contains "ASAP" or task contains "urgent"
(deadline) = deadline is before now

// Usage in conditions
if task is urgent:
    print "🚨 Immediate attention required"

if deadline is overdue:
    print "⚠️ Past due date"

// Combine conditions
if task is high_priority and not deadline is overdue:
    print "✅ Priority task on schedule"
```

## Examples

### Complete Application Example

```oviya
// task_manager.ov - Simple task management system

// Task data structure
tasks = []

// Task priority functions
(task) = task["priority"] = "high" and task["due"] is before now + 1day
(task) = task["due"] is before now
(task) = task["status"] = "done"

// Create new task
func create_task(title, priority, due_date):
    task = {
        "id": tasks.length + 1,
        "title": title,
        "priority": priority,
        "due": due_date,
        "status": "pending",
        "created": now
    }
    tasks.append(task)
    return task

// Get filtered tasks
get_urgent_tasks() = tasks.filter(urgent)
get_overdue_tasks() = tasks.filter(overdue)
get_pending_tasks() = tasks.filter(not completed)

// Task statistics (reactive)
total_tasks => tasks.length
pending_count => get_pending_tasks().length
urgent_count => get_urgent_tasks().length
overdue_count => get_overdue_tasks().length

// Monitoring
when urgent_count > 5:
    print "🚨 Too many urgent tasks: {urgent_count}"

when overdue_count > 0:
    print "⚠️ {overdue_count} overdue tasks need attention"

// Sample usage
create_task("Complete documentation", "high", now + 2days)
create_task("Review code", "medium", now + 1week)
create_task("Fix critical bug", "high", now + 6hours)

// Display statistics
print "Task Summary:"
print "Total: {total_tasks}"
print "Pending: {pending_count}"
print "Urgent: {urgent_count}"
print "Overdue: {overdue_count}"

// Process urgent tasks
urgent_tasks = get_urgent_tasks()
for urgent_tasks task:
    print "🔥 URGENT: {task['title']} (due: {task['due']})"
```

### Data Processing Pipeline

```oviya
// data_analysis.ov - Process and analyze data

// Sample dataset
sales_data = [
    {"product": "Laptop", "price": 1200, "quantity": 5, "date": now - 1day},
    {"product": "Mouse", "price": 25, "quantity": 20, "date": now - 2days},
    {"product": "Monitor", "price": 300, "quantity": 8, "date": now - 1day},
    {"product": "Keyboard", "price": 75, "quantity": 15, "date": now - 3days}
]

// Processing functions
calculate_revenue(sale) = sale["price"] * sale["quantity"]
is_recent(sale) = sale["date"] is after now - 2days
is_high_value(sale) = sale["price"] > 100

// Reactive analytics
total_revenue => sales_data
    -> map(calculate_revenue)
    -> sum

recent_sales => sales_data
    -> filter(is_recent)
    -> length

high_value_revenue => sales_data
    -> filter(is_high_value)
    -> map(calculate_revenue)
    -> sum

// Real-time monitoring
when total_revenue > 10000:
    print "🎉 Revenue target exceeded: ${total_revenue}"

when recent_sales > 10:
    print "📈 High sales activity: {recent_sales} recent sales"

// Async reporting
~generate_detailed_report(sales_data)
~send_daily_summary(total_revenue, recent_sales)

// Display results
print "Sales Analytics:"
print "Total Revenue: ${total_revenue}"
print "Recent Sales: {recent_sales}"
print "High-Value Revenue: ${high_value_revenue}"

// Top products by revenue
top_products = sales_data
    -> map(λ sale: {
        "product": sale["product"],
        "revenue": calculate_revenue(sale)
    })
    -> sort_by("revenue", descending: true)
    -> take(3)

print "\nTop Products:"
for top_products product:
    print "{product['product']}: ${product['revenue']}"
```


### Development Setup

```bash
# Clone the repository
git clone https://github.com/n3h3m/oviya.git
cd oviya

# Install dependencies
go mod download

# Run tests
go test ./...

# Build development version
go build -o oviya-dev cmd/main.go
```

---

**Maintainers**: Nehemiah Jacob
**Repository**: [github.com/n3h3m/oviya](https://github.com/n3h3m/oviya)