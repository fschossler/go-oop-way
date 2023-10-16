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

### Exercise 4: Inheritance

Implement a simple inheritance example. Define a base struct called `Shape` with fields for `Color` (string) and `Area` (float64). Create 
two child structs, `Circle` and `Rectangle`, that inherit from `Shape`. Add specific fields and methods for each shape type.

### Exercise 5: Interfaces

Define an interface called `Drawable` with a `Draw` method that returns a string. Implement this interface in the `Circle` and `Rectangle` structs from Exercise 5.

### Exercise 6: Polymorphism

Create a function that takes a slice of objects that implement the `Drawable` interface. The function should iterate through the slice and call the `Draw` method for each object.

### Exercise 7: Delegation (Composition)

Define a struct called `Calculator` with fields for two numbers and methods for addition, subtraction, multiplication, and division. Create another struct called `AdvancedCalculator` that contains a `Calculator` as a field. The `AdvancedCalculator` should have additional methods like square root, power, and logarithm, but these methods should delegate the basic arithmetic operations to the embedded `Calculator` struct.

### Exercise 8: Abstraction

Define an abstract class called `Animal` with an `AnimalSound` method that should be implemented by specific animal types. Create two concrete animal types that inherit from Animal and provide their own implementations of AnimalSound.

### Exercise 9: Multiple Inheritance (Composition)

Define a struct called SmartPhone that combines the functionalities of a phone and a camera. Create methods to make calls, take photos, and describe the device's capabilities.