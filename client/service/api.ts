import { User, RegisterRequest, Response } from '../types/api';

class Client {
	private readonly baseUrl: string;

	constructor() {
		this.baseUrl = process.env.API_URL || 'http://localhost:8080';
	}

	private async get<T>(end: string, options?: RequestInit): Promise<Response<T>> {
		try {
			const response = await fetch(`${this.baseUrl}${end}`, {
				headers: {
					'Content-Type': 'application/json',
				},
				...options,
			});

			const result = await response.json();

			if (!response.ok) throw new Error(result.error || `API request failed: ${end}`);

			return result;
		} catch (error) {
			throw new Error('Unknown API error occurred');
		}
	}

	async register(data: RegisterRequest): Promise<Response<User>> {
		return this.get<User>('/api/register', {
			method: 'POST',
			body: JSON.stringify(data),
		});
	}

	async getLeaderboard(params?: { sort?: string; page?: string; search?: string }): Promise<Response<User[]>> {
		const query = new URLSearchParams(params).toString();
		const end = query ? `/api/leaderboard?${query}` : '/api/leaderboard';
		return this.get<User[]>(end, {
			method: 'GET',
		});
	}

	async getShareLink(id: string): Promise<Response<User>> {
		return this.get<User>(`/api/share/${id}`);
	}
}

export const api = new Client();
export default api;