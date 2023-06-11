"use client";

import classNames from "classnames";
import { Sidebar as FlowbiteSidebar } from "flowbite-react";
import type { FC, PropsWithChildren } from "react";
import { useContext } from "react";
import {
  useSidebarContext,
  NavigationContext,
} from "@/context/sidebar_context";
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

export function ActualSidebar(): JSX.Element {
  const current_page = useContext(NavigationContext).current_page;

  return (
    <FlowbiteSidebar>
      <FlowbiteSidebar.Items>
        <FlowbiteSidebar.ItemGroup>
          <FlowbiteSidebar.Item
            href="dashboard"
            icon={HiPresentationChartLine}
            active={current_page == "dashboard" ? true : false}
          >
            Dashboard
          </FlowbiteSidebar.Item>
          <FlowbiteSidebar.Item
            href="cluster"
            icon={HiHome}
            active={current_page == "cluster" ? true : false}
          >
            Cluster
          </FlowbiteSidebar.Item>
          <FlowbiteSidebar.Item
            href="nodes"
            icon={HiTemplate}
            active={current_page == "nodes" ? true : false}
          >
            Nodes
          </FlowbiteSidebar.Item>
          <FlowbiteSidebar.Item
            href="pods"
            icon={HiTable}
            active={current_page == "pods" ? true : false}
          >
            Pods
          </FlowbiteSidebar.Item>
          <FlowbiteSidebar.Item
            href="containers"
            icon={HiFolder}
            active={current_page == "containers" ? true : false}
          >
            Containers
          </FlowbiteSidebar.Item>
          <FlowbiteSidebar.Item
            href="volumes"
            icon={HiServer}
            active={current_page == "volumes" ? true : false}
          >
            Volumes
          </FlowbiteSidebar.Item>
          <FlowbiteSidebar.Item
            href="services"
            icon={HiChip}
            active={current_page == "services" ? true : false}
          >
            Services
          </FlowbiteSidebar.Item>
        </FlowbiteSidebar.ItemGroup>
        <FlowbiteSidebar.ItemGroup>
          <FlowbiteSidebar.Item
            href="settings"
            icon={HiCog}
            active={current_page == "settings" ? true : false}
          >
            Settings
          </FlowbiteSidebar.Item>
        </FlowbiteSidebar.ItemGroup>
      </FlowbiteSidebar.Items>
    </FlowbiteSidebar>
  );
}
