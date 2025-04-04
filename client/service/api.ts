import { User, RegisterRequest, Response } from '../types/api';

class Client {
	private readonly baseUrl: string;

	constructor() {
		this.baseUrl = process.env.API_URL || 'http://localhost:8000';
	}

	private async get<T>(endpoint: string, options?: RequestInit): Promise<Response<T>> {
		try {
			const response = await fetch(`${this.baseUrl}${endpoint}`, {
				headers: {
					'Content-Type': 'application/json',
				},
				...options,
			});

			const result = await response.json();

			if (!response.ok) throw new Error(result.error || `API request failed: ${endpoint}`);

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

	async getLeaderboard(): Promise<Response<User[]>> {
		return this.get<User[]>('/api/leaderboard');
	}
}

export const api = new Client();
export default api;