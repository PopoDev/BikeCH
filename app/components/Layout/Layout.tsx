import Head from "next/head";

import Header from "../Header/Header";
import Footer from "../Footer/Footer";

import styles from "./Layout.module.scss";

interface LayoutProps {
  children: React.ReactNode;
  className?: string;
  [x: string]: any;
}

const Layout = ({ children, className, ...rest }: LayoutProps) => {
  return (
    <div className={styles.layout}>
      <Head>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Header />
      <main className={styles.main}>{children}</main>
      <Footer />
    </div>
  );
};

export default Layout;
