import { NextPage } from "next";

export interface AccessTypes {
  onlyUser?: boolean;
}

export type NextPageAuth<T = {}> = NextPage<T> & AccessTypes;

export interface NextComponentAuth {
  Component: AccessTypes;
}
