import { PAGES_LINKS } from "@/shared/config/links.config";
import { AppLayout } from "@/shared/layouts/app-layout";
import { Button } from "@/shared/ui/button";
import { InputField } from "@/shared/ui/input";
import Link from "next/link";
import React from "react";

export const SignupScreen = () => {
  return (
    <AppLayout title="Signup">
      <main className="mx-auto max-w-md flex flex-col justify-center mt-44">
        <div className="flex flex-col justify-center items-center">
          <h1 className="font-bold text-3xl text-center text-transparent bg-clip-text bg-gradient-to-r from-indigo-500 via-purple-500 to-pink-500">
            Meow, Create Account!
          </h1>
          <p className="mt-2 text-gray-500">Please enter your details.</p>
        </div>

        <div className="mt-8 md:mt-12 flex flex-col gap-y-4">
          <InputField
            label="Username"
            type="text"
            placeholder="Enter your username"
          />

          <InputField
            label="Email"
            type="email"
            placeholder="Enter your email"
          />

          <InputField
            label="Password"
            type="password"
            placeholder="Enter your password"
          />
        </div>

        <Button style={{ marginTop: 34 }} solid>
          Signup
        </Button>

        <div className="flex justify-center items-center gap-x-2 text-sm mt-8 md:mt-12">
          <p>Already have an account?</p>
          <Link href={PAGES_LINKS.Login.link}>
            <p className="text-blue-600 underline text-sm">Login</p>
          </Link>
        </div>
      </main>
    </AppLayout>
  );
};
