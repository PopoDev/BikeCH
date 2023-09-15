import { icon } from "leaflet";

export default function MarkerIcon(
  color:
    | "black"
    | "blue"
    | "gold"
    | "green"
    | "grey"
    | "orange"
    | "red"
    | "violet"
    | "yellow"
) {
  return icon({
    iconUrl: `/leaflet/marker-icon-${color}.png`,
    iconRetinaUrl: `/leaflet/marker-icon-2x-${color}.png`,
    shadowUrl: "/leaflet/marker-shadow.png",
    iconSize: [25, 41],
    iconAnchor: [12, 41],
    popupAnchor: [1, -34],
    tooltipAnchor: [16, -28],
    shadowSize: [41, 41],
  });
}
