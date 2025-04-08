import { queryKeys } from "@/query/queryKeys";
import { getLeaderboard } from "@/services/leaderboardService";
import { useQuery } from "@tanstack/react-query";

export const useLeaderboard = () => {
  return useQuery({
    queryKey: [queryKeys.leaderboard()],
    queryFn: () => getLeaderboard(),
  });
};
