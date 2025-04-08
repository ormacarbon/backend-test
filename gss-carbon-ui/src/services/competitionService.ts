import { User } from "@/interfaces/user.interface";
import apiClient from "../api/apiClient";
import { AxiosResponse } from "axios";

export const finishCompetition = async (): Promise<AxiosResponse<User[]>> => {
  const response = await apiClient.post("/competition/finish");
  return response;
};
