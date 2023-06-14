"use client";

import { Badge } from "flowbite-react";

export type Health = "Running" | "Stopped" | "Error" | "Warning";
// Widget for component health
export function HealthIndicatorWidget({
  name,
  status,
}: {
  name: string;
  status: Health;
}): JSX.Element {
  function get_status_color() {
    switch (status) {
      case "Running":
        return "#DEF7EC";
      case "Stopped":
        return "#F3F4F6";
      case "Error":
        return "#FDE8E8";
      case "Warning":
        return "#FDF6B2";
    }
  }

  function get_status_text_color() {
    switch (status) {
      case "Running":
        return "#03543F";
      case "Stopped":
        return "#1F2937";
      case "Error":
        return "#9B1C1C";
      case "Warning":
        return "#713B13";
    }
  }

  return (
    <div className="flex p-2">
      <div
        style={{
          height: "200px",
          width: "200px",
          background: get_status_color(),
        }}
        className="bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700"
      >
        <p className="pt-4 pl-4 font-bold">Health of</p>
        <p className="pl-4">{name}</p>
        <p
          style={{ color: get_status_text_color() }}
          className="pt-4 font-bold  text-4xl text-center"
        >
          {status}
        </p>
      </div>
    </div>
  );
}

// Badge for component health
export function HealthIndicatorBadge({
  status,
}: {
  status: Health;
}): JSX.Element {
  function get_status_color_label() {
    switch (status) {
      case "Running":
        return "success";
      case "Stopped":
        return "gray";
      case "Error":
        return "failure";
      case "Warning":
        return "warning";
    }
  }

  return (
    <div className="flex p-2">
      <Badge color={get_status_color_label()} className="!text-2xl">
        {status}
      </Badge>
    </div>
  );
}
