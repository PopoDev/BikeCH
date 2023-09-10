import Container from "../Container/Container";

import styles from "./Footer.module.scss";

const Footer = ({ ...rest }) => {
  return (
    <footer className={styles.footer} {...rest}>
      <Container className={`${styles.footerContainer} ${styles.footerLegal}`}>
        <p>{new Date().getFullYear()}</p>
      </Container>
    </footer>
  );
};

export default Footer;
