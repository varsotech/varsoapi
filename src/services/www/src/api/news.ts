import { AxiosResponse } from "axios";
import { UseQueryOptions, useMutation, useQuery } from "./reactquery";
import { BlurToggleRequest, GetNewsResponse } from "../proto/src/app/requests";
import axios from "./axios";

export function useNews(
  options?: UseQueryOptions<AxiosResponse<GetNewsResponse>>
) {
  return useQuery<GetNewsResponse>(
    async () => await axios.get("/api/v1/app/news"),
    {
      queryKey: ["news"],
      ...options,
    }
  );
}

export function useToggleBlur(newsItemId: string) {
  return useMutation<unknown, AxiosResponse<any>>(
    async () =>
      await axios.post(
        "/api/v1/app/news/item/blur/toggle",
        BlurToggleRequest.create({
          rssItemId: newsItemId,
        })
      ),
    { mutationKey: ["news"] }
  );
}
