import { COLORS } from "@/shared/config/colors.config";
import React from "react";
import { MessageCircle, Monitor, Smile } from "react-feather";

const offerIcons = [
  {
    title: "User-friendly experience",
    icon: <Smile size={40} color={COLORS.primary} />,
    text: "High accecability and user experience",
  },
  {
    title: "Collaboration in chats",
    icon: <MessageCircle size={40} color={COLORS.primary} />,
    text: "Make your comunication in chat easy",
  },
  {
    title: "Just use any devices",
    icon: <Monitor size={40} color={COLORS.primary} />,
    text: "Use any device you wish, and get in queue online",
  },
];

export const HomeOfferSection = () => {
  return (
    <div className="bg-gray-100 w-full bg-cgray py-8 rounded-xl mt-16 md:mt-44 flex flex-col items-center justify-center">
      <p className="text-sm text-gray-500">Benefits for you</p>
      <h1
        className="mt-3 font-bold text-3xl md:text-5xl text-primary"
      >
        Our Offer
      </h1>
      <p className="mt-5 text-center text-gray-500 text-sm">
        Mission of Meowchat is to make your life easier in
        <br />
        boring collaboration with digital and modern approach
      </p>

      <div className="mt-7 flex items-center justify-center flex-wrap sm:flex-nowrap gap-x-5">
        {offerIcons.map(item => (
          <div
            key={item.title}
            className="flex flex-col items-center justify-start w-full sm:w-1/4 p-5"
          >
            <div className="h-full">{item.icon}</div>
            <p className="text-center mt-6 font-semibold text-black">
              {item.title}
            </p>
            <p className="text-center mt-2 text-gray-500 text-sm">
              {item.text}
            </p>
          </div>
        ))}
      </div>
    </div>
  );
};
