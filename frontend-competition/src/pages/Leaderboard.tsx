import { api } from "@/api/api";
import { useEffect, useState } from "react";

interface Leader {
  name: string;
  points: number;
}

export default function Leaderboard() {
  const [leaders, setLeaders] = useState<Leader[]>([]);

  useEffect(() => {
    api.get("/api/winners")
      .then((response) => setLeaders(response.data))
      .catch((error) => console.error("Error fetching leaderboard:", error));
  }, []);

  return (
    <div className="max-w-md mx-auto mt-10 p-6 bg-white shadow-md rounded-lg">
      <h2 className="text-xl font-semibold text-center mb-4">Competition Leaderboard</h2>
      <ul className="space-y-2">
        {leaders.map((leader, index) => (
          <li key={index} className="flex justify-between border-b pb-2">
            <span className="font-medium">{leader.name}</span>
            <span className="text-gray-600">{leader.points} pts</span>
          </li>
        ))}
      </ul>
    </div>
  );
}
