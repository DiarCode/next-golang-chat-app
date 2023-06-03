import { AppLayout } from "@/shared/layouts/app-layout";
import { Button } from "@/shared/ui/button";
import React from "react";
import { HomeOfferSection } from "./components/home-offer";
import Link from "next/link";
import { PAGES_LINKS } from "@/shared/config/links.config";

export const HomeScreen = () => {
  return (
    <AppLayout title="Home">
      <div className="flex flex-col justify-center items-center mt-16 md:mt-32">
        <p className="text-base tracking-wide font-semibold text-slate-500">
          Welcome, buddy!
        </p>
        <h1
          className="max-w-xl font-bold  text-3xl md:text-5xl text-center tracking-wide py-3 text-transparent bg-clip-text bg-gradient-to-r
             from-indigo-600 via-purple-600 to-pink-600"
        >
          Communication is easy with Meowchat
        </h1>

        <div className="mt-12 flex items-center gap-x-5">
          <Link href={PAGES_LINKS.Signup.link}>
            <Button solid>Get Started</Button>
          </Link>
          <Link href={PAGES_LINKS.Signup.link}>
            <Button>Create Account</Button>
          </Link>
        </div>
      </div>

      <HomeOfferSection />
    </AppLayout>
  );
};
