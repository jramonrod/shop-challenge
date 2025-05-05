interface Props {
  total: string;
  onCheckout: () => void;
}

export function CartTotals({ total, onCheckout }: Props) {
  return (
    <div className="flex flex-col space-y-4">
      <div className="flex justify-between items-center">
        <span className="text-xl font-bold">Total</span>
        <span className="text-2xl font-bold">{total} €</span>
      </div>
      <button className="btn btn-primary btn-block" onClick={onCheckout}>
        Proceed to Checkout
      </button>
    </div>
  );
}
