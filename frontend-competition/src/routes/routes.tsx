import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Share from "./pages/Share";
import Leaderboard from "./pages/Leaderboard";
import Signup from "@/pages/Signup";

export default function AppRoutes() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Signup />} />
        <Route path="/share" element={<Share />} />
        <Route path="/leaderboard" element={<Leaderboard />} />
      </Routes>
    </Router>
  );
}
