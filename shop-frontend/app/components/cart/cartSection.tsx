import React from "react";
import type { Part } from "~/types/part";
import { CartItemCard } from "./cartItemCard";

interface Props {
  slug: string;
  parts: Part[];
}

export function CartSection({ slug, parts }: Props) {
  return (
    <section>
      <h2 className="text-xl font-semibold mb-2">{slug}</h2>
      <div className="space-y-2 mb-4">
        {parts.map((part) => (
          <CartItemCard key={part.uuid} part={part} />
        ))}
      </div>
    </section>
  );
}
