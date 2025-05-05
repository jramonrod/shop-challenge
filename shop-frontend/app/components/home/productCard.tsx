import type { Product } from "~/types/products";

interface Props {
  product: Product;
}

export function ProductCard({ product }: Props) {
  const isUnavailable = !product.is_available;

  return (
    <div className="relative group overflow-hidden diagonal-cut rounded-lg shadow-lg transform transition-transform hover:scale-105 hover:shadow-2xl">
      <a
        {...(!isUnavailable ? { href: `/products/${product.slug}` } : {})}
        className={`block ${isUnavailable ? "pointer-events-none cursor-not-allowed" : ""}`}
        aria-disabled={isUnavailable}
      >
        <img
          src={product.location}
          alt={product.name}
          className={`${isUnavailable ? "opacity-50" : ""} w-full h-auto block`}
        />
        {isUnavailable && (
          <div className="absolute inset-0 bg-black bg-opacity-50 flex items-center justify-center text-white font-bold text-xl opacity-0 group-hover:opacity-50 transition-opacity">
            Coming Soon!
          </div>
        )}
      </a>
    </div>
  );
}
