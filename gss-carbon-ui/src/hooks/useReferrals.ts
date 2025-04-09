import { Referrals } from "@/interfaces/user.interface";
import { queryKeys } from "@/query/queryKeys";
import { getReferrals } from "@/services/referralsService";
import { useQuery } from "@tanstack/react-query";
import { AxiosResponse, AxiosError } from "axios";

export const useReferrals = () => {
    return useQuery<AxiosResponse<Referrals[]>, AxiosError>({
      queryKey: queryKeys.referrals(),
      queryFn: () => getReferrals(),
    });
  }
  