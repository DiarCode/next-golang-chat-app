import { Button } from "@/shared/ui/button";
import { Modal } from "@/shared/ui/modal";
import React, { useContext } from "react";
import { CreateChatModalContext } from "../context/create-chat-modal.context";

export const CreateChatModal = () => {
  const { visible, setVisible } = useContext(CreateChatModalContext);

  return (
    <Modal visible={visible} setVisibleState={setVisible}>
      <h1 className="font-semibold text-xl">Create Chat</h1>

      <div className="mt-4">
        <input
          type="text"
          className=" w-full h-28 p-0 text-3xl md:text-4xl font-semibold bg-transparent focus:outline-none border-b-2"
          placeholder="Enter chat name"
        />
      </div>

      <div className="mt-10">
        <Button style={{ width: "100%" }} solid>
          Create
        </Button>
      </div>
    </Modal>
  );
};
