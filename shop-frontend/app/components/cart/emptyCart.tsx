import React from "react";
import { useNavigate } from "react-router";

export function EmptyCart() {
  const navigate = useNavigate();
  return (
    <div className="max-w-md mx-auto p-6 text-center">
      <h1 className="text-2xl font-bold mb-4">Your Cart is Empty</h1>
      <button className="btn btn-primary" onClick={() => navigate("/")}>
        Shop Now
      </button>
    </div>
  );
}
