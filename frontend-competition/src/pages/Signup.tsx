import { useEffect, useState } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { api } from "@/api/api";

export default function Signup() {
  const navigate = useNavigate(); 
  const [formData, setFormData] = useState({
    name: "",
    email: "",
    phone_number: "",
  });

  const [searchParams] = useSearchParams();
  const [referralCode, setReferralCode] = useState("");

  useEffect(() => {
    const ref = searchParams.get("ref") || "";
    console.log("Updated referral code:", ref);
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
      alert(`Signup successful! Your sharing link: ${shareLink}`);
      navigate(`/share?link=${encodeURIComponent(shareLink)}`);
    } catch (error) {
      console.error("Error:", error);
      alert("Failed to sign up.");
    }
  };

  return (
    <div className="max-w-md mx-auto mt-10 p-6 bg-white shadow-md rounded-lg">
      <h2 className="text-xl font-semibold text-center mb-4">Sign Up for the Competition</h2>
      <form onSubmit={handleSubmit} className="space-y-4">
        <Input name="name" placeholder="Name" onChange={handleChange} required />
        <Input name="email" type="email" placeholder="Email" onChange={handleChange} required />
        <Input name="phone_number" placeholder="Phone Number" onChange={handleChange} required />
        <Button type="submit" className="w-full">Join</Button>
      </form>
    </div>
  );
}
