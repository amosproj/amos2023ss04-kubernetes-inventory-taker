"use client";
import { H1 } from "@/components/style_elements";
import { NavigationContext } from "@/context/sidebar_context";
import { useState } from "react";

/* @ts-expect-error Async Server Component */
export default async function Index(): JSX.Element {
  const current_page = useState("dashboard");

  return (
    <NavigationContext.Provider value={ current_page }>
      <div className="p-6">
        <H1 content={"Dashboard"} />
      </div>
    </NavigationContext.Provider>
  );
}
