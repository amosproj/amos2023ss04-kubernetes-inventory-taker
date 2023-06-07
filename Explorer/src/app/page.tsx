"use client";
import React from "react";
import StripedTable from "@/components/containerTable";
import containers from "@/components/containerTestData";

const ContainerPage: React.FC = () => {
  return (
    <div className="container mx-auto px-4">
      <h1 className="text-5xl mb-6 mt-10">Containers</h1>
      <StripedTable containers={containers} />
    </div>
  );
};

export default ContainerPage;
