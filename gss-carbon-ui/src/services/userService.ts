import { UserRegister, User } from "@/interfaces/user.interface";
import apiClient from "../api/apiClient";
import { AxiosResponse } from "axios";

export const getUserById = async (
  userId: string,
): Promise<AxiosResponse<User>> => {
  const response = await apiClient.get(`/user/${userId}`);
  return response;
};

export const registerUser = async (
  user: UserRegister,
): Promise<AxiosResponse<User>> => {
  const response = await apiClient.post<User>("/user/register", user);
  return response;
};

export const getUserByReferralToken = async (
  token: string,
): Promise<AxiosResponse<User>> => {
  const response = await apiClient.get(`/user/referral/${token}`);
  return response;
};
