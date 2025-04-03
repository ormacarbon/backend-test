import { useEffect, useState } from "react";
import { useNavigate, useSearchParams } from "react-router-dom"; // Importa useNavigate
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { api } from "@/api/api";

export default function Signup() {
  const navigate = useNavigate(); // Hook para redirecionamento
  const [formData, setFormData] = useState({
    name: "",
    email: "",
    phone_number: "",
  });

  const [searchParams] = useSearchParams();
  const [referralCode, setReferralCode] = useState("");

  useEffect(() => {
    const ref = searchParams.get("ref") || "";
    console.log("Código de indicação atualizado:", ref);
    setReferralCode(ref);
  }, [searchParams]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
  
    const data = {
      ...formData,
      referred_By: referralCode,
    };
  
  
    try {
      const response = await api.post("/api/signup", data);
      const shareLink = response.data.share_link;
      alert(`Inscrição feita! Seu link de compartilhamento: ${shareLink}`);
      navigate(`/share?link=${encodeURIComponent(shareLink)}`);
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
        <Input name="phone_number" placeholder="Telefone" onChange={handleChange} required />
        <Button type="submit" className="w-full">Participar</Button>
      </form>
    </div>
  );
}
