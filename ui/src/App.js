import { Todos } from "./components/Todos";
import { About } from "./components/About";
import { BrowserRouter, Routes, Route } from "react-router-dom";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route
          path="/"
          element={
            <div className="container">
              <Todos />
            </div>
          }
        />
        <Route
          path="/about"
          element={
            <div className="container">
              <About />
            </div>
          }
        />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
