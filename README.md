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
1. **The lexer**
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

2. **The parser**
3. **The Abstract Syntax Tree (AST)**
4. **The internal object system**
5. **The evaluator**


