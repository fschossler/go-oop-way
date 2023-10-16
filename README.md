# go-oop-way

Some implementations of Go OOP supported principles. Remembering that Go is not a OOP language. It support and uses many programming paradigms.

## Exercises

### Exercise 1: Create a Struct

Define a struct named `Person` with fields for `Name` (string), `Age` (int), and `Address` (string).

### Exercise 2: Create a Method

Add a method to the `Person` struct that returns a formatted string with the person's details.

### Exercise 3: Composition

Define a struct named `Employee` that contains a `Person` as a field. The `Employee` should also have fields for `EmployeeID` (int) and Salary 
(float64).

### Exercise 4: Interfaces

Define an interface called `Drawable` with a `Draw` method that returns a string. Implement this interface in the `Person` and `Employee` structs from Exercises 1 and 3.

### Exercise 5: Polymorphism

Create a function that takes a slice of objects that implement the `Drawable` interface. The function should iterate through the slice and call the `Draw` method for each object.

### Exercise 6: Delegation (Composition)

Define a struct called `Calculator` with fields for two numbers and methods for `addition`, `subtraction`, `multiplication`, and `division`. Create another struct called `AdvancedCalculator` that contains a `Calculator` as a field. The `AdvancedCalculator` should have additional methods like square root, power, and logarithm, **but these methods should delegate the basic arithmetic operations to the embedded `Calculator` struct**.