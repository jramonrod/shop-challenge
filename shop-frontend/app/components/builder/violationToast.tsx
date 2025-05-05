import type { Violation } from "~/types/violations";

interface Props {
  violations: Violation[];
}

export function ViolationsToast({ violations }: Props) {
  if (violations.length === 0) return null;
  return (
    <div className="toast toast-top toast-end z-50 space-y-2">
      {violations.map((v, i) => (
        <div key={i} className="alert alert-error shadow-lg">
          <span>{v.message}</span>
        </div>
      ))}
    </div>
  );
}
