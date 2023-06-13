import ContainerTable from "@/components/containerTable";
import { getContainerList } from "@/lib/db";
import { H1 } from "@/components/style_elements";

export const dynamic = "force-dynamic";
/* @ts-expect-error Async Server Component */
export default async function Index(): JSX.Element {
  const data = await getContainerList();
  return (
    <div className="container mx-auto px-4">
      <H1 content={"Containers"} />
      <ContainerTable list={data} />
    </div>
  );
}
