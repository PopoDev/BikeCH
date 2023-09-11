import { icon } from "leaflet";

const Icon = icon({
  iconUrl: "/leaflet/images/marker-icon.png", // "leaflet/dist/images/marker-icon.png
  iconRetinaUrl: "/leaflet/images/marker-icon-2x.png", // "leaflet/dist/images/marker-icon-2x.png
  shadowUrl: "/leaflet/images/marker-shadow.png", // "leaflet/dist/images/marker-shadow.png
  iconSize: [25, 41],
  iconAnchor: [12, 41],
  popupAnchor: [1, -34],
  tooltipAnchor: [16, -28],
  shadowSize: [41, 41],
});

export default Icon;
