export interface User {
  id: number;
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
