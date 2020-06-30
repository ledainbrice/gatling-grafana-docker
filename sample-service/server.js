const express = require('express')
const app = express()

const PORT = 8000;
var counter = 0;

app.get('/', (req, res) => {
  counter++
  setTimeout(function() {
    res.send('Hello World -> ' + counter);
  }, counter);
});

app.listen(PORT);
console.log(`Running on http://localhost:${PORT}`);