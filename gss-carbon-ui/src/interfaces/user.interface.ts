export interface User {
  id: number;
  name: string;
  email: string;
  phone: string;
  points: number;
  referralToken: string;
  createdAt: string;
  updatedAt: string;
}

export interface UserRegister {
  name: string;
  email: string;
  phone: string;
  referralCode?: string | null;
}
