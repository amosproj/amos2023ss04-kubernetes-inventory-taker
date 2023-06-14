import { H1 } from "@/components/style_elements";

/* @ts-expect-error Async Server Component */
export default async function Index(): JSX.Element {
  return (
    <div className="p-6">
      <H1 content={"Volumes"} />
    </div>
  );
}
