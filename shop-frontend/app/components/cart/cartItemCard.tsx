import type { Part } from "~/types/part";

interface Props {
  part: Part;
}

export function CartItemCard({ part }: Props) {
  return (
    <div className="card bg-base-100 shadow-md">
      <div className="card-body flex items-center">
        <div className="flex-1">
          <p className="text-lg">{part.name}</p>
        </div>
        <div className="text-right">
          <span className="text-xl font-semibold">{part.price} €</span>
        </div>
      </div>
    </div>
  );
}
