import ContainerDetailPage from "@/components/container_detail_page";
import { getContainerDetails } from "@/lib/db";
// Having to use this annotation is a known bug in React TypeScript
// https://nextjs.org/docs/app/building-your-application/data-fetching/fetching#async-and-await-in-server-components
export const dynamic = "force-dynamic";

export default async function Index({
  params,
}: {
  params: { id: string };
  /* @ts-expect-error Async Server Component */
}): JSX.Element {
  const data = await getContainerDetails(parseInt(params.id));
  return (
    <div className="p-6">
      <ContainerDetailPage container_details={data} />
    </div>
  );
}
