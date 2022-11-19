# Problem

Design and implement, using Golang, a JSON RPC server that provides big-number computation services for clients.

The service allows creating/updating/deleting named number objects that the server manages for clients.

Other than create/update/delete operations, your server should support basic arithmetic operations on these named number objects: addition, subtraction, multiplication, and division.

## Requirements
A client may perform computations like this:
 
### Create a number:
 
```
{"jsonrpc":"1.0","method":"create","params":["grav_const", "0.000000000066731039356729"],"id":1}
```
 
### Create another number:
 
```
{"jsonrpc":"1.0","method":"create","params":["planet_mass", "6416930923733925522307001.29472615"],"id":2}
```
 
### Request calculation results:
 
```
{"jsonrpc":"1.0","method":"multiply","params":["grav_const", "planet_mass"],"id":3}
```
 
```
{"jsonrpc":"1.0","method":"multiply","params":["planet_mass", "0.5"],"id":4}
```
 
### Update named number values:
 
```
{"jsonrpc":"1.0","method":"update","params":["grav_param", "428208470021099.94"],"id":5}
```
 
### Delete a named number:
 
```
{"jsonrpc":"1.0","method":"delete","params":["grav_param"],"id":6}
```

A number object may have multiple clients operating on it at any given time.
The server should be able to handle multiple concurrent requests and provide error information when an operation fails.

```
{"jsonrpc":"1.0","method":"create","params":["dayinmonth", "30"],"id":7}
{"jsonrpc":"1.0","method":"delete","params":["dayinmonth"],"id":8}
{"jsonrpc":"1.0","method":"update","params":["dayinmonth", “31”],"id":9} // ERROR!
// `dayinmonth` is already been deleted so it cannot be updated.
```
 
Implementation should be production-ready and maintainable. You should consider this task as the one you’ll be assigned during your work. Please implement your server with a mission-critical mindset and handle errors as much as possible. Anything else not mentioned you are free to use your imagination.

## Bonus
Try to implement the mechanism to differentiate between multiple users.

## Useful Resources

* https://pkg.go.dev/net/rpc/jsonrpc
* https://pkg.go.dev/math/big
* https://www.jsonrpc.org/specification_v1
