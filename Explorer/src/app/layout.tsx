"use client";

import "@/styles/globals.css";
import { Header } from "@/components/header";
import { SidebarProvider } from "@/context/sidebar_context";
import { ActualSidebar } from "@/components/sidebar";
import { Flowbite } from "flowbite-react";
import PeriodicRefresh from "@/components/period_refresh";

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>
        <PeriodicRefresh delay_ms={1000} />

        <Flowbite>
          <SidebarProvider>
            <div className="flex dark:bg-gray-900 h-screen">
              <div className="order-1 float-left bg-black">
                <ActualSidebar />
              </div>
              <main className="order-2 w-full overflow-x-hidden">
                <div className="float-left w-full">
                  <Header />
                </div>
                <div className="float-left w-full">
                  <div className="mx-4 mt-4 mb-24">{children}</div>
                </div>
              </main>
            </div>
          </SidebarProvider>
        </Flowbite>
      </body>
    </html>
  );
}
