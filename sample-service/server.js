const express = require('express')
const app = express()

const PORT = 8000;
var counter = 0;

app.get('/', (req, res) => {
  counter++
  /*setTimeout(function() {
    res.send('Hello World -> ' + counter);
  }, counter % 10000);*/
  res.send('Hello World -> ' + counter);
});

app.get('/error', (req, res) => {
  res.status(500);
  res.send('Test error 500');
});

app.listen(PORT);
console.log(`Running on http://localhost:${PORT}`);