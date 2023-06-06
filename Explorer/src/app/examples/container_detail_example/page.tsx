"use client";

import ContainerDetailPage from "@/components/container_detail_page";
import { getContainerDetails } from "@/lib/db";
/* @ts-expect-error Async Server Component */
export default async function Index(): JSX.Element {
  return (
    <div className="p-6">
      <ContainerDetailPage
        container_details={await getContainerDetails(undefined)}
      />
    </div>
  );
}
