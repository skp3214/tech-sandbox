const express =  require('express');
const app = express();
const port = 3000;

app.use(express.json());

app.use(express.urlencoded({ extended: true }));

app.get('/', (req, res) => {
    res.status(200).send('Welcome to the Go Lang Web Server!');
});

app.get('/get', (req, res) => {
    res.status(200).json({ message: 'GET request received' });
});

app.post('/post', (req, res) => {
    const data = req.body;
    res.status(200).send(data);
});


app.post('/postform', (req, res) => {
    const data = req.body;
    res.status(200).send(JSON.stringify(data));
});

app.listen(port, () => {
    console.log(`Server is running at http://localhost:${port}`);
});
