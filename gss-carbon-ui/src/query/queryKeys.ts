export const queryKeys = {
  leaderboard: () => ["leaderboard"] as const,
  userById: (id: string) => ["user", id] as const,
  userByReferralToken: (token: string) => ["userByReferralToken", token] as const,
  referrals: () => ["referrals"] as const,
};
