import React from "react";
import { useLoaderData, useNavigate } from "react-router";
import type { Category } from "~/types/category";
import type { Part } from "~/types/part";
import type { Product } from "~/types/products";
import type { VerifyResults } from "~/types/verify_results";
import type { Violation } from "~/types/violations";
import { PartOption } from "./partOption";
import { ViolationsToast } from "./violationToast";
import { Stepper } from "./stepper";

export async function loader({ params }: { params: { slug: string } }) {
  const prod_res = await fetch(`http://localhost:8080/products/${params.slug}`);
  const prod: Product = await prod_res.json();

  const res = await fetch(`http://localhost:8080/products/${prod.uuid}/categories`);
  if (!res.ok) throw new Error("Failed to load categories");
  const categories: Category[] = await res.json();
  return { product: prod, categories };
}

export default function BuilderPage() {
  const { product, categories } = useLoaderData() as {
    product: Product;
    categories: Category[];
  };
  const navigate = useNavigate();

  const [stepIndex, setStepIndex] = React.useState(0);
  const [selection, setSelection] = React.useState<Record<string, string>>({});
  const [parts, setParts] = React.useState<Part[]>([]);
  const [violations, setViolations] = React.useState<Violation[]>([]);
  const [loading, setLoading] = React.useState(false);
  const [error, setError] = React.useState<string | null>(null);

  const currentCategory = categories[stepIndex];

  // get parts
  React.useEffect(() => {
    let canceled = false;
    setLoading(true);
    setError(null);
    fetch(`http://localhost:8080/products/${product.uuid}/categories/${currentCategory.uuid}/parts`)
      .then((r) => { if (!r.ok) throw new Error("Failed to load parts"); return r.json(); })
      .then((data: Part[]) => !canceled && setParts(data))
      .catch((err) => !canceled && setError(err.message))
      .finally(() => !canceled && setLoading(false));
    return () => { canceled = true; };
  }, [product.uuid, currentCategory.uuid]);

  // validate
  React.useEffect(() => {
    let canceled = false;
    setLoading(true);
    fetch(`http://localhost:8080/products/${product.uuid}/restrictions/verify`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ parts: Object.values(selection) }),
    })
      .then((r) => { if (!r.ok) throw new Error("Failed to validate"); return r.json(); })
      .then((data: VerifyResults) => !canceled && setViolations(data.violations))
      .catch((err) => !canceled && setError(err.message))
      .finally(() => !canceled && setLoading(false));
    return () => { canceled = true; };
  }, [product.uuid, currentCategory.uuid, JSON.stringify(selection)]);

  function choose(partId: string) {
    if (violations.some((v) => v.resource_id === partId)) return;
    setSelection((s) => ({ ...s, [currentCategory.uuid]: partId }));
  }

  async function submitCart() {
    const res = await fetch("http://localhost:8080/cart", {
      method: "POST", headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ parts: Object.values(selection) }),
    });
    if (!res.ok) setError("failed to create cart");
    const { uuid } = await res.json();
    navigate(`/cart/${uuid}`, { replace: true });
  }

  function next() {
    if (stepIndex < categories.length - 1) setStepIndex(stepIndex + 1);
    else submitCart();
  }
  function back() { if (stepIndex > 0) setStepIndex(stepIndex - 1); }

  return (
    <div className="max-w-3xl mx-auto p-6 relative">
      <ViolationsToast violations={violations} />

      <Stepper categories={categories} stepIndex={stepIndex} onStepClick={setStepIndex} />

      <h2 className="text-xl font-bold mb-4">{currentCategory.name}</h2>
      {loading && <p>Loading options…</p>}
      {error && <p className="text-red-600">Error: {error}</p>}

      {!loading && !error && (
        <div className="grid grid-cols-2 gap-4 mb-6">
          {parts.map((part) => {
            const isSelected = selection[currentCategory.uuid] === part.uuid;
            const isInvalid = violations.some((v) => v.resource_id === part.uuid);
            return (
              <PartOption
                key={part.uuid}
                part={part}
                isSelected={isSelected}
                isInvalid={isInvalid}
                onChoose={choose}
              />
            );
          })}
        </div>
      )}

      <div className="flex justify-between">
        <button className="btn btn-outline" onClick={back} disabled={stepIndex === 0}>
          Back
        </button>
        <button className="btn btn-primary" onClick={next}
          disabled={!selection[currentCategory.uuid] || error != null}
        >
          {stepIndex < categories.length - 1 ? "Next" : "Add To Cart"}
        </button>
      </div>
    </div>
  );
}
