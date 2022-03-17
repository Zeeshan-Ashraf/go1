# go1
Covering all basics of Go by writing practical running code.

Prerequisite: basic knowledge of c programming language as go is very similar to c in syhtax & pointers

This repo go1 contains:
  1. Basic syntaxt of go viz: variables, pointer, basic data type, composite datatype array slice maps structure
  2. print the data using printf / sprintf
  3. function & methods
  4. how to do send get & post request to web with timeout and print the data to STDOUT
  5. build http server using gin (as well as using inbuild go library) to handle get post delete update APIs
  6. send local HTML files to user on hitting localhost:port/index.html
  7. handle json marshal (convert object to json) & unmarshal (convert recieved json from web to object), validate json, json tagging
  8. read OS environment variable 
  9. handle URL parameter URL query & URL grouping 
  10. go orm = gorm for DB connection & DAO layer to communicate with DB
  11. connect and read kafka
  12. add caching layer
  13. unit testing

How to run:
go to dir where is located main.go
run in terminal : go run *.go