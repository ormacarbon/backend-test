import { Request, Response } from 'express';
import { create, linkedUser, findWinners } from '../repository/User';

export class UserController {
    async new(request: Request, response: Response) {  
        const { body } = request;         
        const { httpCode, ...rest } = await create(body);
        return response.status(httpCode).json(rest);
    }

    async newUserLinked(request: Request, response: Response) {
        const { id } = request.query;
        const { body } = request;         
        const { httpCode, ...rest } = await linkedUser(id as string, body);
        return response.status(httpCode).json(rest);
    }

    async top10(request: Request, response: Response) { 
        const { httpCode, ...rest } = await findWinners();
        return response.status(httpCode).json(rest);
    }
}