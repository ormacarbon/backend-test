import { User } from "@/interfaces/user.interface";
import queryClient from "@/query/queryClient";
import { queryKeys } from "@/query/queryKeys";
import { finishCompetition } from "@/services/competitionService";
import { useMutation } from "@tanstack/react-query";
import { AxiosResponse, AxiosError } from "axios";

export const useFinishCompetition = () => {
  return useMutation<AxiosResponse<User[]>, AxiosError>({
    mutationFn: () => finishCompetition(),
    onSuccess: () => {
      queryClient.invalidateQueries({queryKey: queryKeys.leaderboard()})
      queryClient.invalidateQueries({queryKey: queryKeys.referrals()});
      },
    onError: () => { },
  });
};