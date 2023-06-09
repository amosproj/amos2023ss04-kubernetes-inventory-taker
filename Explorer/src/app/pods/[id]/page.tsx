import PodDetailPage from "@/components/pod_detail_page";
import { getPodDetails } from "@/lib/db";
// Having to use this annotation is a known bug in React TypeScript
// https://nextjs.org/docs/app/building-your-application/data-fetching/fetching#async-and-await-in-server-components
export const dynamic = "force-dynamic";

export default async function Index({
  params,
}: {
  params: { id: string };
  /* @ts-expect-error Async Server Component */
}): JSX.Element {
  const data = await getPodDetails(decodeURIComponent(params.id));
  if (!data) {
    throw "Unkown ID";
  }
  return (
    <div className="p-6">
      <PodDetailPage pod_details={data.pod_data} containers={data.containers} />
    </div>
  );
}
