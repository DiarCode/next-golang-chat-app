export interface LoginDto {
  email: string;
  password: string;
}

export interface SignupDto {
  username: string;
  email: string;
  password: string;
}

export interface AuthState {
  id: number;
  token: string;
  username: string;
  email: string;
}
