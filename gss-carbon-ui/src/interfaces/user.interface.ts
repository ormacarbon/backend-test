export interface User {
  id: number;
  name: string;
  email: string;
  phone: string;
  points: number;
  referralToken: string;
  referredBy: number;
  createdAt: string;
  updatedAt: string;
}

export interface ReducedUser {
  id: number;
  name: string;
  email: string;
}

export interface UserRegister {
  name: string;
  email: string;
  phone: string;
  referralToken?: string | null;
}

export interface Referrals {
  id: number;
  referrer: ReducedUser;
  referred: ReducedUser;
  createdAt: string;
}