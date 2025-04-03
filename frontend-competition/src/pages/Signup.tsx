import { useState } from "react";
import { useSearchParams } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { api } from "@/api/api";

export default function Signup() {
  const [formData, setFormData] = useState({
    name: "",
    email: "",
    phone: "",
  });

  const [searchParams] = useSearchParams();
  const referralCode = searchParams.get("ref") || "";

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      const response = await api.post("/signup", {
        ...formData,
        referredBy: referralCode,
      });

      alert(`Inscrição feita! Seu link de compartilhamento: ${response.data.share_link}`);
    } catch (error) {
      console.error("Erro:", error);
      alert("Erro ao se inscrever.");
    }
  };

  return (
    <div className="max-w-md mx-auto mt-10 p-6 bg-white shadow-md rounded-lg">
      <h2 className="text-xl font-semibold text-center mb-4">Inscreva-se na Competição</h2>
      <form onSubmit={handleSubmit} className="space-y-4">
        <Input name="name" placeholder="Nome" onChange={handleChange} required />
        <Input name="email" type="email" placeholder="E-mail" onChange={handleChange} required />
        <Input name="phone" placeholder="Telefone" onChange={handleChange} required />
        <Button type="submit" className="w-full">Participar</Button>
      </form>
    </div>
  );
}
