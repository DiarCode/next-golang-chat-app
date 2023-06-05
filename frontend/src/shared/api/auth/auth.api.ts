import { AuthState, LoginDto, SignupDto } from "@/shared/types/auth/auth.dto";
import axios from "axios";
import { ApiResponse, getApiUrl } from "../api";

export class AuthApiService {
  public static login(dto: LoginDto) {
    return axios.post<AuthState & { message?: string }>(
      getApiUrl("auth/login"),
      dto
    );
  }

  public static signup(dto: SignupDto) {
    return axios.post<AuthState & { message?: string }>(
      getApiUrl("auth/signup"),
      dto
    );
  }
}
