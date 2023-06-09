import {
  Dispatch,
  FC,
  PropsWithChildren,
  SetStateAction,
  createContext,
  useMemo,
  useState,
} from "react";

interface CreateChatModalContextState {
  visible: boolean;
  setVisible: Dispatch<SetStateAction<boolean>>;
}

export const CreateChatModalContext =
  createContext<CreateChatModalContextState>({} as CreateChatModalContextState);

export const CreateChatModalContextProvider: FC<PropsWithChildren> = ({
  children,
}) => {
  const [visible, setVisible] = useState(false);

  const memoValue: CreateChatModalContextState = useMemo(
    () => ({ visible, setVisible }),
    [visible]
  );

  return (
    <CreateChatModalContext.Provider value={memoValue}>
      {children}
    </CreateChatModalContext.Provider>
  );
};
