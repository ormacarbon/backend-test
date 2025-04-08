import { User } from "@/interfaces/user.interface";
import apiClient from "../api/apiClient";

export const getLeaderboard = async (): Promise<User[]> => {
  const response = await apiClient.get("/leaderboard");
  return response.data;
};
