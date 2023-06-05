import { PAGES_LINKS } from "@/shared/config/links.config";
import { Button } from "@/shared/ui/button";
import Link from "next/link";
import React from "react";
import { ProfileNavigation } from "./profile-navigation";
import { useAuth } from "@/shared/hooks/useAuth";

export const AuthNavigation = () => {
  const { isAuth } = useAuth();

  return (
    <>
      {!isAuth ? (
        <div className="flex items-center gap-x-4">
          <Link href={PAGES_LINKS.Login.link} className="hidden sm:block">
            <p>{PAGES_LINKS.Login.name}</p>
          </Link>

          <Link href={PAGES_LINKS.Signup.link}>
            <Button solid>
              <p className="text-white">{PAGES_LINKS.Signup.name}</p>
            </Button>
          </Link>
        </div>
      ) : (
        <div className="flex items-center gap-x-5">
          <ProfileNavigation />
        </div>
      )}
    </>
  );
};
