import React, { useState } from 'react';
import { useNavigate } from 'react-router'
import "./signup.css";

interface SignupResponse {
  data: string
  error?: {
    type:     string
    message:  string 
    details?: FieldErrors 
  }
}

interface FieldErrors {
  name?: string
  email?: string
  password?: string
}

function SignupPage() {
  const [fieldErrors, setFieldErrors] = useState<FieldErrors>();
  const [formError, setFormError] = useState("");
  const navigate = useNavigate()

  const apiBase= import.meta.env.VITE_API_BASE_URL;

  const handleSubmit = async (e: React.SubmitEvent<HTMLFormElement>) => {
    e.preventDefault();

    const formData = new FormData(e.target);
    const payLoad = Object.fromEntries(formData.entries());

    console.log(payLoad);

    try {
      const response = await fetch(`${apiBase}/user/signup`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payLoad),
      });

      const result: SignupResponse = await response.json();
      console.log(result)

      if (!response.ok) {
        const error = result.error;

        // check if validation error
        if (error.type == "validation_error") {
          console.log(error.details);
          setFieldErrors(error.details);
        }

        setFormError(error.message) 
        return
      }

      // navigate to login page
      console.log("Account created! Please log in")
      navigate("/login")

    } catch (err) {
      console.log(err);
      setFormError(err.message);
    }
  }

  return (
    <div className="signup-page">
      {formError && <p className="form-error">{formError}</p>}
      <div className="signup-form">
        <a href="/" className="close-btn" aria-label="Back to home">&times;</a>

        <h2>Create an account</h2>

        <form onSubmit={handleSubmit}>
          <div className="input-group">
            <label>Name</label>
            <input
              type="name"
              id="name"
              name="name"
              required
            />
              {fieldErrors?.name && <p className="field-error">{fieldErrors.name}</p>}
          </div>
          <div className="input-group">
            <label htmlFor="email">Email</label>
            <input
              type="email"
              id="email"
              name="email"
              required
            />
              {fieldErrors?.email && <p className="field-error">{fieldErrors.email}</p>}
          </div>

          <div className="input-group">
            <label htmlFor="password">Password</label>
            <input
              type="password"
              id="password"
              name="password"
              required
            />
              {fieldErrors?.password && <p className="field-error">{fieldErrors.password}</p>}
          </div>

          <button className="submit-btn">Sign up</button>
        </form>
      </div>
    </div>
  );
}

export default SignupPage;
