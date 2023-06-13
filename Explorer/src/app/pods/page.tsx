"use client";
import { H1 } from "@/components/style_elements";
//import { useNavigationStore } from "@/context/sidebar_context";

/* @ts-expect-error Async Server Component */
export default async function Index(): JSX.Element {
  //useNavigationStore((state) => state.changePage("dashboard"));

  return (
    <div className="p-6">
      <H1 content={"Pods"} />
    </div>
  );
}
