import axios, { AxiosInstance } from "axios"

const axiosInstance: AxiosInstance = fetchClient()

function fetchClient() {
  return axios.create({
    baseURL: process.env.SYSTEM_EXTERNAL_URL,
  })
}

export function setAxiosToken(token: string, deleteTokenFn: () => void) {
  // Add an interceptor to inject the token in the Authorization header
  axiosInstance.interceptors.request.use(function (config) {
    config.headers.Authorization = token ? `Bearer ${token}` : ""
    return config
  })

  // Add an interceptor to handle 401 responses
  axiosInstance.interceptors.response.use(
    (response) => response,
    (error) => {
      if (error.response && error.response.status === 401) {
        deleteTokenFn()
      }
    }
  )
}

export default axiosInstance
