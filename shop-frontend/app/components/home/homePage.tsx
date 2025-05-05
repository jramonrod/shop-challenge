// src/components/Home/HomePage.tsx
import React from "react";
import { useLoaderData } from "react-router";
import type { Product } from "~/types/products";
import { ProductGrid } from "./productGrid";
import type { Route } from "./+types/homePage";

export async function loader({}: Route.LoaderArgs): Promise<{ products: Product[] }> {
  const res = await fetch("http://localhost:8080/products");
  if (!res.ok) throw new Error("Failed to load products");
  const products: Product[] = await res.json();
  return { products };
}

export default function HomePage() {
  const { products } = useLoaderData<typeof loader>();
  return <ProductGrid products={products} />;
}
