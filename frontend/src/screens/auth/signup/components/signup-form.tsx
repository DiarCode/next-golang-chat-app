import { useAuth } from "@/shared/hooks/useAuth";
import { SignupDto } from "@/shared/types/auth/auth.dto";
import { Button } from "@/shared/ui/button";
import { InputField } from "@/shared/ui/input";
import React from "react";
import { useForm } from "react-hook-form";

export const SignupForm = () => {
  const { signup } = useAuth();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<SignupDto>();

  const onSignupSubmit = (data: SignupDto) => {
    console.log(data);
    signup(data);
  };

  return (
    <form
      onSubmit={handleSubmit(onSignupSubmit)}
      className="mt-8 md:mt-12 flex flex-col gap-y-4"
    >
      <InputField
        label="Username"
        type="text"
        placeholder="Enter your username"
        error={errors.username?.message}
        {...register("username", { required: "Username is required" })}
      />

      <InputField
        label="Email"
        type="email"
        placeholder="Enter your email"
        error={errors.email?.message}
        {...register("email", { required: "Email is required" })}
      />

      <InputField
        label="Password"
        type="password"
        placeholder="Enter your password"
        error={errors.password?.message}
        {...register("password", { required: "Password is required", min: 5 })}
      />

      <Button type="submit" style={{ marginTop: 16 }} solid>
        Signup
      </Button>
    </form>
  );
};
