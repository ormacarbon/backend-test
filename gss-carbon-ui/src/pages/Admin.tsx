import React, { useEffect, useState } from "react";
import { Button } from "../components/ui/button";

interface User {
  id: string;
  name: string;
  email: string;
  phone: string;
  points: number;
}

interface Referral {
  id: string;
  createdAt: string | Date;
  referrerId: string;
  referredId: string;
}

const AdminPage: React.FC = () => {
  const [leaderboard, setLeaderboard] = useState<User[]>([]);
  const [referrals, setReferrals] = useState<Referral[]>([]);
  const [loading, setLoading] = useState(true);

  const endCompetition = async (e: React.FormEvent) => {};

  if (loading) {
    return <div>Carregando...</div>;
  }

  return (
    <div className="container mx-auto p-6">
      <h1 className="mb-6 text-3xl font-bold">Admin Dashboard</h1>

      <div className="mb-8 rounded-lg border bg-white p-6 shadow-sm">
        <h2 className="mb-4 text-xl font-semibold">Competition Management</h2>
        <form onSubmit={endCompetition}>
          <Button type="submit" className="bg-red-600 hover:bg-red-700">
            End Competition & Notify Winners
          </Button>
        </form>
      </div>

      <div className="mb-8 rounded-lg border bg-white p-6 shadow-sm">
        <h2 className="mb-4 text-xl font-semibold">Leaderboard</h2>
        <div className="overflow-x-auto">
          <table className="w-full">
            <thead>
              <tr className="border-b">
                <th className="px-4 py-2 text-left">Rank</th>
                <th className="px-4 py-2 text-left">Name</th>
                <th className="px-4 py-2 text-left">Email</th>
                <th className="px-4 py-2 text-left">Phone</th>
                <th className="px-4 py-2 text-right">Points</th>
              </tr>
            </thead>
            <tbody>
              {leaderboard.map((user, index) => (
                <tr key={user.id} className="border-b">
                  <td className="px-4 py-2">{index + 1}</td>
                  <td className="px-4 py-2">{user.name}</td>
                  <td className="px-4 py-2">{user.email}</td>
                  <td className="px-4 py-2">{user.phone}</td>
                  <td className="px-4 py-2 text-right">{user.points}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>

      <div className="rounded-lg border bg-white p-6 shadow-sm">
        <h2 className="mb-4 text-xl font-semibold">Referral History</h2>
        <div className="overflow-x-auto">
          <table className="w-full">
            <thead>
              <tr className="border-b">
                <th className="px-4 py-2 text-left">Date</th>
                <th className="px-4 py-2 text-left">Referrer ID</th>
                <th className="px-4 py-2 text-left">Referred ID</th>
              </tr>
            </thead>
            <tbody>
              {referrals.map((referral) => (
                <tr key={referral.id} className="border-b">
                  <td className="px-4 py-2">
                    {new Date(referral.createdAt).toLocaleString()}
                  </td>
                  <td className="px-4 py-2">{referral.referrerId}</td>
                  <td className="px-4 py-2">{referral.referredId}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
};

export default AdminPage;
