import {
  Dispatch,
  FC,
  PropsWithChildren,
  SetStateAction,
  createContext,
  useMemo,
  useState,
} from "react";

interface CreatePostModalContextState {
  visible: boolean;
  setVisible: Dispatch<SetStateAction<boolean>>;
}

export const CreatePostModalContext = createContext<CreatePostModalContextState>(
  {} as CreatePostModalContextState
);

export const CreatePostModalContextProvider: FC<PropsWithChildren> = ({
  children,
}) => {
  const [visible, setVisible] = useState(false);

  const memoValue: CreatePostModalContextState = useMemo(
    () => ({ visible, setVisible }),
    [visible]
  );

  return (
    <CreatePostModalContext.Provider value={memoValue}>
      {children}
    </CreatePostModalContext.Provider>
  );
};
