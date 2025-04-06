'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { Leaf, Copy, Crown } from 'lucide-react';
import { getMe, getRanking } from '@/lib/services/api';

interface User {
  id: string;
  name: string;
  email: string;
  phone: string;
  points: number;
  link_code: string;
}

interface RankingUser {
  user_id: string;
  name: string;
  points: number;
}

export default function Dashboard() {
  const router = useRouter();
  const [user, setUser] = useState<User | null>(null);
  const [ranking, setRanking] = useState<RankingUser[]>([]);
  const [copied, setCopied] = useState(false);

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (!token) {
      router.push('/');
      return;
    }

    const fetchData = async () => {
      try {
        const [userData, rankingData] = await Promise.all([
          getMe(token),
          getRanking(),
        ]);
        setUser(userData);
        setRanking(rankingData);
      } catch (err) {
        localStorage.removeItem('token');
        router.push('/');
      }
    };

    fetchData();
  }, [router]);

  const handleCopyInviteCode = () => {
    if (user) {
      navigator.clipboard.writeText(user.link_code);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    }
  };

  if (!user) return null;

  return (
    <div className="min-h-screen bg-[var(--background)]">
      <header className="bg-white shadow">
        <div className="max-w-7xl mx-auto px-4 py-4 sm:px-6 lg:px-8 flex justify-between items-center">
          <div className="flex items-center">
            <Leaf className="w-8 h-8 text-[var(--accent)]" />
            <h1 className="text-2xl font-bold ml-2">vbio</h1>
          </div>
          <button
            onClick={() => {
              localStorage.removeItem('token');
              router.push('/');
            }}
            className="text-[var(--accent)] hover:text-[var(--accent-hover)]"
          >
            Sair
          </button>
        </div>
      </header>

      <main className="max-w-7xl mx-auto px-4 py-8 sm:px-6 lg:px-8">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
          <div className="bg-white p-6 rounded-lg shadow">
            <h2 className="text-xl font-semibold mb-4">Seu perfil</h2>
            <div className="space-y-2">
              <p><span className="font-medium">Nome:</span> {user.name}</p>
              <p><span className="font-medium">Email:</span> {user.email}</p>
              <p><span className="font-medium">Telefone:</span> {user.phone}</p>
              <p><span className="font-medium">Pontos:</span> {user.points}</p>
              <div className="mt-4">
                <p className="font-medium mb-2">Seu código de convite:</p>
                <button
                  onClick={handleCopyInviteCode}
                  className="flex items-center gap-2 px-4 py-2 bg-[var(--muted)] rounded-md hover:bg-gray-200 transition-colors"
                >
                  <span className="truncate">{user.id}</span>
                  <Copy className="w-4 h-4" />
                </button>
                {copied && (
                  <p className="text-sm text-green-600 mt-1">Código copiado!</p>
                )}
              </div>
            </div>
          </div>

          <div className="bg-white p-6 rounded-lg shadow">
            <h2 className="text-xl font-semibold mb-4 flex items-center gap-2">
              <Crown className="w-6 h-6 text-[var(--accent)]" />
              Ranking
            </h2>
            <div className="space-y-4">
              {ranking.map((rank, index) => (
                <div
                  key={rank.user_id}
                  className={`flex items-center justify-between p-3 rounded-md ${
                    rank.user_id === user.id
                      ? 'bg-[var(--accent)] text-white'
                      : 'bg-[var(--muted)]'
                  }`}
                >
                  <div className="flex items-center gap-3">
                    <span className="font-bold">{index + 1}º</span>
                    <span>{rank.name}</span>
                  </div>
                  <span className="font-medium">{rank.points} pontos</span>
                </div>
              ))}
            </div>
          </div>
        </div>
      </main>
    </div>
  );
}