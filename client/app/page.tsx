"use client";

import { api } from "@/service/api";
import { User } from "@/types/api";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import Header from "@/components/shared/Header";
import Footer from "@/components/shared/Footer";

export default function Home() {
	const [leaderboard, setLeaderboard] = useState<User[]>([]);
	const [sort, setSort] = useState<string>("points");
	const [loading, setLoading] = useState<boolean>(true);
	const [error, setError] = useState<string | null>(null);
	const [search, setSearch] = useState<string>("");
	const [page, setPage] = useState<number>(1);
	const [totalPages, setTotalPages] = useState<number>(0);
	const router = useRouter();

	const fetchLeaderboard = async () => {
		try {
			setLoading(true);
			const res = await api.getLeaderboard({ 
				sort, 
				page: page.toString(), 
				search 
			});
			setLeaderboard(res.leaderboard ?? []);
			setTotalPages(res.totalPages);
			setLoading(false);
		} catch (error) {
			setError("Failed to load leaderboard");
			setLoading(false);
		}
	};

	useEffect(() => {
		fetchLeaderboard();
	}, [sort, page, search]);

	return (
		<div className="min-h-screen bg-gray-50 flex flex-col">
			<Header />
			<div className="flex-1 p-8">
				<div className="max-w-6xl mx-auto">
					<div className="mb-8">
						<h1 className="text-4xl font-bold text-gray-800">Leaderboard</h1>
						<p className="text-gray-600 mt-2">Track top performers and rankings</p>
					</div>

					<div className="flex flex-col sm:flex-row justify-between gap-4 mb-6">
						<div className="relative">
							<input
								type="text"
								placeholder="Search by name"
								value={search}
								onChange={(e) => setSearch(e.target.value)}
								className="w-full sm:w-80 pl-4 pr-10 py-2 rounded-lg border-2 border-gray-200 focus:border-blue-500 focus:ring-2 focus:ring-blue-200 transition-all outline-none"
							/>
							<svg className="w-5 h-5 absolute right-3 top-3 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
								<path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
							</svg>
						</div>
						<select 
							value={sort} 
							onChange={e => setSort(e.target.value)}
							className="px-4 py-2 rounded-lg border-2 border-gray-200 focus:border-blue-500 outline-none cursor-pointer"
						>
							<option value="points">Order by points</option>
							<option value="name">Order by name</option>
							<option value="email">Order by e-mail</option>
						</select>
					</div>

					<div className="bg-white rounded-xl shadow-sm overflow-hidden">
						{loading ? (
							<div className="p-8 text-center">
								<div className="animate-spin w-8 h-8 border-4 border-blue-500 border-t-transparent rounded-full mx-auto"></div>
								<p className="text-gray-600 mt-4">Loading data...</p>
							</div>
						) : error ? (
							<div className="p-8 text-center text-red-500">
								<svg className="w-12 h-12 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
									<path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
								</svg>
								<p>{error}</p>
							</div>
						) : leaderboard.length === 0 ? (
							<div className="p-8 text-center text-gray-500">
								<svg className="w-12 h-12 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
									<path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4" />
								</svg>
								<p>No users found</p>
							</div>
						) : (
							<table className="w-full">
								<thead className="bg-gray-50">
									<tr>
										<th className="px-6 py-4 text-left text-sm font-semibold text-gray-600">Points</th>
										<th className="px-6 py-4 text-left text-sm font-semibold text-gray-600">Name</th>
										<th className="px-6 py-4 text-left text-sm font-semibold text-gray-600">Email</th>
									</tr>
								</thead>
								<tbody className="divide-y divide-gray-100">
									{leaderboard.map((user) => (
										<tr key={user.id} className="hover:bg-gray-50 transition-colors">
											<td className="px-6 py-4 whitespace-nowrap font-semibold text-blue-600">{user.points}</td>
											<td className="px-6 py-4 whitespace-nowrap">{user.name}</td>
											<td className="px-6 py-4 whitespace-nowrap text-gray-500">{user.email}</td>
										</tr>
									))}
								</tbody>
							</table>
						)}
					</div>

					<div className="mt-6 flex items-center justify-between">
						<button 
							disabled={page <= 1} 
							onClick={() => setPage(page - 1)} 
							className="px-4 py-2 rounded-lg bg-white border-2 border-gray-200 text-gray-600 disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-50 transition-colors"
						>
							Previous
						</button>
						<span className="text-gray-600">Page {page} of {totalPages}</span>
						<button 
							disabled={page >= totalPages} 
							onClick={() => setPage(page + 1)}
							className="px-4 py-2 rounded-lg bg-blue-500 text-white disabled:opacity-50 disabled:cursor-not-allowed hover:bg-blue-600 transition-colors"
						>
							Next
						</button>
					</div>
				</div>
			</div>
			<Footer />
		</div>
	);
}
