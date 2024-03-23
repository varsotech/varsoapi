import { BrowserRouter, Route, Routes } from "react-router-dom";
import "./App.css";
import Home from "./routes/home/Home";
import { PersistQueryClientProvider } from "./api/reactquery";

function App() {
  return (
    <div className="App">
      <PersistQueryClientProvider>
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<Home />} />
          </Routes>
        </BrowserRouter>
      </PersistQueryClientProvider>
    </div>
  );
}

export default App;
