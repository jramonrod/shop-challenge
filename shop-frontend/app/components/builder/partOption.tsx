import type { Part } from "~/types/part";

interface Props {
  part: Part;
  isSelected: boolean;
  isInvalid: boolean;
  onChoose: (id: string) => void;
}

export function PartOption({ part, isSelected, isInvalid, onChoose }: Props) {
  const base = "cursor-pointer border rounded-lg p-4 flex flex-col items-center transition";
  const cls = isInvalid
    ? "border-red-500 bg-red-50 opacity-60 cursor-not-allowed"
    : isSelected
    ? "border-primary bg-primary/10"
    : "hover:border-primary hover:bg-primary/5";

  return (
    <div className={`${base} ${cls}`} onClick={() => !isInvalid && onChoose(part.uuid)}>
      <button className="btn btn-sm btn-outline mb-2" disabled={isInvalid}>
        {part.name}
      </button>
      <span className="text-sm text-gray-500">{part.price} €</span>
    </div>
  );
}
