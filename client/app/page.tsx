"use client";

import { api } from "@/service/api";
import { User } from "@/types/api";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";

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
			const res = await api.getLeaderboard({ sort, page: page.toString(), search });
			setLeaderboard(res.data ?? []);
			setLoading(false);
		} catch (error) {
			setError("Failed to fetch leaderboard");
			setLoading(false);
		}
	};

	useEffect(() => {
		fetchLeaderboard();
	}, [sort, page, search]);

	return (
		<div>
			
		</div>
	);
}
