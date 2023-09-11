import dynamic from "next/dynamic";

const Map = dynamic(() => import("./DynamicMap"), {
  ssr: false,
});

export default Map;
