const request = require('supertest');
const express = require('express');

const app = express();
app.use(express.json());

function isDiagonal(matrix) {
    if (!matrix || matrix.length === 0) return false;
    for (let i = 0; i < matrix.length; i++) {
        for (let j = 0; j < matrix[0].length; j++) {
            if (i !== j && Math.abs(matrix[i][j]) > 1e-6) {
                return false;
            }
        }
    }
    return true;
}

app.post('/analyze-matrix', (req, res) => {
    const { rotated_matrix } = req.body;
    if (!rotated_matrix || rotated_matrix.length === 0) {
        return res.status(400).json({ error: 'Matriz vacía o inválida' });
    }
    const values = rotated_matrix.flat();
    const max = Math.max(...values);
    const min = Math.min(...values);
    const sum = values.reduce((acc, curr) => acc + curr, 0);
    const avg = sum / values.length;

    return res.json({
        max_value: max,
        min_value: min,
        average: avg,
        total_sum: sum,
        is_diagonal: isDiagonal(rotated_matrix)
    });
});

describe('Pruebas Unitarias y de Integración en Node.js', () => {
    test('Verifica correctamente si una matriz es diagonal', () => {
        const diagonalMatrix = [[1, 0], [0, 2]];
        const nonDiagonalMatrix = [[1, 3], [0, 2]];
        expect(isDiagonal(diagonalMatrix)).toBe(true);
        expect(isDiagonal(nonDiagonalMatrix)).toBe(false);
    });

    test('POST /analyze-matrix calcula estadísticas correctamente', async () => {
        const response = await request(app)
            .post('/analyze-matrix')
            .send({ rotated_matrix: [[1, 2], [3, 4]] });

        expect(response.statusCode).toBe(200);
        expect(response.body.max_value).toBe(4);
        expect(response.body.min_value).toBe(1);
        expect(response.body.total_sum).toBe(10);
        expect(response.body.average).toBe(2.5);
    });

    test('POST /analyze-matrix responde 400 ante entrada inválida', async () => {
        const response = await request(app)
            .post('/analyze-matrix')
            .send({ rotated_matrix: [] });

        expect(response.statusCode).toBe(400);
    });
});