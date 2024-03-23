import { AxiosResponse } from "axios";
import { UseQueryOptions, useQuery } from "./reactquery";
import { GetNewsResponse } from "@varsotech/varsoapi/src/app/requests";
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
