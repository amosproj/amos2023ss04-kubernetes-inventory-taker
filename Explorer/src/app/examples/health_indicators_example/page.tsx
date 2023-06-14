import {
  HealthIndicatorWidget,
  HealthIndicatorBadge,
} from "@/components/health_indicators";

export default function Index(): JSX.Element {
  return (
    <div className="p-6">
      <div className="float-left">
        <HealthIndicatorBadge status={"Running"} />
        <HealthIndicatorBadge status={"Stopped"} />
        <HealthIndicatorBadge status={"Error"} />
        <HealthIndicatorBadge status={"Warning"} />
      </div>
      <div className="float-left">
        <HealthIndicatorWidget
          name={"Container ID 234767"}
          status={"Running"}
        />
      </div>
      <div className="float-left">
        <HealthIndicatorWidget
          name={"Container ID 234767"}
          status={"Stopped"}
        />
      </div>
      <div className="float-left">
        <HealthIndicatorWidget name={"Container ID 234767"} status={"Error"} />
      </div>
      <div className="float-left">
        <HealthIndicatorWidget
          name={"Container ID 234767"}
          status={"Warning"}
        />
      </div>
    </div>
  );
}
