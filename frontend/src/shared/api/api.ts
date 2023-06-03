import { QueryClient } from "@tanstack/react-query";

export interface ApiResponse<T> {
  data: T;
  message?: string;
}

export const API_URL = process.env.API_URL ?? "http://localhost:8080/api/v1";

export const getApiUrl = (path: string) => {
  return `${API_URL}/${path}`;
};

export const queryClient = new QueryClient();
