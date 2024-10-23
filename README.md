### Feature List:
- C-like syntax
- variable bindings
- integers and booleans
- arithmetic exp
- built-in functions
- first-class and higher-order functions 
- closures
- a string data structure
- an array data structure
- a hash data structure

### Parts of the project:
1. **The lexer (tokenizer)**  
  It will take source code as input and output the tokens that represent the 
  source code. It will go through its input and output the next token that it 
  recognizes. It doesn't need to buffer or save tokens, since there will be only 
  one method called `NextToken()`, which will output the next token.

- Ideally, it makes sense to think of adding something like `io.Reader` to get
filename and line-numbers as that would allow the interpreter to easily identify
and throw it during errors.

- When introducing identifiers and keywords, we need to recognize whether the 
current character is a letter or not and if so, continue reading the rest till 
a non-letter-character is encountered. After reading the entire identifier or 
keyword, we can then check the appropriate token type.

- You also need to account for whitespaces in your document so a function called 
`skipWhitespace()` or `eatWhitespace()` or `consumeWhitespace()` is found in a 
lot of parsers/lexers

- Next, we need to support different symbols and tokens
`==, !, !=, -, /, *, <, >, true, false, if, else, return`

2. **The parser (consumes tokens)**  
> A parser is a software component that takes input data (frequently text) and 
> builds a data structure - often some kind of parse tree, abstract syntax tree or 
> other hierarchial structure - giving a structural representation of the input, 
> checking for correct syntax in the process. The parser is often preceeded by a 
> separate lexical analyzer, which creates tokens from the sequence of input 
> characters.

The parser outputs an abstract syntax tree. It can look similar to this
```js
> var input = `if (3 * 5 > 10) { return "Hello"; } else { return "goodbye" }`;
> var tokens = MagixLexer.parse(input);
> MagicParser.parse(tokens);
```
```js
{
  type: "if-statement",
  condition: {
    type: "operator-expression",
    operator: ">",
    left: {
      type: "opeartor-expression",
      operator: "*",
      left: { type: "integer-literal", value: 3 },
      right: { type: "integer-literal", value: 5 }
    },
    right: { type: "integer-literal", value: 10 }
  },
  consequence: {
    type: "return-statement",
    returnValue: { type: "string-literal", value: "hello" }
  },
  alternative: {
    type: "return-statement",
    returnValue: { type: "string-literal", value: "goodbye" }
  }
}
```

Look into parser generators like yacc or bison. There are different ways to parse
things :
  1. top-down parsing (builds the root note first and then descends)
  2. bottom-up parsing
  3. recursive descent parsing (variant of top-down) ~ Pratt parser
  4. early parsing
  5. predictive parsing

- Starting to build our abstract syntax tree by parsing "let" statements only.
```js
// Expressions produce value, statements do not produce value
let <identifier> = <expression>
```
This is the usual structure of a "let" statement. We need to handle the 
statements and the expressions separately. Some nodes in our tree are going to 
be statements and others are going to be expressions

3. **The Abstract Syntax Tree (AST)**
4. **The internal object system**
5. **The evaluator**

