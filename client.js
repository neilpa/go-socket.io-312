#!/usr/bin/env node
io = require("socket.io-client")

now = () => new Date().toISOString().split('T')[1]

var addr = "http://localhost:5432";
console.log(`client: ${now()} connecting to ${addr}`)

var s = io(addr)

s.on('connect', function(d) {
  console.log(`client: ${now()} sending foo: hello`)
  s.emit('foo', 'hello');
})
s.on('bar', function(d) {
  console.log(`client: ${now()} recv'd bar: ${d}`)
})
