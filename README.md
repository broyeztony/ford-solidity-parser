```ford
.----------------.  .----------------.  .----------------. 
| .--------------. || .--------------. || .--------------. |
| |              | || |     ___      | || |     __       | |
| |              | || |    |_  |     | || |    \_ `.     | |
| |              | || |      | |     | || |      | |     | |
| |              | || |      | |     | || |       > >    | |
| |              | || |     _| |     | || |     _| |     | |
| |   _______    | || |    |___|     | || |    /__.'     | |
| |  |_______|   | || |              | || |              | |
| '--------------' || '--------------' || '--------------' |
 '----------------'  '----------------'  '----------------'
```

This repository is a fork of https://github.com/broyeztony/ford-lang-parser to include EVM semantics
as an attempt to provide an alternative to https://soliditylang.org/ and https://docs.vyperlang.org/en/stable/

# Run the unit-tests
```shell
❯ go test ./tests
```

# Print the AST for a program
Take a look at the example smart contract programs defined in `playground.ford`
Comment/Uncomment accordingly and then 
```shell
❯ go run main.go
```
from your terminal 

# Example
```
contract HelloWorld;
let greet = "Hello World!";

❯ ./run
```

will produce the below AST

```ford
  {
  "body": [
    {
      "declarations": [
        {
          "id": {
            "name": "greet",
            "type": "Identifier"
          },
          "initializer": {
            "type": "StringLiteral",
            "value": "Hello World!"
          },
          "type": "VariableDeclaration"
        }
      ],
      "type": "VariableStatement"
    }
  ],
  "name": "HelloWorld",
  "type": "Contract"
}

```

# Semantics
```ford
// variable declaration
let x = "not a number";

// function definition with explicit arguments declaration
def increment (x) {
    while x < 10 {
        x = x + 1;
    }
    return x;
}

// functions do not need to declare the argument list
// in case they don't, the caller's arguments are accessible through the `_` placeholder object
def square {
    // describe is a native function that outputs a map of key-value pairs of the caller arguments
    describe(_);

    return _.x * _.x;
}

// ObjectLiteral assignment
let A = { x: 1, y: 2, b: true, s: "hello" };

// Empty BlockStatement
{}

// variable initialization, with optional error handler following the error handler operator `->`
let result = square({ x }) -> {
    // This part here is an optional error handler.
    // It receives the `_` object which is an error object:
    // In this case ```{ code: INCOMPATIBLE_TYPE_ERROR, reason: "Incompatible type. Expected: 'Number', Found: 'String'."}```
    // The error handler let us return a 'recovery' value. Here, 0 will be assigned to the variable named `result`.
    // The error handler is not required to return a value.
    recover 0;
};

// 'if' statements, with 'else' alternative
if result > 1 {

}
else {

}

// calling a function, with error handler
result  = increment(0) -> {
    // describe the error object
    describe(_);
};

// Function's arguments can also be passed as a list like in the example below.
// In that case, they need to be accessed by index from the `_` implicit object.
def someOtherFunction {
    print(_[0]);
    print(_[1]);
}

someOtherFunction(1, 2);
```
