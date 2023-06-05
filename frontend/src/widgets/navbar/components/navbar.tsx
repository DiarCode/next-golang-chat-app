import { Logo } from "@/shared/components/logo/logo";
import { DesktopNavigation } from "./desktop-navigation";
import { AuthNavigation } from "./auth-navigation";
import { MobileNavigation } from "./mobile-navigation";

export const Navbar = () => {
  return (
    <nav className="sticky top-0 flex z-20 justify-between items-center py-6 mb-6 bg-white border-b border-gray-200">
      <Logo />

      {/* Appear on mobile screen */}
      <section className="sm:hidden flex items-center gap-x-4">
        <AuthNavigation />
        <MobileNavigation />
      </section>

      <DesktopNavigation />

      <div className="hidden sm:block">
        <AuthNavigation />
      </div>
    </nav>
  );
};
