const express = require('express');
const axios = require('axios');
const app = express();

app.use(express.json());

const PORT = 3000;

// Endpoints  CRUD de filmes 
app.get('/filmes/:id', async (req, res) => {
    const { id } = req.params;
    try {
        const response = await axios.get(`http://localhost:8080/filme/${id}`);
        res.json(response.data);
    } catch (error) {
        res.status(500).send('Erro ao buscar filme');
    }
});

app.post('/filmes', async (req, res) => {
    try {
        console.log("Enviando requisição ao backend Go com dados:", req.body);
        const response = await axios.post('http://localhost:8080/filme', req.body);
        console.log("Resposta do backend Go:", response.data);
        res.status(201).json(response.data);
    } catch (error) {
        console.error("Erro ao adicionar filme:", error);
        res.status(500).send('Erro ao adicionar filme');
    }
});

app.get('/filmes', async (req, res) => {
    try {
        const response = await axios.get('http://localhost:8080/filmes');
        res.json(response.data);
    } catch (error) {
        console.error("Erro ao obter filmes:", error);
        res.status(500).send('Erro ao obter filmes');
    }
});

app.put('/filmes/:id', async (req, res) => {
    const { id } = req.params;
    try {
        const response = await axios.put(`http://localhost:8080/filme/${id}`, req.body);
        res.json(response.data);
    } catch (error) {
        res.status(500).send('Erro ao atualizar filme');
    }
});

app.delete('/filmes/:id', async (req, res) => {
    const { id } = req.params;
    try {
        const response = await axios.delete(`http://localhost:8080/filme/${id}`);
        res.status(204).send();
    } catch (error) {
        res.status(500).send('Erro ao deletar filme');
    }
});

app.listen(PORT, () => {
    console.log(`Servidor rodando em http://localhost:${PORT}`);
});
