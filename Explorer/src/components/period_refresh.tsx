"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";

export default function PeriodicRefresh({
  delay_ms,
}: {
  delay_ms: number;
}): JSX.Element {
  const router = useRouter();
  useEffect(() => {
    const registration = setTimeout(() => {
      router.refresh();
    }, delay_ms);
    return () => {
      clearTimeout(registration);
    };
  });
  return <div></div>;
}
