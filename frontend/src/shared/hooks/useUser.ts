import { useQuery } from "@tanstack/react-query";
import { AuthApiService } from "../api/auth/auth.api";

export const useUser = (id: number | undefined | null) =>
  useQuery({
    queryKey: ["user", id],
    queryFn: () => AuthApiService.getUserById(id!),
    enabled: Boolean(id)
  });
