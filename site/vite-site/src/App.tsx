import { Products } from "./components/Products";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
const queryClient = new QueryClient();

function App() {
  return (
    <div className="container mx-auto py-8">
      <QueryClientProvider client={queryClient}>
        <Products />
      </QueryClientProvider>
    </div>
  );
}

export default App;
