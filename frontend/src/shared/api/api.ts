import { QueryClient } from "@tanstack/react-query";

export interface ApiResponse<T> {
  data: T;
  message?: string;
}

export const API_HOST = process.env.API_HOST ?? "localhost";
export const API_URL = `http://${API_HOST}:8080/api/v1`;

export const getApiUrl = (path: string) => {
  return `${API_URL}/${path}`;
};

export const queryClient = new QueryClient();
