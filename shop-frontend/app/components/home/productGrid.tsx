import type { Product } from "~/types/products";
import { ProductCard } from "./productCard";

interface Props {
  products: Product[];
}

export function ProductGrid({ products }: Props) {
  return (
    <div className="max-w-4xl mx-auto p-4">
      <div className="grid grid-cols-2 gap-4">
        {products.map((p) => (
          <ProductCard key={p.uuid} product={p} />
        ))}
      </div>
    </div>
  );
}
