import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8080/api',
});

export const login = async (email: string, password: string) => {
  const response = await api.post('/login', { email, password });
  return response.data;
};

export const createUser = async (userData: {
  name: string;
  email: string;
  password: string;
  phone: string;
  invite_code?: string;
}) => {
  const response = await api.post('/users', userData);
  return response.data;
};

export const getMe = async (token: string) => {
  const response = await api.get('/me', {
    headers: {
      Authorization: token,
    },
  });
  return response.data;
};

export const getRanking = async () => {
  const response = await api.get('/users/ranking');
  return response.data;
};