import React, { useState, useEffect } from "react";
import { Button } from "../components/ui/button";
import { Input } from "../components/ui/input";
import { Copy, Facebook, Twitter, MessageCircle } from "lucide-react";

interface ShareSectionProps {
  userId: string;
}

const ShareSection: React.FC<ShareSectionProps> = ({ userId }) => {
  const [shareUrl, setShareUrl] = useState("");
  // ver o uso de toast

  useEffect(() => {
    // Cria a URL de referência usando o ID do usuário
    const baseUrl = window.location.origin;
    setShareUrl(`${baseUrl}/register?ref=${userId}`);
  }, [userId]);

  const copyToClipboard = async () => {};

  const shareOnSocialMedia = (platform: string) => {
    let shareLink = "";
    const text = encodeURIComponent(
      "Join the Carbon Offset Competition and help spread awareness! Register using my link:",
    );

    switch (platform) {
      case "twitter":
        shareLink = `https://twitter.com/intent/tweet?text=${text}&url=${encodeURIComponent(
          shareUrl,
        )}`;
        break;
      case "facebook":
        shareLink = `https://www.facebook.com/sharer/sharer.php?u=${encodeURIComponent(
          shareUrl,
        )}`;
        break;
      case "whatsapp":
        shareLink = `https://wa.me/?text=${text}%20${encodeURIComponent(shareUrl)}`;
        break;
      default:
        return;
    }

    window.open(shareLink, "_blank");
  };

  return (
    <div className="space-y-4">
      <div className="flex items-center gap-2">
        <div className="relative flex-grow">
          <Input
            value={shareUrl}
            readOnly
            className="pr-10 border-slate-200 bg-slate-50 text-slate-800"
          />
          <button
            onClick={copyToClipboard}
            className="absolute right-2 top-1/2 -translate-y-1/2 text-slate-400 hover:text-emerald-600 focus:outline-none"
            aria-label="Copy to clipboard"
          >
            <Copy className="h-5 w-5" />
          </button>
        </div>
        <Button
          onClick={copyToClipboard}
          variant="outline"
          className="shrink-0 border-emerald-500 text-emerald-700 hover:bg-emerald-50"
        ></Button>
      </div>

      <div className="grid grid-cols-3 gap-3">
        <Button
          onClick={() => shareOnSocialMedia("twitter")}
          variant="outline"
          className="flex flex-col items-center gap-2 border-blue-400 py-6 text-blue-500 hover:bg-blue-50"
        >
          <Twitter className="h-5 w-5" />
          <span className="text-xs">Twitter</span>
        </Button>
        <Button
          onClick={() => shareOnSocialMedia("facebook")}
          variant="outline"
          className="flex flex-col items-center gap-2 border-blue-600 py-6 text-blue-600 hover:bg-blue-50"
        >
          <Facebook className="h-5 w-5" />
          <span className="text-xs">Facebook</span>
        </Button>
        <Button
          onClick={() => shareOnSocialMedia("whatsapp")}
          variant="outline"
          className="flex flex-col items-center gap-2 border-emerald-500 py-6 text-emerald-500 hover:bg-emerald-50"
        >
          <MessageCircle className="h-5 w-5" />
          <span className="text-xs">WhatsApp</span>
        </Button>
      </div>
    </div>
  );
};

export default ShareSection;
