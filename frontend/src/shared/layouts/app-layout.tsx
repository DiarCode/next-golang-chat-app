import { Navbar } from "@/widgets/navbar";
import Head from "next/head";
import { CSSProperties, FC, PropsWithChildren, useEffect } from "react";

interface AppLayoutProps extends PropsWithChildren {
  title: string;
  description?: string;
  containerized?: boolean;
  style?: CSSProperties;
  showNavbar?: boolean;
}

const metaDescriptionContent = "chat, posts, community";

export const AppLayout: FC<AppLayoutProps> = ({
  title,
  description,
  children,
  containerized = true,
  style,
  showNavbar = true,
}) => {
  useEffect(() => {
    window.scrollTo(0, 0);
  }, []);

  return (
    <div className="bg-white">
      <Head>
        <title>{title}</title>
        <meta name="Keywords" content={metaDescriptionContent} />
        <meta name="description" content={description} />
        <meta
          name="viewport"
          content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0"
        />
      </Head>

      <main
        className={`w-screen min-h-screen relative overflow-clip bg-white font-open_sans text-sm sm:text-base 
          mx-auto px-3 pb-10 text-black ${containerized && "container"}`}
        style={style}
      >
        {showNavbar && <Navbar />}
        {children}
      </main>

      {/* <AlertNotification /> */}
    </div>
  );
};
