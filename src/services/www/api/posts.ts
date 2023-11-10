import { GetPostsResponse } from "@varsotech/varsoapi/src/app/requests"
import axios from "./axios"
import { AxiosResponse } from "axios"
import { useQuery, UseQueryOptions } from "./reactquery"

export function usePosts(options?: UseQueryOptions<AxiosResponse<GetPostsResponse>>) {
  return useQuery(async () => await axios.get("/api/v1/app/posts"), {
    queryKey: ["posts"],
    ...options,
  })
}

// export function useCreateTask() {
//   return useMutation<CreateTaskRequest, AxiosResponse<CreateTaskResponse>>(
//     async (req) => await axios.post("/api/v1/sidekick/tasks", req),
//     { mutationKey: ["tasks"] }
//   )
// }

// export function useUpdateTask() {
//   return useMutation<UpdateTaskRequest, AxiosResponse<UpdateTaskResponse>>(
//     async (req) => await axios.post("/api/v1/sidekick/tasks/update", req),
//     { mutationKey: ["tasks"] }
//   )
// }

// export function useDeleteTask() {
//   return useMutation<DeleteTaskRequest, AxiosResponse<DeleteTaskResponse>>(
//     async (req) => await axios.post("/api/v1/sidekick/tasks/delete", req),
//     {
//       mutationKey: ["tasks"],
//     }
//   )
// }
