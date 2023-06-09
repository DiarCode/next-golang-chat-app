import { AuthState, LoginDto, SignupDto } from "@/shared/types/auth/auth.dto";
import axios from "axios";
import { getApiUrl } from "../api";
import { User } from "@/shared/types/user/user.type";

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

  public static getUserById(id: number) {
    return axios.get<User>(getApiUrl(`users/${id}`));
  }
}
