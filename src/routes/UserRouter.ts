import express from 'express';
import { UserController } from '../controller/UserController';

const router = express.Router();
const controller = new UserController;

router.post('/create', controller.new);
router.post('/create/linkedUser', controller.newUserLinked);

router.get('/winners', controller.top10);

export default router;