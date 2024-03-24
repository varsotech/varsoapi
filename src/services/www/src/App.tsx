import { BrowserRouter, Route, Routes } from "react-router-dom";
import Home from "./routes/home/Home";
import { PersistQueryClientProvider } from "./api/reactquery";
import Login from "./routes/login/Login";
import { AuthProvider } from "components/Auth/AuthProvider";

function App() {
  return (
    <div className="App">
      <PersistQueryClientProvider>
        <AuthProvider>
          <BrowserRouter>
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="/login" element={<Login />} />
            </Routes>
          </BrowserRouter>
        </AuthProvider>
      </PersistQueryClientProvider>
    </div>
  );
}

export default App;
