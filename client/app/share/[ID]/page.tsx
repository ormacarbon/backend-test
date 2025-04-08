"use client";

import { api } from "@/service/api";
import { User, Response } from "@/types/api";
import { useEffect, useState } from "react";
import { useParams, useRouter } from "next/navigation";

export default function SharePage() {
    const router = useRouter();
    const [shareUrl, setShareUrl] = useState<string>("");
    const [user, setUser] = useState<User | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string>("");
    const params = useParams();
    const { ID } = params;

    useEffect(() => {
        const fetchShareData = async () => {
            try {
                const response: Response<User> = await api.getShareLink(ID as string);
                if (response.user) {
                    const shareUrl = `${window.location.origin}/register?ref=${response.user.share_code}`;
                    setShareUrl(shareUrl);
                    setUser(response.user);
                } else {
                    setError("User not found");
                }
            } catch (error) {
                setError("Failed to fetch share data");
            } finally {
                setLoading(false);
            }
        };

        if (ID) {
            fetchShareData();
        }
    }, [ID]);

    const copyToClipboard = () => {
        navigator.clipboard.writeText(shareUrl);
    };

    if (loading) {
        return <div>Loading...</div>;
    }

    return (
        <div className="min-h-screen bg-gray-50 p-8">
            <div className="max-w-2xl mx-auto bg-white rounded-xl shadow-sm p-8">
                {error ? (
                    <div className="text-center">
                        <div className="mb-4 p-4 bg-red-50 text-red-600 rounded-lg">
                            {error}
                        </div>
                        <button
                            onClick={() => router.push('/')}
                            className="px-6 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-colors"
                        >
                            Go to homepage
                        </button>
                    </div>
                ) : (
                    <>
                        <h1 className="text-3xl font-bold text-gray-800 mb-6">Share and Earn Points!</h1>
                        <p className="text-gray-600 mb-8">
                            Share your unique link with friends and earn points when they register!
                        </p>
                        
                        {user && (
                            <>
                                <div className="flex gap-2 mb-6">
                                    <input
                                        type="text"
                                        readOnly
                                        value={shareUrl}
                                        className="flex-1 px-4 py-2 rounded-lg border-2 border-gray-200"
                                    />
                                    <button
                                        onClick={copyToClipboard}
                                        className="px-6 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-colors"
                                    >
                                        Copy
                                    </button>
                                </div>

                                <div className="bg-blue-50 p-4 rounded-lg mb-6">
                                    <p className="text-blue-800">
                                        Current points: <span className="font-bold">{user.points}</span>
                                    </p>
                                </div>

                                <button
                                    onClick={() => router.push('/')}
                                    className="px-6 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-colors"
                                >
                                    Go to homepage
                                </button>
                                
                            </>
                        )}
                    </>
                )}
            </div>
        </div>
    );
}
