import ContainerTable from "@/components/containerTable";
import { getContainerList } from "@/lib/db";
export const dynamic = "force-dynamic";
/* @ts-expect-error Async Server Component */
export default async function Index(): JSX.Element {
  const data = await getContainerList();
  return (
    <div className="container mx-auto px-4">
      <h1 className="text-5xl mb-6 mt-10">Containers</h1>
      <ContainerTable list={data} />
    </div>
  );
}
