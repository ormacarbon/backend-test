import React from "react";
import { Link } from "react-router-dom";
import { Button } from "../components/ui/button";
import { Trophy, Medal, Award } from "lucide-react";
import { useFinishCompetition } from "@/hooks/useCompetition";
import { useLeaderboard } from "@/hooks/useLeaderboard";
import { useReferrals } from "@/hooks/useReferrals";

const Admin: React.FC = () => {
  const { mutateAsync, isPending: finishing } = useFinishCompetition();
  const { data: leaderboard, isLoading: leaderboardLoading } = useLeaderboard();
  const { data: referrals, isLoading: referralsLoading } = useReferrals();

  const endCompetition = async (e: React.FormEvent) => {
    e.preventDefault();
    await mutateAsync();
  };

  return (
    <div className="container mx-auto p-6">
      <h1 className="mb-6 text-3xl font-bold text-emerald-700">Admin Dashboard</h1>

      {/* Competition Management Section */}
      <div className="mb-8 rounded-lg border bg-white p-6 shadow-sm">
        <h2 className="mb-4 text-xl font-semibold text-emerald-700">Competition Management</h2>
        <form onSubmit={endCompetition}>
          <Button
            type="submit"
            className="bg-red-600 hover:bg-red-700 text-white"
            disabled={finishing}
          >
            {finishing ? "Ending Competition..." : "End Competition & Notify Winners"}
          </Button>
        </form>
      </div>

      {/* Leaderboard Section */}
      <div className="mb-8 rounded-lg border bg-white p-6 shadow-sm">
        <h2 className="mb-4 text-xl font-semibold text-emerald-700">Leaderboard</h2>
        {leaderboardLoading ? (
          <p className="text-center text-slate-600">Loading leaderboard...</p>
        ) : leaderboard && leaderboard.length > 0 ? (
          <div className="overflow-x-auto">
            <table className="w-full">
              <thead>
                <tr className="bg-slate-50">
                  <th className="px-6 py-3 text-left text-sm font-semibold text-slate-700">
                    Rank
                  </th>
                  <th className="px-6 py-3 text-left text-sm font-semibold text-slate-700">
                    Name
                  </th>
                  <th className="px-6 py-3 text-left text-sm font-semibold text-slate-700">
                    Email
                  </th>
                  <th className="px-6 py-3 text-left text-sm font-semibold text-slate-700">
                    Phone
                  </th>
                  <th className="px-6 py-3 text-right text-sm font-semibold text-slate-700">
                    Points
                  </th>
                </tr>
              </thead>
              <tbody className="divide-y divide-slate-200">
                {leaderboard.map((user, index) => (
                  <tr key={user.id} className={index < 10 ? "bg-emerald-50/30" : ""}>
                    <td className="px-6 py-4 text-sm">
                      <div className="flex items-center">
                        {index === 0 ? (
                          <div className="mr-2 flex h-8 w-8 items-center justify-center rounded-full bg-yellow-100 text-yellow-600">
                            <Trophy className="h-4 w-4" />
                          </div>
                        ) : index === 1 ? (
                          <div className="mr-2 flex h-8 w-8 items-center justify-center rounded-full bg-slate-100 text-slate-600">
                            <Medal className="h-4 w-4" />
                          </div>
                        ) : index === 2 ? (
                          <div className="mr-2 flex h-8 w-8 items-center justify-center rounded-full bg-amber-100 text-amber-600">
                            <Award className="h-4 w-4" />
                          </div>
                        ) : (
                          <div className="mr-2 flex h-8 w-8 items-center justify-center rounded-full bg-slate-50 text-slate-500">
                            {index + 1}
                          </div>
                        )}
                      </div>
                    </td>
                    <td className="whitespace-nowrap px-6 py-4 text-sm font-medium text-slate-800">
                      {user.name}
                    </td>
                    <td className="whitespace-nowrap px-6 py-4 text-sm text-slate-800">
                      {user.email}
                    </td>
                    <td className="whitespace-nowrap px-6 py-4 text-sm text-slate-800">
                      {user.phone}
                    </td>
                    <td className="whitespace-nowrap px-6 py-4 text-right text-sm">
                      <span className="font-medium text-emerald-600">
                        {user.points}
                      </span>
                      <span className="ml-1 text-slate-500">points</span>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        ) : (
          <div className="rounded-lg border border-dashed border-slate-300 bg-slate-50 p-8 text-center">
            <Trophy className="mx-auto h-12 w-12 text-slate-300" />
            <h3 className="mt-4 text-lg font-medium text-slate-700">
              No participants yet
            </h3>
            <p className="mt-2 text-slate-500">Be the first to join the competition!</p>
            <div className="mt-6">
              <Link to="/register">
                <Button className="bg-gradient-to-r from-emerald-500 to-teal-500 text-white hover:from-emerald-600 hover:to-teal-600">
                  Join Now
                </Button>
              </Link>
            </div>
          </div>
        )}
      </div>

      {/* Referral History Section */}
      <div className="rounded-lg border bg-white p-6 shadow-sm">
        <h2 className="mb-4 text-xl font-semibold text-emerald-700">Referral History</h2>
        {referralsLoading ? (
          <p className="text-center text-slate-600">Loading referral history...</p>
        ) : referrals && referrals.data && referrals.data.length > 0 ? (
          <div className="overflow-x-auto">
            <table className="w-full">
              <thead>
                <tr className="border-b">
                  <th className="px-4 py-2 text-left text-sm font-semibold text-slate-700">
                    Date
                  </th>
                  <th className="px-4 py-2 text-left text-sm font-semibold text-slate-700">
                    Referrer Name
                  </th>
                  <th className="px-4 py-2 text-left text-sm font-semibold text-slate-700">
                    Referred Name
                  </th>
                </tr>
              </thead>
              <tbody className="divide-y divide-slate-200">
                {referrals.data.map((referral) => (
                  <tr key={referral.id} className="border-b">
                    <td className="px-4 py-2">
                      {new Date(referral.createdAt).toLocaleString()}
                    </td>
                    <td className="px-4 py-2">{referral.referrer.name}</td>
                    <td className="px-4 py-2">{referral.referred.name}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        ) : (
          <div className="rounded-lg border border-dashed border-slate-300 bg-slate-50 p-8 text-center">
            <Trophy className="mx-auto h-12 w-12 text-slate-300" />
            <h3 className="mt-4 text-lg font-medium text-slate-700">
              No referral history available.
            </h3>
          </div>
        )}
      </div>
    </div>
  );
};

export default Admin;
