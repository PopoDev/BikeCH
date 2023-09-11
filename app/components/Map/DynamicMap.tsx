import "leaflet/dist/leaflet.css";
import { MapContainer, TileLayer, Marker, Popup } from "react-leaflet";
import Icon from "../Icon";

import styles from "./Map.module.scss";

// Set default sizing to control aspect ratio which will scale responsively
// but also help avoid layout shift
const DEFAULT_WIDTH = 600;
const DEFAULT_HEIGHT = 600;

interface MapProps {
  className?: string;
  width?: number;
  height?: number;
  center: [number, number];
  zoom: number;
  scrollWheelZoom?: boolean;
}

export default function Map(props: MapProps) {
  const {
    className,
    width = DEFAULT_WIDTH,
    height = DEFAULT_HEIGHT,
    center,
    zoom,
    scrollWheelZoom = true,
  } = props;

  let mapClassName = styles.map;

  if (className) {
    mapClassName = `${mapClassName} ${className}`;
  }

  return (
    <div style={{ aspectRatio: width / height }}>
      <MapContainer
        className={mapClassName}
        center={center}
        zoom={zoom}
        scrollWheelZoom={scrollWheelZoom}
      >
        <TileLayer
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
        />
        <Marker position={center} icon={Icon}>
          <Popup>
            A pretty CSS3 popup. <br /> Easily customizable.
          </Popup>
        </Marker>
      </MapContainer>
    </div>
  );
}
