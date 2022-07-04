import {
    Column, Entity, JoinColumn, OneToOne, PrimaryGeneratedColumn,
} from 'typeorm';
import { UserEntity } from './user';

@Entity({ name: 'point' })
export class PointEntity {
    @PrimaryGeneratedColumn()
    id!: number;

    @OneToOne(() => UserEntity, (user) => user.id)
    @JoinColumn({ name: 'user_id'})
    user_id!: UserEntity | number;

    @Column()
    points!: number;    
}