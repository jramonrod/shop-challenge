import { useLoaderData, useNavigate } from "react-router";
import type { Cart } from "~/types/cart";
import { EmptyCart } from "./emptyCart";
import { CartSection } from "./cartSection";
import { CartTotals } from "./cartTotals";
interface LoaderData {
  cart: Cart;
}

export async function loader({
  params,
}: {
  params: { cart_id: string };
}): Promise<LoaderData> {
  const res = await fetch(`http://localhost:8080/carts/${params.cart_id}`);
  if (!res.ok) throw new Error("Failed to load cart");
  const cart: Cart = await res.json();
  return { cart };
}

export default function CartPage() {
  const { cart } = useLoaderData<LoaderData>();
  const navigate = useNavigate();

  const entries = Object.entries(cart.products);
  const total = entries
    .flatMap(([_, parts]) => parts)
    .reduce((sum, part) => sum + parseFloat(part.price), 0)
    .toFixed(2);

  if (entries.length === 0) {
    return <EmptyCart />;
  }

  return (
    <div className="max-w-3xl mx-auto p-6 space-y-6">
      <h1 className="text-2xl font-bold">Your Cart</h1>

      {entries.map(([slug, parts]) => (
        <CartSection key={slug} slug={slug} parts={parts} />
      ))}

      <CartTotals
        total={total}
        onCheckout={() =>
          navigate("/checkout", { state: { cartId: cart.uuid } })
        }
      />
    </div>
  );
}
