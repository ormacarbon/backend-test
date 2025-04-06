import React from "react";
import RegistrationForm from "../components/registration-form";
import { Link, useSearchParams } from "react-router-dom";
import { ArrowLeft, UserCircle } from "lucide-react";
import { useGetUserByReferralToken } from "@/hooks/useUser";

const Register: React.FC = () => {
  const [searchParams] = useSearchParams();
  const referralToken = searchParams.get("ref") || null;
  const { data: userData, isLoading } = useGetUserByReferralToken(
    referralToken!,
  );

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
          <div className="bg-gradient-to-r from-emerald-500 to-teal-500 px-6 py-8 text-white">
            <h1 className="text-2xl font-bold">Join the Competition</h1>
            <p className="mt-2 text-emerald-50">
              Register to earn your first point and grab your unique sharing
              link.
            </p>
          </div>

          <div className="px-6 py-8 space-y-6">
            {referralToken && (
              <div className="flex items-center gap-4 rounded-lg bg-emerald-100 p-4 text-sm text-emerald-800 shadow-md">
                <UserCircle className="h-6 w-6 text-emerald-600" />
                <p>
                  You were invited by{" "}
                  {isLoading ? (
                    <span className="animate-pulse">...</span>
                  ) : (
                    userData?.data.name.split(" ")[0] || "a friend"
                  )}
                  . Your awesome referrer gets a point when you sign up! 😄
                </p>
              </div>
            )}

            <RegistrationForm referralCode={referralToken} />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Register;
