import React from "react";
import RegistrationForm from "../components/registration-form";
import { Link, useSearchParams } from "react-router-dom";
import { ArrowLeft } from "lucide-react";

const Register: React.FC = () => {
  const [searchParams] = useSearchParams();
  const referralCode = searchParams.get("ref") || null;

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
              Register to earn your first point and get your unique sharing link
            </p>
          </div>

          <div className="px-6 py-8">
            {referralCode && (
              <div className="mb-6 rounded-lg bg-emerald-50 p-4 text-sm text-emerald-800">
                You were invited by a friend! They'll earn a point when you
                register.
              </div>
            )}

            <RegistrationForm referralCode={referralCode} />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Register;
