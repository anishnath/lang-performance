const express = require('express');
const app = express();

// Define Message struct
class Message {
    constructor(text) {
        this.text = text;
    }
}

// Define FactorialResponse struct
class FactorialResponse {
    constructor(input, result) {
        this.input = input;
        this.result = result;
    }
}

// Function to compute factorial
function factorial(n) {
    if (n <= 1) {
        return 1;
    }
    return n * factorial(n - 1);
}

// Ping handler
app.get('/ping', (req, res) => {
    const message = new Message('pong');
    res.json(message);
});

// Factorial handler
app.get('/factorial/:number', (req, res) => {
    const input = parseInt(req.params.number);
    if (isNaN(input)) {
        res.status(400).json({ error: 'Invalid input' });
        return;
    }

    const result = factorial(input);
    const response = new FactorialResponse(input, result);
    res.json(response);
});

// Start the server
const port = 18083;
app.listen(port, () => {
    console.log(`Server started on port ${port}`);
});