import { BrowserRouter, Routes, Route } from "react-router";
import Home from "./pages/homepage";
import SignupPage from "./pages/user/signup";
import LoginPage from "./pages/user/login";
import ListingPage from "./pages/product/listing";
import "./App.css"

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/signup" element={<SignupPage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/listing" element={<ListingPage />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
