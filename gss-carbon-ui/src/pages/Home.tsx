import { Button } from "@/components/ui/button";
import { ArrowRight, Award, Share2, UserPlus } from "lucide-react";
import React from "react";
import { Link } from "react-router-dom";

const Home: React.FC = () => {
  return (
    <div className="min-h-screen bg-gradient-to-br from-emerald-50 via-teal-50 to-cyan-50">
      <div className="container mx-auto px-4 py-12 sm:px-6 lg:px-8">
        {/* Hero Section */}
        <div className="mx-auto max-w-3xl text-center">
          <div className="mb-6 inline-flex items-center rounded-full bg-emerald-100 px-4 py-1.5 text-sm font-medium text-emerald-800">
            <span className="relative flex h-2 w-2">
              <span className="absolute inline-flex h-full w-full animate-ping rounded-full bg-emerald-400 opacity-75"></span>
              <span className="relative inline-flex h-2 w-2 rounded-full bg-emerald-500"></span>
            </span>
            <span className="ml-2">Competition in progress</span>
          </div>

          <h1 className="mb-6 bg-gradient-to-r from-emerald-600 to-teal-600 bg-clip-text text-4xl font-extrabold tracking-tight text-transparent sm:text-5xl md:text-6xl">
            Carbon Offset Competition
          </h1>

          <p className="mb-10 text-xl text-slate-600">
            Join our mission to spread awareness about carbon offsetting.
            Register, share, and be part of the solution to climate change.
          </p>

          <div className="flex flex-col space-y-4 sm:flex-row sm:space-x-4 sm:space-y-0 sm:justify-center">
            <Link to="/register">
              <Button
                size="lg"
                className="group w-full bg-gradient-to-r from-emerald-500 to-teal-500 text-white hover:from-emerald-600 hover:to-teal-600 sm:w-auto"
              >
                Join Now
                <ArrowRight className="ml-2 h-4 w-4 transition-transform group-hover:translate-x-1" />
              </Button>
            </Link>
            <Link to="/leaderboard">
              <Button
                size="lg"
                variant="outline"
                className="w-full border-emerald-500 text-emerald-700 hover:bg-emerald-50 sm:w-auto"
              >
                View Leaderboard
              </Button>
            </Link>
          </div>
        </div>

        {/* How It Works Section */}
        <div className="mt-16 mx-auto max-w-5xl">
          <h2 className="mb-12 text-center text-3xl font-bold text-slate-800">
            How It Works
          </h2>

          <div className="grid gap-8 md:grid-cols-3">
            {/* Step 1 */}
            <div className="group rounded-xl bg-white p-8 shadow-md transition-all hover:shadow-lg">
              <div className="mb-4 flex h-12 w-12 items-center justify-center rounded-full bg-emerald-100 text-emerald-600 transition-transform group-hover:scale-110">
                <UserPlus className="h-6 w-6" />
              </div>
              <h3 className="mb-3 text-xl font-semibold text-slate-800">
                Register
              </h3>
              <p className="text-slate-600">
                Sign up with your details to join the competition and earn your
                first point instantly.
              </p>
            </div>

            {/* Step 2 */}
            <div className="group rounded-xl bg-white p-8 shadow-md transition-all hover:shadow-lg">
              <div className="mb-4 flex h-12 w-12 items-center justify-center rounded-full bg-teal-100 text-teal-600 transition-transform group-hover:scale-110">
                <Share2 className="h-6 w-6" />
              </div>
              <h3 className="mb-3 text-xl font-semibold text-slate-800">
                Share
              </h3>
              <p className="text-slate-600">
                Spread the word by sharing your unique link. Each registration
                through your link earns you extra points.
              </p>
            </div>

            {/* Step 3 */}
            <div className="group rounded-xl bg-white p-8 shadow-md transition-all hover:shadow-lg">
              <div className="mb-4 flex h-12 w-12 items-center justify-center rounded-full bg-cyan-100 text-cyan-600 transition-transform group-hover:scale-110">
                <Award className="h-6 w-6" />
              </div>
              <h3 className="mb-3 text-xl font-semibold text-slate-800">Win</h3>
              <p className="text-slate-600">
                The top 10 participants with the most points win prizes and
                recognition for their contribution.
              </p>
            </div>
          </div>
        </div>

        {/* Stats Section */}
        <div className="mt-24 mx-auto max-w-5xl rounded-2xl bg-gradient-to-r from-emerald-500 to-teal-500 p-1">
          <div className="rounded-xl bg-white p-8">
            <div className="grid gap-8 text-center md:grid-cols-3">
              <div className="space-y-2">
                <p className="text-4xl font-bold text-emerald-600">1 Point</p>
                <p className="text-slate-600">For each registration</p>
              </div>
              <div className="space-y-2">
                <p className="text-4xl font-bold text-teal-600">+1 Point</p>
                <p className="text-slate-600">For each referral</p>
              </div>
              <div className="space-y-2">
                <p className="text-4xl font-bold text-cyan-600">Top 10</p>
                <p className="text-slate-600">Winners get prizes</p>
              </div>
            </div>
          </div>
        </div>

        {/* CTA Section */}
        <div className="mt-24 mx-auto max-w-3xl rounded-2xl bg-gradient-to-r from-emerald-500 to-teal-500 p-8 text-center text-white">
          <h2 className="mb-4 text-3xl font-bold">Ready to make an impact?</h2>
          <p className="mb-8 text-lg">
            Join thousands of others in spreading awareness about carbon
            offsetting.
          </p>
          <Link to="/register">
            <Button
              size="lg"
              className="bg-white font-semibold text-emerald-600 hover:bg-emerald-50"
            >
              Join the Competition
            </Button>
          </Link>
        </div>
      </div>
    </div>
  );
};

export default Home;
