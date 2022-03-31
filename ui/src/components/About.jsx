import React from "react";
import { Link } from "react-router-dom";

export const About = () => {
  return (
    <>
      <h3>About</h3>
      <p>This is it folks, this is why we do it.</p>
      <p>
        Todos is an app that helps you (yes you) finally get organized and get
        your life together.
      </p>

      <Link className="nav-link" to={"/"}>
        Return
      </Link>
    </>
  );
};
