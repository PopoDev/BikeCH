import React, { useState, useEffect, useRef } from "react";
import * as d3 from "d3";

export default function ElevationProfile({
  elevations,
  distances,
}: {
  elevations: number[];
  distances: number[];
}) {
  const chartRef = useRef<HTMLDivElement | null>(null);
  const [dimensions, setDimensions] = useState({ width: 0, height: 0 });

  useEffect(() => {
    const container = d3.select(chartRef.current);

    // Define a function to update the chart dimensions
    const updateDimensions = () => {
      container.selectAll("*").remove(); // Clear any existing content.
      const bounding = container.node()!.getBoundingClientRect();
      const width = bounding.width;
      const height = bounding.height;
      setDimensions({ width, height });
    };

    // Initial call to set the dimensions
    updateDimensions();

    // Add a resize event listener to update dimensions on window resize
    window.addEventListener("resize", updateDimensions);

    return () => {
      // Cleanup: remove the event listener on component unmount
      window.removeEventListener("resize", updateDimensions);
    };
  }, []);

  useEffect(() => {
    const container = d3.select(chartRef.current);
    container.selectAll("*").remove(); // Clear any existing content.
    const margin = { top: 20, right: 16, bottom: 20, left: 40 };

    // Create a new D3 scale for x-axis
    const xScale = d3
      .scaleLinear()
      .domain([0, d3.max(distances) || 0])
      .range([margin.left, dimensions.width - margin.right]);

    // Create a new D3 scale for y-axis
    const yScale = d3
      .scaleLinear()
      .domain([d3.min(elevations) || 0, d3.max(elevations) || 0])
      .range([dimensions.height - margin.bottom, margin.top]);

    const data = d3.zip(distances, elevations);

    // Define the area function
    const area = d3
      .area<number[]>()
      .y0(dimensions.height - margin.bottom)
      .x((d) => xScale(d[0]))
      .y1((d) => yScale(d[1]));

    // Create the SVG element
    const svg = d3
      .select(chartRef.current)
      .append("svg")
      .attr("width", dimensions.width)
      .attr("height", dimensions.height);

    // Append the area path to the SVG
    svg
      .append("path")
      .datum(data)
      .attr("fill", "red")
      .attr("d", area)
      .style("opacity", 0.6);

    // Add x-axis
    svg
      .append("g")
      .attr("transform", `translate(0,${dimensions.height - margin.bottom})`)
      .call(d3.axisBottom(xScale));

    // Add y-axis
    svg
      .append("g")
      .attr("transform", `translate(${margin.left},0)`)
      .call(d3.axisLeft(yScale));
  }, [elevations, dimensions]);

  return (
    <div
      ref={chartRef}
      className="elevation-profile flex-grow"
      style={{ width: "100%", height: "100%" }}
    ></div>
  );
}
