# Notes

- If we want to eventually use this for json parsing we will eventually need to handle unicode values beyond ASCII (the first 128 bytes), the json spec: https://www.json.org/json-en.html
- Go a string is a sequence of bytes, go source code is utf-8 , utf-8 is a way to encode characters that varies in length, one to four bytes that are called code points. In go, these are called runes.
  - https://en.wikipedia.org/wiki/UTF-8#cite_ref-code-point-count_2-0
  -
