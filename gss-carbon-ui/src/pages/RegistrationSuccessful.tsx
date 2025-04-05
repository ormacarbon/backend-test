import React, { useEffect, useState } from "react";
import { useParams, Link, useNavigate } from "react-router-dom";
import { Button } from "../components/ui/button";
import { ArrowLeft, Award, Share2 } from "lucide-react";
import ShareSection from "../components/share-section";

interface User {
  id: string;
  name: string;
  points: number;
}

const RegistrationSuccessful: React.FC = () => {
  const { userId } = useParams<{ userId: string }>();
  const navigate = useNavigate();
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function fetchUser() {
      if (!userId) {
        navigate("/");
        return;
      }

      try {
        const data = await getUserById(userId);
        if (!data) {
          navigate("/");
          return;
        }
        setUser(data);
      } catch (error) {
        console.error("Erro ao buscar usuário:", error);
        navigate("/");
      } finally {
        setLoading(false);
      }
    }
    fetchUser();
  }, [userId, navigate]);

  if (loading) {
    return <div>Carregando...</div>;
  }

  if (!user) {
    return null;
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-emerald-50 via-teal-50 to-cyan-50 py-12 px-4 sm:px-6 lg:px-8">
      <div className="mx-auto max-w-md">
        <Link
          to="/"
          className="mb-8 inline-flex items-center text-sm font-medium text-emerald-600 hover:text-emerald-800"
        >
          <ArrowLeft className="mr-2 h-4 w-4" />
          Back to home
        </Link>

        <div className="overflow-hidden rounded-2xl bg-white shadow-xl">
          <div className="relative bg-gradient-to-r from-emerald-500 to-teal-500 px-6 py-12 text-center text-white">
            <div className="absolute -top-6 left-1/2 -translate-x-1/2 transform">
              <div className="flex h-12 w-12 items-center justify-center rounded-full bg-white text-emerald-500 shadow-lg">
                <Award className="h-6 w-6" />
              </div>
            </div>
            <h1 className="text-2xl font-bold">Registration Successful!</h1>
            <p className="mt-2 text-emerald-50">
              Congratulations, {user.name}! You've earned your first point.
            </p>

            <div className="mt-4 inline-flex items-center rounded-full bg-white/20 px-4 py-2 text-sm font-medium backdrop-blur-sm">
              <span>Your current score:</span>
              <span className="ml-2 rounded-full bg-white px-2 py-0.5 text-emerald-600">
                {user.points} points
              </span>
            </div>
          </div>

          <div className="px-6 py-8">
            <div className="mb-6 rounded-lg bg-emerald-50 p-4">
              <div className="flex items-center">
                <Share2 className="mr-3 h-5 w-5 text-emerald-600" />
                <p className="font-medium text-emerald-800">
                  Share your link to earn more points!
                </p>
              </div>
              <p className="mt-2 text-sm text-emerald-700">
                When someone registers using your link, you'll earn an extra
                point.
              </p>
            </div>

            <ShareSection userId={userId as string} />

            <div className="mt-8 text-center">
              <Link to="/leaderboard">
                <Button
                  variant="outline"
                  className="border-emerald-500 text-emerald-700 hover:bg-emerald-50"
                >
                  View Leaderboard
                </Button>
              </Link>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default RegistrationSuccessful;
