import { useSearchParams } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { useState } from "react";

export default function Share() {
  const [copied, setCopied] = useState(false);
  const [searchParams] = useSearchParams();

  // Pega o link correto da query string
  const shareLink = searchParams.get("link") || "";

  const copyToClipboard = async () => {
    if (!shareLink) return;
    await navigator.clipboard.writeText(shareLink);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  return (
    <div className="max-w-md mx-auto mt-10 p-6 bg-white shadow-md rounded-lg text-center">
      <h2 className="text-xl font-semibold mb-4">Convide Amigos!</h2>
      <p className="text-gray-600 mb-4">Compartilhe este link e ganhe pontos extras:</p>
      
      <input
        type="text"
        value={shareLink}
        readOnly
        className="w-full p-2 border rounded mb-4 text-center"
      />

      <Button onClick={copyToClipboard} className="w-full">
        {copied ? "Copiado!" : "Copiar Link"}
      </Button>
    </div>
  );
}
