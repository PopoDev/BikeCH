import Map from "./components/Map";

export default function Home() {
  return (
    <div className="flex-grow flex flex-col text-center px-4 pt-[60px] pb-12 ">
      <h1 className="text-lg sm:text-2xl md:text-3xl lg:text-4xl font-mono py-6">
        Bike route planner in Switzerland
      </h1>

      <Map />
    </div>
  );
}
