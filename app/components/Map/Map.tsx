"use client";

import "leaflet/dist/leaflet.css";
import { MapContainer, TileLayer, Marker, Popup } from "react-leaflet";
import styles from "./Map.module.scss";
import Icon from "../Icon";

// Set default sizing to control aspect ratio which will scale responsively
// but also help avoid layout shift
const DEFAULT_WIDTH = 800;
const DEFAULT_HEIGHT = 400;

const DEFAULT_CENTER: [number, number] = [46.94831, 7.45997];
const DEFAULT_ZOOM = 12;

interface MapProps {
  width?: number;
  height?: number;
  center?: [number, number];
  zoom?: number;
  scrollWheelZoom?: boolean;
}

export default function Map(props: MapProps) {
  const {
    width = DEFAULT_WIDTH,
    height = DEFAULT_HEIGHT,
    center = DEFAULT_CENTER,
    zoom = DEFAULT_ZOOM,
    scrollWheelZoom = true,
  } = props;

  return (
    <div style={{ aspectRatio: width / height }}>
      <MapContainer
        className={styles.map}
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
