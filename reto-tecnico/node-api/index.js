const express = require('express');
const swaggerUi = require('swagger-ui-express');
const swaggerDocument = require('./swagger.json');

const app = express();
app.use(express.json());

app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(swaggerDocument));

app.post('/analyze-matrix', (req, res) => {
    const { rotated_matrix } = req.body;

    if (!rotated_matrix || !Array.isArray(rotated_matrix) || rotated_matrix.length === 0) {
        return res.status(400).json({ error: 'Se requiere una matriz válida' });
    }

    let maxVal = -Infinity;
    let minVal = Infinity;
    let sumTotal = 0;
    let totalElements = 0;
    let isDiagonal = true;

    const rows = rotated_matrix.length;

    for (let i = 0; i < rows; i++) {
        const cols = rotated_matrix[i].length;
        if (rows !== cols) {
            isDiagonal = false;
        }

        for (let j = 0; j < cols; j++) {
            const val = rotated_matrix[i][j];
            if (val > maxVal) maxVal = val;
            if (val < minVal) minVal = val;
            sumTotal += val;
            totalElements++;

            if (i !== j && val !== 0) {
                isDiagonal = false;
            }
        }
    }

    const average = totalElements > 0 ? sumTotal / totalElements : 0;

    return res.json({
        valor_maximo: maxVal,
        valor_minimo: minVal,
        promedio: average,
        suma_total: sumTotal,
        esDiagonal: isDiagonal
    });
});

const PORT = 3000;
app.listen(PORT, () => {
    console.log(`API en Node.js corriendo en el puerto ${PORT}`);
    console.log(`Documentación disponible en http://localhost:${PORT}/api-docs`);
});