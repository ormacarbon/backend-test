import { ConnectionOptions } from "typeorm";

const dbConfig: ConnectionOptions = {
    type: "mysql",
    host: process.env.MYSQL_HOST || "localhost",
    port: Number(process.env.MYSQL_PORT) || 3306,
    username: process.env.MYSQL_USER || "root",
    password: process.env.MYSQL_PASSWORD || "test",
    database: process.env.MYSQL_DB || "test",
    entities: [process.env.NODE_ENV === 'local' ? './src/entity/*.ts' : './dist/entity/*.js'],
    synchronize: false,
    logging: true,
};

export default dbConfig;