"use client";
import Head from "next/head";

import Layout from "./components/Layout/Layout";
import Section from "./components/Section/Section";
import Container from "./components/Container/Container";
import Map from "./components/Map/Map";
import Button from "./components/Button/Button";

import styles from "./Home.module.scss";

const DEFAULT_CENTER: [number, number] = [46.94831, 7.45997];

export default function Home() {
  return (
    <Layout>
      <Head>
        <title>Next.js Leaflet Starter</title>
        <meta
          name="description"
          content="Create mapping apps with Next.js Leaflet Starter"
        />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <Section>
        <Container>
          <h1 className={styles.title}>Next.js Leaflet Starter</h1>

          <Map
            className={styles.homeMap}
            width={800}
            height={400}
            center={DEFAULT_CENTER}
            zoom={12}
          />

          <p className={styles.description}>
            <code className={styles.code}>
              npx create-next-app -e
              https://github.com/colbyfayock/next-leaflet-starter
            </code>
          </p>

          <p className={styles.view}>
            <Button href="https://github.com/colbyfayock/next-leaflet-starter">
              Vew on GitHub
            </Button>
          </p>
        </Container>
      </Section>
    </Layout>
  );
}
