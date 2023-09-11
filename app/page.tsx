import Map from "./components/Map";

export default function Home() {
  return (
    <div className="text-center p-4">
      <h1 className="text-2xl md:text-3xl lg:text-4xl font-mono py-6">
        Bike route planner in Switzerland
      </h1>

      <Map width={800} height={400} />
    </div>
  );
}
