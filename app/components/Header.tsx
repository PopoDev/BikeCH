import Link from "next/link";
import { FaGithub } from "react-icons/fa";

import Logo from "./Logo";

export default function Header() {
  return (
    <header className="bg-black text-white w-full fixed top-0 left-0">
      <div className="w-4/5 m-auto flex justify-between items-center">
        <Logo />
        <Link
          href="https://github.com/PopoDev/BikeCH"
          className="p-4"
          target="_blank"
          rel="noopener noreferrer"
        >
          <FaGithub size="28px" />
        </Link>
      </div>
    </header>
  );
}
