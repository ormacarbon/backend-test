import apiClient from "@/api/apiClient";
import { Referrals } from "@/interfaces/user.interface";
import { AxiosResponse } from "axios";

export const getReferrals = async (): Promise<AxiosResponse<Referrals[]>> => {
    const response = await apiClient.get("/referrals");
    return response;
  }
  