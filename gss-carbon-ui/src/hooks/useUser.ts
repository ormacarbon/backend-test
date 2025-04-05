import { User, UserRegister } from "@/interfaces/user.interface";
import { queryKeys } from "@/query/queryKeys";
import { registerUser } from "@/services/userService";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { AxiosResponse, AxiosError } from "axios";
import { toast } from "react-toastify";

export const useRegisterUser = () => {
  const queryClient = useQueryClient();

  return useMutation<AxiosResponse<User>, AxiosError, UserRegister>({
    mutationFn: (user: UserRegister) => registerUser(user),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.leaderboard() });
    },
    onError: (error) => {
      if (error.response && error.response.status === 409) {
        toast.info("This e-mail is already in use 🥲", {
          progressClassName: "bg-gradient-to-r! from-emerald-500! to-teal-500!",
        });
      }
    },
  });
};
