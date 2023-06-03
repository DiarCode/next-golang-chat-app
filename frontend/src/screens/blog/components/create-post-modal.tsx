import { Button } from "@/shared/ui/button";
import { Modal } from "@/shared/ui/modal";
import React, { useContext } from "react";
import { CreatePostModalContext } from "../context/create-post-modal.context";

export const CreatePostModal = () => {
  const { visible, setVisible } = useContext(CreatePostModalContext);

  return (
    <Modal visible={visible} setVisibleState={setVisible}>
      <h1 className="font-semibold text-xl">Create Post</h1>

      <div className="mt-4">
        <div className="mt-6">
          <p className="font-semibold px-4 mb-2">Title</p>
          <input
            className="w-full p-3 text-base font-semibold bg-transparent focus:outline-none border-2 rounded-md"
            placeholder="Enter title"
          />
        </div>

        <div className="mt-6">
          <p className="font-semibold px-4 mb-2">Content</p>
          <textarea
            className="w-full h-32 p-3 text-base font-semibold bg-transparent focus:outline-none border-2 rounded-md"
            placeholder="Write post content"
          />
        </div>
      </div>

      <div className="mt-6">
        <Button style={{ width: "100%" }} solid>
          Create
        </Button>
      </div>
    </Modal>
  );
};
