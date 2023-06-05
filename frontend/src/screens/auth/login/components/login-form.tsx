import { useAuth } from "@/shared/hooks/useAuth";
import { LoginDto } from "@/shared/types/auth/auth.dto";
import { Button } from "@/shared/ui/button";
import { InputField } from "@/shared/ui/input";
import React from "react";
import { useForm } from "react-hook-form";

export const LoginForm = () => {
  const { login } = useAuth();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginDto>();

  const onLoginSubmit = (data: LoginDto) => {
    console.log(data);
    login(data);
  };

  return (
    <form
      onSubmit={handleSubmit(onLoginSubmit)}
      className="mt-8 md:mt-12 flex flex-col gap-y-4"
    >
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
        Login
      </Button>
    </form>
  );
};
