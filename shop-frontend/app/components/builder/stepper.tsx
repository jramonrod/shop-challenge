import type { Category } from "~/types/category";

interface Props {
  categories: Category[];
  stepIndex: number;
  onStepClick?: (i: number) => void;
}

export function Stepper({ categories, stepIndex, onStepClick }: Props) {
  return (
    <ul className="steps w-full mb-6">
      {categories.map((c, i) => (
        <li
          key={c.uuid}
          onClick={() => onStepClick && onStepClick(i)}
          className={
            "step " +
            (i < stepIndex
              ? "step-primary"
              : i === stepIndex
              ? "step-secondary"
              : "")
          }
        >
          {c.name}
        </li>
      ))}
    </ul>
  );
}
