import ContainerDetailPage from "@/components/container_detail_page";
import { getData } from "@/lib/db";

export default function Index(): JSX.Element {
  return (
    <div className="p-6">
      <ContainerDetailPage container={await getData()} />
    </div>
  );
}
