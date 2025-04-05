import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "../pages/Home";
import Admin from "../pages/Admin";
import Leaderboard from "../pages/Leaderboard";
import RegistrationSuccessful from "@/pages/RegistrationSuccessful";
import Register from "@/pages/Register";

const AppRoutes: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/admin" element={<Admin />} />
        <Route path="/leaderboard" element={<Leaderboard />} />
        <Route path="/register" element={<Register />} />
        <Route path="/success/:userId" element={<RegistrationSuccessful />} />
      </Routes>
    </Router>
  );
};

export default AppRoutes;
