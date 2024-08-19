# Notes

## Extending to support utf-8 encoding

- If we want to eventually use this for json parsing we will eventually need to handle unicode values beyond ASCII (the first 128 bytes), the json spec: https://www.json.org/json-en.html
- Go a string is a sequence of bytes, go source code is utf-8 , utf-8 is a way to encode characters that varies in length, one to four bytes that are called code points. In go, these are called runes.
  - https://en.wikipedia.org/wiki/UTF-8#cite_ref-code-point-count_2-0
  - Utf8 has a simple encoding strategy that allows for variable length of bytes, the first byte always says how many bytes there are to be represented a code point in unicode. They are signified by the left most bits being set to 1.
    - For code points with only one byte: it follows the ASCII encoding scheme
    - For code points with 2 bytes wide the first byte will have 11xxxxxx , where xs are either 1 or 0 used to represent a digit in the code point scheme, and the next byte has 10xxxxxx, where the 10 signifies a continuation byte
    - For code points with 3 bytes wide the first byte will have 111xxxxx, where the remaining bytes and 0s are used to represent a digit in code points, and it is followed by a continuation byte
    - For code points with 4 bytes wide the first byte will have 1111xxxx, where the remaining bytes and xs are used to represent a digit that has a code point, followed by continuation bytes
  - This is better that utf-32 since it is generally more backwards compatible since it supports ascii, does not waste space and sometimes when computers encountered a byte with all zeroes that signified the end of the message, utf-8 avoids that.

## Parsers

- databases have a lexer and parser component, like for example Postgres uses the flex (lexer) and bison (parser compiler/generator). bison was mentioned in the book
- expressions produce values, statements don't
  - what is a statement? a statement is like a declaration, an aliasing of a value in another name... much like in `let x = 5` is a statement saying x is 5.
  - how a programming language handles this is language dependent, some languages have functions as expressions and can be use where expressions are allowed. Whereas some languages have don't treat functions as expressions, they are treated as statements purely.
- we are constructing a tree of nodes where the root nodes is the program
- we are doing a recursive descent parser
