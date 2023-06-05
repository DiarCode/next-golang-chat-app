import dayjs from "dayjs";

export const getDateFromUnix = (unix: number): Date => {
  let date = new Date();
  date.setTime(unix);

  return date;
};

export const formatDateFromUnix = (
  unix: number,
  format: string = "hh:mm A"
): string => {
  const date = getDateFromUnix(unix);
  return dayjs(date).format(format);
};
