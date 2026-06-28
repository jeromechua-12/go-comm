import React, { useState } from 'react';
import { useNavigate } from 'react-router'
import "./login.css";

interface LoginResponse {
  data: UserInfo 
  error?: {
    type:     string
    message:  string 
  }
}

interface UserInfo {
    id:    number,
    name:  string
    email: string
    role:  string
}

function LoginPage() {
  const [formError, setFormError] = useState("");
  const navigate = useNavigate();

  const apiBase= import.meta.env.VITE_API_BASE_URL;

  const handleSubmit = async (e: React.SubmitEvent<HTMLFormElement>) => {
    e.preventDefault();

    const formData = new FormData(e.target);
    const payLoad = Object.fromEntries(formData.entries());

    console.log(payLoad);

    try {
      const response = await fetch(`${apiBase}/user/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payLoad),
        credentials: "include",
      });

      const result: LoginResponse = await response.json();

      if (!response.ok) {
        const error = result.error;
        setFormError(error.message);
        return
      }

      // navigate to account main page
      console.log("Successfully logged in!");

      navigate("/listing");
    } catch (err) {
      console.log(err);
      setFormError(err.message);
    }
  }

  return (
    <div className="login-page">
      {formError && <p className="form-error">{formError}</p>}
      <div className="login-form">
        <a href="/" className="close-btn" aria-label="Back to home">&times;</a>

        <h2>Login</h2>

        <form onSubmit={handleSubmit}>
          <div className="input-group">
            <label htmlFor="email">Email</label>
            <input
              type="email"
              id="email"
              name="email"
              required
            />
          </div>

          <div className="input-group">
            <label htmlFor="password">Password</label>
            <input
              type="password"
              id="password"
              name="password"
              required
            />
          </div>

          <button className="submit-btn">Login</button>
        </form>
      </div>
    </div>
  );
}

export default LoginPage;
