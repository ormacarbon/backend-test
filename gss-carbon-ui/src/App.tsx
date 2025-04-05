import React from "react";
import Routes from "./routes/Routes";
import { QueryClientProvider } from "@tanstack/react-query";
import queryClient from "./query/queryClient";
import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

const App: React.FC = () => {
  return (
    <QueryClientProvider client={queryClient}>
      <Routes />
      <ToastContainer position="bottom-right" autoClose={5000} icon={false} />
    </QueryClientProvider>
  );
};

export default App;
