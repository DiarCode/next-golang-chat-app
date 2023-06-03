import { AppLayout } from "@/shared/layouts/app-layout";
import { Button } from "@/shared/ui/button";
import { InputField } from "@/shared/ui/input";
import React from "react";

export const LoginScreen = () => {
  return (
    <AppLayout title="Login">
      <main className="mx-auto max-w-md flex flex-col justify-center mt-44">
        <div className="flex flex-col justify-center items-center">
          <h1 className="font-bold text-3xl text-transparent bg-clip-text bg-gradient-to-r from-indigo-500 via-purple-500 to-pink-500">
            Meow, welcome back!
          </h1>
          <p className="mt-2 text-gray-500">Please enter your details.</p>
        </div>

        <div className="mt-12 flex flex-col gap-y-4">
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
          Login
        </Button>
      </main>
    </AppLayout>
  );
};
