import AppRoutes from "./routes/routes";

export default function App() {
  return (
    <div className="min-h-screen flex flex-col bg-gray-100">
      <header className="p-4 bg-white shadow-md">
        <h1 className="text-xl font-bold text-center">Competição de Carbono</h1>
      </header>
      <main className="flex-1 flex justify-center items-center">
        <AppRoutes />
      </main>
    </div>
  );
}
