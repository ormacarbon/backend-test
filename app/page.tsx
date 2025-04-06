'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { Leaf } from 'lucide-react';
import LoginForm from '@/components/LoginForm';
import RegisterForm from '@/components/RegisterForm';

export default function Home() {
  const [isLogin, setIsLogin] = useState(true);
  const router = useRouter();

  const handleSuccess = (token: string) => {
    localStorage.setItem('token', token);
    router.push('/dashboard');
  };

  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-[url('https://images.unsplash.com/photo-1535913989690-f90e1c2d4cfa?auto=format&fit=crop&q=80')] bg-cover bg-center">
      <div className="absolute inset-0 bg-black/40 backdrop-blur-sm" />
      
      <div className="relative z-10 w-full max-w-md p-8 bg-white/95 rounded-lg shadow-xl">
        <div className="flex items-center justify-center mb-8">
          <Leaf className="w-8 h-8 text-[var(--accent)]" />
          <h1 className="text-3xl font-bold ml-2">vbio</h1>
        </div>

        {isLogin ? (
          <LoginForm onSuccess={handleSuccess} />
        ) : (
          <RegisterForm onSuccess={() => setIsLogin(true)} />
        )}

        <div className="mt-6 text-center">
          <button
            onClick={() => setIsLogin(!isLogin)}
            className="text-[var(--accent)] hover:text-[var(--accent-hover)] transition-colors"
          >
            {isLogin ? 'Criar uma nova conta' : 'Já tenho uma conta'}
          </button>
        </div>
      </div>
    </div>
  );
}