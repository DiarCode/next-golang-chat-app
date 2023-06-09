import { Button } from "@/shared/ui/button";
import { Modal } from "@/shared/ui/modal";
import React, { useContext } from "react";
import { CreateChatModalContext } from "../context/create-chat-modal.context";
import { useForm } from "react-hook-form";
import { CreateChatDto } from "@/shared/types/chat/chat.dto";
import { ChatApiService } from "@/shared/api/chat/chat.api";
import { useChats } from "@/shared/hooks/useChat";

export const CreateChatModal = () => {
  const { register, handleSubmit, reset } = useForm<CreateChatDto>();
  const { addChat } = useChats();
  const { visible, setVisible } = useContext(CreateChatModalContext);

  const onCreateChatSubmit = async (data: CreateChatDto) => {
    console.log(data);
    const res = await ChatApiService.createChatRoom(data);
    if (res.status !== 200) {
      console.log(res);
      setVisible(false);
      return;
    }

    reset();
    addChat(res.data);
    setVisible(false);
  };

  return (
    <Modal visible={visible} setVisibleState={setVisible}>
      <h1 className="font-semibold text-xl">Create Chat</h1>

      <form className="mt-4" onSubmit={handleSubmit(onCreateChatSubmit)}>
        <input
          {...register("name", { required: true })}
          type="text"
          className=" w-full h-28 p-0 text-3xl md:text-4xl font-semibold bg-transparent focus:outline-none border-b-2"
          placeholder="Enter chat name"
        />

        <div className="mt-10">
          <Button type="submit" style={{ width: "100%" }} solid>
            Create
          </Button>
        </div>
      </form>
    </Modal>
  );
};
