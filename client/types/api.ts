export interface User {
	ID?: number;
	id: number;
	name: string;
	email: string;
	phone: string;
	points: number;
	share_code: string;
	referred_by?: string;
	created_at: string;
	updated_at: string;
}

export interface RegisterRequest {
	name: string;
	email: string;
	phone_number: string;
	referral_code?: string;
}

export interface Response<T> {
	message?: string;
	error?: string;
	data?: T;
	leaderboard?: T;
	user?: T;
	total: number;
	page: number;
	totalPages: number;
}
