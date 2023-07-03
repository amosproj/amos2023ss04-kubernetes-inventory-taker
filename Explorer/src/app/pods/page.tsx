import PodsTable from "@/components/podsTable";
import { H1 } from "@/components/style_elements";
import { getPodsList } from "@/lib/db";

/* @ts-expect-error Async Server Component */
export default async function Index(): JSX.Element {
  const data = await getPodsList();

  return (
    <div className="p-6">
      <H1 content={"Pods"} />
      <PodsTable list={data} />
    </div>
  );
}
