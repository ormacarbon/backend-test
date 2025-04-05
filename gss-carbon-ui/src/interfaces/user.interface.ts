export interface User {
  id: string;
  name: string;
  email: string;
  phone: string;
  points: number;
}

export interface Referral {
  id: string;
  createdAt: string;
  referrerId: string;
  referredId: string;
}

export interface UserRegister {
  name: string;
  email: string;
  phone: string;
  referralCode?: string | null;
}
