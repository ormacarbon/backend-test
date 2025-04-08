"use client";

import { api } from "@/service/api";
import { useState } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import Header from "@/components/shared/Header";
import Footer from "@/components/shared/Footer";

export default function RegisterPage() {
    const router = useRouter();
    const searchParams = useSearchParams();
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState("");

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setLoading(true);
        setError("");

        const formData = new FormData(e.currentTarget);
        const data = {
            name: formData.get("name") as string,
            email: formData.get("email") as string,
            phone_number: formData.get("phone") as string,
            referral_code: searchParams.get("ref") as string 
                        || formData.get("referral_code") as string,
        };

        try {
            const response = await api.register(data);
            debugger
            router.push(`/share/${response.user?.ID}`);
        } catch (err) {
            setError("Failed to register. Please try again.");
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="min-h-screen bg-gray-50 flex flex-col">
            <Header />
            <div className="flex-1 p-8">
                <div className="max-w-md mx-auto bg-white rounded-xl shadow-sm p-8">
                    <h1 className="text-3xl font-bold text-gray-800 mb-6">Register</h1>
                    
                    {error && (
                        <div className="mb-4 p-4 bg-red-50 text-red-600 rounded-lg">
                            {error}
                        </div>
                    )}

                    <form onSubmit={handleSubmit}>
                        <div className="space-y-4">
                            <div>
                                <label className="block text-sm font-medium text-gray-700 mb-1">Name</label>
                                <input
                                    required
                                    type="text"
                                    name="name"
                                    className="w-full px-4 py-2 rounded-lg border-2 border-gray-200"
                                />
                            </div>
                            <div>
                                <label className="block text-sm font-medium text-gray-700 mb-1">Email</label>
                                <input
                                    required
                                    type="email"
                                    name="email"
                                    className="w-full px-4 py-2 rounded-lg border-2 border-gray-200"
                                />
                            </div>
                            <div>
                                <label className="block text-sm font-medium text-gray-700 mb-1">Phone</label>
                                <input
                                    required
                                    type="tel"
                                    name="phone"
                                    className="w-full px-4 py-2 rounded-lg border-2 border-gray-200"
                                />
                            </div>
                            <div>
                                <label className="block text-sm font-medium text-gray-700 mb-1">Referral code (optional)</label>
                                <input
                                    type="text"
                                    name="referral_code"
                                    defaultValue={searchParams.get("ref") || ""}
                                    className="w-full px-4 py-2 rounded-lg border-2 border-gray-200"
                                />
                            </div>
                            <button
                                type="submit"
                                disabled={loading}
                                className="w-full py-2 px-4 bg-blue-500 text-white rounded-lg hover:bg-blue-600 disabled:opacity-50"
                            >
                                {loading ? "Registering..." : "Register"}
                            </button>
                        </div>
                    </form>
                </div>
            </div>
            <Footer />
        </div>
    );
}
