import React from "react";
import { Link } from "react-router-dom";
import { Button } from "../components/ui/button";
import { ArrowLeft, Trophy, Medal, Award } from "lucide-react";
import { useLeaderboard } from "@/hooks/useLeaderboard";

const LeaderboardPage: React.FC = () => {
  const { data: leaderboardData, isLoading } = useLeaderboard();

  return (
    <div className="min-h-screen bg-gradient-to-br from-emerald-50 via-teal-50 to-cyan-50 py-12 px-4 sm:px-6 lg:px-8">
      <div className="mx-auto max-w-4xl">
        <Link
          to="/"
          className="mb-8 inline-flex items-center text-sm font-medium text-emerald-600 hover:text-emerald-800"
        >
          <ArrowLeft className="mr-2 h-4 w-4" />
          Back to home
        </Link>

        <div className="mt-8 rounded-2xl bg-white shadow-xl">
          <div className="relative bg-gradient-to-r from-emerald-500 to-teal-500 px-6 py-12 text-center text-white">
            <div className="absolute -top-6 left-1/2 -translate-x-1/2 transform">
              <div className="flex h-12 w-12 items-center justify-center rounded-full bg-white text-emerald-500 shadow-lg">
                <Trophy className="h-6 w-6" />
              </div>
            </div>
            <h1 className="text-2xl font-bold">Competition Leaderboard</h1>
            <p className="mt-2 text-emerald-50">
              Top participants spreading awareness about carbon offsetting
            </p>
          </div>

          <div className="px-6 py-8">
            {isLoading ? (
              <div className="text-center">
                <p className="text-lg text-slate-600">Loading leaderboard...</p>
              </div>
            ) : leaderboardData && leaderboardData.length > 0 ? (
              <div className="overflow-hidden rounded-lg border border-slate-200">
                <table className="w-full">
                  <thead>
                    <tr className="bg-slate-50">
                      <th className="px-6 py-3 text-left text-sm font-semibold text-slate-700">Rank</th>
                      <th className="px-6 py-3 text-left text-sm font-semibold text-slate-700">Name</th>
                      <th className="px-6 py-3 text-right text-sm font-semibold text-slate-700">Points</th>
                    </tr>
                  </thead>
                  <tbody className="divide-y divide-slate-200">
                    {leaderboardData.map((user, index) => (
                      <tr
                        key={user.id}
                        className={index < 10 ? "bg-emerald-50/30" : ""}
                      >
                        <td className="whitespace-nowrap px-6 py-4 text-sm">
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
                            <span className={index < 3 ? "font-medium" : ""}>
                              {index < 3 ? (
                                <span className="sr-only">
                                  {index === 0
                                    ? "First place"
                                    : index === 1
                                    ? "Second place"
                                    : "Third place"}
                                </span>
                              ) : (
                                <span className="sr-only">Rank {index + 1}</span>
                              )}
                            </span>
                          </div>
                        </td>
                        <td className="whitespace-nowrap px-6 py-4 text-sm font-medium text-slate-800">
                          {user.name}
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
                <p className="mt-2 text-slate-500">
                  Be the first to join the competition!
                </p>
                <div className="mt-6">
                  <Link to="/register">
                    <Button className="bg-gradient-to-r from-emerald-500 to-teal-500 text-white hover:from-emerald-600 hover:to-teal-600">
                      Join Now
                    </Button>
                  </Link>
                </div>
              </div>
            )}

            <div className="mt-8 text-center">
              <Link to="/register">
                <Button className="bg-gradient-to-r from-emerald-500 to-teal-500 text-white hover:from-emerald-600 hover:to-teal-600">
                  Join the Competition
                </Button>
              </Link>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default LeaderboardPage;
