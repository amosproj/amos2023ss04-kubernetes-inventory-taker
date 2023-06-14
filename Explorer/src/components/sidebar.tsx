"use client";

import classNames from "classnames";
import { Sidebar as FlowbiteSidebar } from "flowbite-react";
import type { FC, PropsWithChildren } from "react";
import { useSidebarContext } from "@/context/sidebar_context";
import {
  HiPresentationChartLine,
  HiCog,
  HiFolder,
  HiServer,
  HiTemplate,
  HiChip,
  HiTable,
  HiHome,
} from "react-icons/hi";

const Sidebar: FC<PropsWithChildren<Record<string, unknown>>> = function ({
  children,
}) {
  const { isOpenOnSmallScreens: isSidebarOpenOnSmallScreens } =
    useSidebarContext();

  return (
    <div
      className={classNames(
        "fixed overflow-auto top-0 h-screen z-10 lg:sticky lg:!block",
        {
          hidden: !isSidebarOpenOnSmallScreens,
        }
      )}
    >
      <FlowbiteSidebar>{children}</FlowbiteSidebar>
    </div>
  );
};

export default Object.assign(Sidebar, { ...FlowbiteSidebar });

import { usePathname } from "next/navigation";

export function ActualSidebar(): JSX.Element {
  const pathname = usePathname();
  const splitPathname = pathname.split("/");
  const current_page = splitPathname[1];

  return (
    <FlowbiteSidebar>
      <FlowbiteSidebar.Items>
        <FlowbiteSidebar.ItemGroup>
          <FlowbiteSidebar.Item
            href="dashboard"
            icon={HiPresentationChartLine}
            active={current_page.localeCompare("dashboard") ? false : true}
          >
            Dashboard
          </FlowbiteSidebar.Item>
          <FlowbiteSidebar.Item
            href="cluster"
            icon={HiHome}
            active={current_page.localeCompare("cluster") ? false : true}
          >
            Cluster
          </FlowbiteSidebar.Item>
          <FlowbiteSidebar.Item
            href="deployments"
            icon={HiTemplate}
            active={current_page.localeCompare("deployments") ? false : true}
          >
            Deployments
          </FlowbiteSidebar.Item>
          <FlowbiteSidebar.Item
            href="nodes"
            icon={HiTemplate}
            active={current_page.localeCompare("nodes") ? false : true}
          >
            Nodes
          </FlowbiteSidebar.Item>
          <FlowbiteSidebar.Item
            href="pods"
            icon={HiTable}
            active={current_page.localeCompare("pods") ? false : true}
          >
            Pods
          </FlowbiteSidebar.Item>
          <FlowbiteSidebar.Item
            href="containers"
            icon={HiFolder}
            active={current_page.localeCompare("containers") ? false : true}
          >
            Containers
          </FlowbiteSidebar.Item>
          <FlowbiteSidebar.Item
            href="volumes"
            icon={HiServer}
            active={current_page.localeCompare("volumes") ? false : true}
          >
            Volumes
          </FlowbiteSidebar.Item>
          <FlowbiteSidebar.Item
            href="services"
            icon={HiChip}
            active={current_page.localeCompare("services") ? false : true}
          >
            Services
          </FlowbiteSidebar.Item>
        </FlowbiteSidebar.ItemGroup>
        <FlowbiteSidebar.ItemGroup>
          <FlowbiteSidebar.Item
            href="settings"
            icon={HiCog}
            active={current_page.localeCompare("settings") ? false : true}
          >
            Settings
          </FlowbiteSidebar.Item>
        </FlowbiteSidebar.ItemGroup>
      </FlowbiteSidebar.Items>
    </FlowbiteSidebar>
  );
}
