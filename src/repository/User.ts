import  Response  from '../config/utils/response';
import { getManager, getRepository } from 'typeorm';
import { ResponsePattern } from '../types/response';
import { UserEntity } from './../entity/user';
import { PointEntity } from './../entity/point';

const response = new Response();

export const create = async(body: UserEntity): Promise<ResponsePattern> => {
    if (!body) return  response.badRequest<null, string>(400, null, 'Requisição sem conteúdo!');
    
    const { name, email, phone } = body;

    if (!name || !email || !phone) {
        return response.badRequest<null, string>(400, null, 'Preencha todos os campos!');
    }

    const phoneNumber = Number(phone);

    if (Number.isNaN(phoneNumber) || !phoneNumber) return response.badRequest<null, string>(400, null, 'Formato do número inválido!');

    const userExists = await getRepository(UserEntity).findOne({
        where: {
            email,
        }
    });
    if (userExists) return response.badRequest<null, string>(400, null, 'Usuário já cadastrado!');

    try {
        const transaction = await getManager().transaction(async (transactionalEntityManager) => {
            const data = {
                name,
                email,
                phone,
            };
            const user = await transactionalEntityManager.getRepository(UserEntity).save(data);

            const point = {
                user_id: user.id,
                points: 1,
            }
            await transactionalEntityManager.getRepository(PointEntity).save(point);

            return 'Novo usuário cadastrado e 1 ponto atribuído!';
        });
        return response.successfulRequest<string, null>(201, transaction, null);
    } catch (error) {
        return response.error<string>(`Error: ${error}`);
    }
}

export const linkedUser = async(userId: string, body: UserEntity): Promise<ResponsePattern> => {
    if (!body) return  response.badRequest<null, string>(400, null, 'Requisição sem conteúdo!');
    
    const { name, email, phone } = body;

    if (!name || !email || !phone) {
        return response.badRequest<null, string>(400, null, 'Preencha todos os campos!');
    }

    const phoneNumber = Number(phone);

    if (Number.isNaN(phoneNumber) || !phoneNumber) return response.badRequest<null, string>(400, null, 'Formato do número inválido!');

    const userExists = await getRepository(UserEntity).findOne({
        where: {
            email,
        }
    });

    if (userExists) return response.badRequest<null, string>(400, null, 'Usuário já cadastrado!');

    const userIdNumber = Number(userId);
    if (Number.isNaN(userIdNumber) || !userIdNumber) return response.badRequest<null, string>(400, null, 'Formato do ID inválido!');

    const linkedUser = await getRepository(UserEntity).findOne({ id: userIdNumber});

    try {
        const transaction = linkedUser 
            ?  await getManager().transaction(async (transactionalEntityManager) => {
                const data = {
                    name,
                    email,
                    phone,
                };
                const user = await transactionalEntityManager.getRepository(UserEntity).save(data);

                const point = {
                    user_id: user.id,
                    points: 1,
                }
                await transactionalEntityManager.getRepository(PointEntity).save(point);

                const searchPointsLinkedUser = await getRepository(PointEntity).findOne({ user_id: linkedUser.id }) as PointEntity;
                const points = searchPointsLinkedUser?.points ?? 0;

                await transactionalEntityManager.getRepository(PointEntity).update(
                    { 
                        id: linkedUser.id 
                    },
                    { 
                        user_id: linkedUser.id,
                        points: points ? points + 1 : 1, 
                     }
                );

                return `Novo usuário cadastrado e 1 ponto atribuído a ele e ao usuário ${linkedUser!.id}!`;
            })
            : await getManager().transaction(async (transactionalEntityManager) => {
                const data = {
                    name,
                    email,
                    phone,
                };
                const user = await transactionalEntityManager.getRepository(UserEntity).save(data);

                const point = {
                    user_id: user.id,
                    points: 1,
                }
                await transactionalEntityManager.getRepository(PointEntity).save(point)

                return `Novo usuário cadastrado e 1 ponto atribuído!`;
            });
        return response.successfulRequest<string, null>(201, transaction, null);
    } catch (error) {
        return response.error<string>(`Error: ${error}`);
    }
}

export const findWinners = async(): Promise<ResponsePattern> => {
    
    const data = await getRepository(PointEntity).find({
        order: { 
            points: 'DESC',
        },
        take: 10,
        relations: ['user_id'],
    })

    return response.successfulRequest<PointEntity | PointEntity[], null>(200, data, null);
}