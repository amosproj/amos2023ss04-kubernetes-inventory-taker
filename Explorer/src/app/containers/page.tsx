"use client";

import { H1 } from "@/components/style_elements";
import React from "react";
import StripedTable from "@/components/containerTable";
import containers from "@/components/containerTestData";
//import { useNavigationStore } from "@/context/sidebar_context";

export default async function Index(): Promise<JSX.Element> {
  //useNavigationStore((state) => state.changePage("nodes"));

  return (
    <div className="p-6">
      <H1 content={"Containers"} />
      <StripedTable containers={containers} />
    </div>
  );
}
