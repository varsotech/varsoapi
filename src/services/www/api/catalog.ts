import { CatalogSearchRequest, CatalogSearchResponse } from "@varsotech/varsoapi/src/app/requests"
import axios from "./axios"
import { AxiosResponse } from "axios"
import { useMutation, useQuery, UseQueryOptions } from "./reactquery"

export function useCatalogSearch() {
  return useMutation<CatalogSearchRequest, AxiosResponse<CatalogSearchResponse>>(
    async (req) => await axios.post("/api/v1/catalog/search", req),
    { mutationKey: ["catalog_search"] }
  )
}

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
