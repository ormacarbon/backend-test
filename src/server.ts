import 'dotenv/config'
import "reflect-metadata"
import express from 'express';
import dbConfig from './config/database';
import routes from './routes';
import { createConnection } from 'typeorm';

const app = express();
const PORT = process.env.PORT || 8080;
app.use(express.json());

createConnection(dbConfig)
    .then(() => {
        app.listen(PORT, () => console.log('Server is running on port', PORT));
    })
    .catch((err) => {
        console.error(`Error during Database initialization: ${err}`);
        process.exit(1);
    })

app.use(routes);






