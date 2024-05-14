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
as an attempt to provide an alternative to https://soliditylang.org/ and https://docs.vyperlang.org/en/stable/ programming languages
to develop EVM-compatible smart contracts

# Run the unit-tests
```shell
❯ go test ./tests
```

# Print the AST for a program
Take a look at the example smart contracts defined in `playground.ford`
Comment/Uncomment accordingly and then 
```shell
❯ go run main.go
```
from your terminal to generate the corresponding program's Abstract Syntax Tree.

# Example
```
contract HelloWorld;
let greet = "Hello World!";

❯ ./run
```

will produce the following AST

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

# Ford Semantics
Ford is a dynamically-typed language and it does not support Object-oriented programming.
In order to deal with variable's type, we resort to a few `reserved` functions like `address("0x...")` or `u8(255)`.

Every `.ford` contract must have a companion metadata file where state variable's types are declared explicitly.
This is also where functions' parameters and return types, visibility, state mutability are declared, as well as event's names and parameters. 

```ford
// contract's name
contract Semantics;

// state variable declaration
let x = "hello"; // a string
let y = u8(0); // a uint8 
let z = address("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"); // an ethereum address

// public function definition with explicit arguments declaration
def increment (b) {
    while y < 10 {
        y = y + b;
    }
    return y;
}

// functions can omit parenthesis and argument's list even when they accept parameters.
// in that case, the caller's arguments are accessible through the `_` placeholder object
def square {
    // describe is a native function that outputs a map of key-value pairs of the caller arguments
    describe(_);

    return _.y * _.y;
}

def demo() {
    // ObjectLiteral assignment
    let A = { x: u8(1), y: u8(2), b: true, s: "hello" };
    
    // Empty BlockStatement
    {}
    
    // variable initialization
    let result = square({ y });
    
    // 'if' statements, with 'else' alternative block
    if result > 1 { }
    else {}

    // calling functions    
    result  = increment(0);
    someOtherFunction(1, 2);    
}

// Function's arguments can also be passed as a list like in the example below.
// In that case, they need to be accessed by index from the `_` placeholder implicit object.
def someOtherFunction {
    print(_[0]);
    print(_[1]);
}
```

# Examples

### Primitive types
```ford
contract PrimitiveTypes;

let aString = ""; // string
let aBool = true; // boolean
let aUint8 = u8(255); // uint8
let anAddress = address("0xCA35b7d915458EF540aDe6068dFe2F44E8fa733c"); // address
```

```yaml
state:
  - name: aString
    type: string
  - name: aBool
    type: boolean
  - name: aUint8
    type: u8
  - name: anAddress
    type: address

defs: {}
events: {}
```

### Functions
```ford
contract View;

let x = u8(1);

// by default, all defs are `public`
// `y` is of type `u8` (uint8)
// `addToX` return type is `u8` (uint8)
def addToX (y) {
    return x + y;
}
```

```yaml
state:
  - name: x
    type: u8

defs:
  - name: addToX
    parameters:
      y: u8
    returnType: u8
    visibility:
      - public
    stateMutability:
      - view

events: {}
```

### Events
```ford
contract Events;

myAddress = address("0xfe3091F63A0b0b1cf81ff53102434aa287aC5289");

def receive(amount) {

    let sent = send(myAddress, amount);
    if sent {
        emit("Received", { address: msg.sender, amount: amount, message: "Received some ether" });
    } else {
        emit("TransferError", { address: msg.sender, amount: amount, message: "Transfer failure" });
    }
}
```

```yaml
state:
  - name: myAddress
    type: address
    payable: true

defs:
  - name: receive
    parameters:
      amount: u256
    returnType: null
    visibility:
      - public
    payable: true

events:
  Received:
    address: address
    amount: u256
    message: string
  TransferError:
    address: address
    amount: u256
    message: string
```

### Mappings
```ford
contract Mappings;

// string => bool
let m1 = mapping();
// string => (string => address)
let m2 = mapping();

def constructor {

    m1["abc"] = true;
    m1["def"] = false;

    m2["X"]["Y"] = address("0xfe3091F63A0b0b1cf81ff53102434aa287aC5289");
}
```

```yaml
state:
  - name: m1
    type: mapping(string => bool)
  - name: m2
    type: mapping(string => mapping(string => address))

defs:
  - name: constructor
    returnType: null
    visibility:
      - public

events: {}
```

### Lists
```ford
contract Lists;

def doList {

	let myList = listU8(1, 1, 2, 3, 5);

    let itemAtIndex3 = getItem(myList, 4); // i.e `5`
	append(myList, u8(8)); // add u8(8) at the end of the list
	remove(myList, 0); // removes u8(1) at the beginning of the list
	length(myList); // returns the size of the list, i.e `5`

	/* introduce utility functions
	let doesListContainElement = false;
	let needle = u8(4);
    for (i = 0; i < length(myList); i = i + 1) {
        if getItem(myList, i) == needle {
            doesListContainElement = true;
            break;
        }
    }
	*/
	let found = contains(myList, u8(2)); // false
}
```

```yaml
state: {}

defs:
  - name: doList
    returnType: null
    visibility:
      - private

events: {}
```
