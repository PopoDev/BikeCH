import Leaflet from "leaflet";
import * as ReactLeaflet from "react-leaflet";
import "leaflet/dist/leaflet.css";

import styles from "./Map.module.scss";

const { MapContainer } = ReactLeaflet;

interface MapProps {
  children: (ReactLeaflet: any, Leaflet: any) => React.ReactNode;
  className?: string;
  width?: number;
  height?: number;
  [x: string]: any;
}

const Map = ({ children, className, width, height, ...rest }: MapProps) => {
  let mapClassName = styles.map;

  if (className) {
    mapClassName = `${mapClassName} ${className}`;
  }

  return (
    <MapContainer className={mapClassName} {...rest}>
      {children(ReactLeaflet, Leaflet)}
    </MapContainer>
  );
};

export default Map;
