import { UserRegister, User } from "@/interfaces/user.interface";
import apiClient from "../api/apiClient";
import { AxiosResponse } from "axios";

export const getUserById = async (userId: string): Promise<User> => {
  const response = await apiClient.get(`/users/${userId}`);
  return response.data;
};

export const registerUser = async (
  user: UserRegister,
): Promise<AxiosResponse<User>> => {
  const response = await apiClient.post<User>("/register", user);
  return response;
};
