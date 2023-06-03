import { queryClient } from "@/shared/api/api";
import { AuthContextProvider } from "@/shared/context/auth.context";
import { ProvidersLayout } from "@/shared/layouts/providers-layout";
import "@/shared/styles/globals.css";
import { NextComponentAuth } from "@/shared/types/page/page.type";
import type { AppProps } from "next/app";
import { QueryClientProvider } from "@tanstack/react-query";

type AuthAppProps = AppProps & NextComponentAuth;

export default function App({ Component, pageProps }: AuthAppProps) {
  return (
    <QueryClientProvider client={queryClient}>
      <AuthContextProvider Component={Component}>
        <ProvidersLayout>
          <Component {...pageProps} />
        </ProvidersLayout>
      </AuthContextProvider>
    </QueryClientProvider>
  );
}
